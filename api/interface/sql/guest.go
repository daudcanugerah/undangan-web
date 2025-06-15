package sql

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"basic-service/domain"
	"basic-service/gen/db/model"
	"basic-service/gen/db/table"

	"braces.dev/errtrace"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/go-jet/jet/v2/sqlite"
)

type GuestManager struct {
	db *SQLite
}

func NewGuestManager(db *SQLite) *GuestManager {
	return &GuestManager{db: db}
}

func (r *GuestManager) Create(ctx context.Context, guest domain.Guest) error {
	tagsJSON, err := json.Marshal(guest.Tags)
	if err != nil {
		return errtrace.Wrap(err)
	}

	stmt := table.Guests.INSERT(
		table.Guests.ID,
		table.Guests.UserTemplateID,
		table.Guests.Name,
		table.Guests.GroupName,
		table.Guests.Person,
		table.Guests.Tags,
		table.Guests.Telp,
		table.Guests.Address,
		table.Guests.Message,
		table.Guests.Attend,
		table.Guests.ViewAt,
		table.Guests.CreatedAt,
	).VALUES(
		guest.ID,
		guest.UserTemplateID,
		guest.Name,
		guest.Group,
		guest.Person,
		tagsJSON,
		guest.Telp,
		guest.Address,
		guest.Message,
		guest.Attend,
		guest.ViewAt,
		time.Now(),
	)

	_, err = stmt.ExecContext(ctx, r.db.db)
	return errtrace.Wrap(err)
}

func (r *GuestManager) List(ctx context.Context, userID, userTemplateID string, page, pageSize int) ([]domain.Guest, int64, error) {
	var guests []model.Guests

	// Calculate offset based on page and pageSize
	offset := (page - 1) * pageSize

	// Modified query with JOIN and user_id check
	stmt := sqlite.SELECT(
		table.Guests.AllColumns,
	).FROM(
		table.Guests.INNER_JOIN(table.UserTemplates,
			table.UserTemplates.ID.EQ(table.Guests.UserTemplateID)),
	).WHERE(
		table.Guests.UserTemplateID.EQ(sqlite.String(userTemplateID)).
			AND(table.UserTemplates.UserID.EQ(sqlite.String(userID))),
	).ORDER_BY(
		table.Guests.CreatedAt.DESC(),
	).LIMIT(
		int64(pageSize),
	).OFFSET(
		int64(offset),
	)

	fmt.Println("SQL Query:", stmt.DebugSql())

	err := stmt.QueryContext(ctx, r.db.db, &guests)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return []domain.Guest{}, 0, nil
		}
		return nil, 0, errtrace.Wrap(err)
	}

	// Get total count with the same user_id filter
	total, err := r.getFilteredCount(ctx, userID, userTemplateID)
	if err != nil {
		return nil, 0, errtrace.Wrap(err)
	}

	result := make([]domain.Guest, 0, len(guests))
	for _, g := range guests {
		var tags []string
		if err := json.Unmarshal([]byte(g.Tags), &tags); err != nil {
			return nil, 0, errtrace.Wrap(err)
		}

		result = append(result, domain.Guest{
			ID:             g.ID,
			UserTemplateID: g.UserTemplateID,
			Name:           g.Name,
			Group:          g.GroupName,
			Person:         int(g.Person),
			Tags:           tags,
			Telp:           g.Telp,
			Address:        g.Address,
			Message:        g.Message,
			Attend:         g.Attend,
			ViewAt:         g.ViewAt,
			CreatedAt:      g.CreatedAt,
		})
	}

	return result, total, nil
}

// New helper function for filtered count
func (r *GuestManager) getFilteredCount(ctx context.Context, userID, userTemplateID string) (int64, error) {
	var count int64

	stmt := sqlite.SELECT(
		sqlite.COUNT(sqlite.STAR).AS("total"),
	).FROM(
		table.Guests.
			INNER_JOIN(table.UserTemplates,
				table.UserTemplates.ID.EQ(table.Guests.UserTemplateID)),
	).WHERE(
		table.Guests.UserTemplateID.EQ(sqlite.String(userTemplateID)).
			AND(table.UserTemplates.UserID.EQ(sqlite.String(userID))),
	)

	var total struct {
		Total int64
	}

	err := stmt.QueryContext(ctx, r.db.db, &total)
	if err != nil {
		return 0, errtrace.Wrap(err)
	}
	count = total.Total

	return count, nil
}

func (r *GuestManager) Get(ctx context.Context, guestID string) (*domain.Guest, error) {
	var guest model.Guests

	stmt := sqlite.SELECT(
		table.Guests.AllColumns,
	).FROM(
		table.Guests,
	).WHERE(
		table.Guests.ID.EQ(sqlite.String(guestID)),
	).LIMIT(1)

	err := stmt.QueryContext(ctx, r.db.db, &guest)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return nil, nil
		}
		return nil, errtrace.Wrap(err)
	}

	var tags []string
	if err := json.Unmarshal([]byte(guest.Tags), &tags); err != nil {
		return nil, errtrace.Wrap(err)
	}

	return &domain.Guest{
		ID:             guest.ID,
		UserTemplateID: guest.UserTemplateID,
		Name:           guest.Name,
		Group:          guest.GroupName,
		Person:         int(guest.Person),
		Tags:           tags,
		Telp:           guest.Telp,
		Address:        guest.Address,
		Message:        guest.Message,
		Attend:         guest.Attend,
		ViewAt:         guest.ViewAt,
		CreatedAt:      guest.CreatedAt,
	}, nil
}

func (r *GuestManager) Update(ctx context.Context, guestID string, guest domain.Guest) error {
	tagsJSON, err := json.Marshal(guest.Tags)
	if err != nil {
		return errtrace.Wrap(err)
	}

	setList := []interface{}{
		table.Guests.Name.SET(sqlite.String(guest.Name)),
		table.Guests.GroupName.SET(sqlite.String(guest.Group)),
		table.Guests.Person.SET(sqlite.Int(int64(guest.Person))),
		table.Guests.Tags.SET(sqlite.String(string(tagsJSON))),
		table.Guests.Telp.SET(sqlite.String(guest.Telp)),
		table.Guests.Address.SET(sqlite.String(guest.Address)),
		table.Guests.Message.SET(sqlite.String(guest.Message)),
		table.Guests.ViewAt.SET(sqlite.DATETIME(guest.ViewAt)),
		table.Guests.CreatedAt.SET(sqlite.DATETIME(guest.CreatedAt)),
	}

	if guest.Attend != nil {
		setList = append(setList, table.Guests.Attend.SET(sqlite.Bool(*guest.Attend)))
	}

	stmt := table.Guests.UPDATE().
		SET(setList[0], setList[1:]...).WHERE(
		table.Guests.ID.EQ(sqlite.String(guestID)),
	)

	if guest.Attend != nil {
		stmt = stmt.SET(table.Guests.Attend.SET(sqlite.Bool(*guest.Attend)))
	}

	result, err := stmt.ExecContext(ctx, r.db.db)
	if err != nil {
		return errtrace.Wrap(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errtrace.Wrap(err)
	}

	if rowsAffected == 0 {
		return errtrace.Wrap(errors.New("no guest found with the given ID"))
	}

	return nil
}

func (r *GuestManager) UpdateMessageAndLastView(ctx context.Context, guestID, message string, attend *bool) error {
	// Validate inputs
	if guestID == "" {
		return errtrace.Wrap(errors.New("guestID cannot be empty"))
	}
	setList := []interface{}{
		table.Guests.Message.SET(sqlite.String(message)),
		table.Guests.ViewAt.SET(sqlite.DATETIME(time.Now())),
	}

	if attend != nil {
		setList = append(setList, table.Guests.Attend.SET(sqlite.Bool(*attend)))
	}

	// Prepare base update statement
	stmt := table.Guests.UPDATE().
		SET(setList[0], setList[1:]...).
		WHERE(table.Guests.ID.EQ(sqlite.String(guestID)))

	fmt.Println("SQL Update Query:", stmt.DebugSql())

	// Execute the statement
	result, err := stmt.ExecContext(ctx, r.db.db)
	if err != nil {
		return errtrace.Wrap(fmt.Errorf("failed to execute update: %w", err))
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errtrace.Wrap(fmt.Errorf("failed to get rows affected: %w", err))
	}

	if rowsAffected == 0 {
		return errtrace.Wrap(fmt.Errorf("no guest found with ID: %s", guestID))
	}

	return nil
}

func (r *GuestManager) Delete(ctx context.Context, guestID string) error {
	stmt := table.Guests.DELETE().
		WHERE(table.Guests.ID.EQ(sqlite.String(guestID)))

	result, err := stmt.ExecContext(ctx, r.db.db)
	if err != nil {
		return errtrace.Wrap(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errtrace.Wrap(err)
	}

	if rowsAffected == 0 {
		return errtrace.Wrap(errors.New("no guest found with the given ID"))
	}

	return nil
}
