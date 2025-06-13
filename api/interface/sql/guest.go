package sql

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"basic-service/domain"
	"basic-service/gen/db/model"
	"basic-service/gen/db/table"

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
		return err
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
	return err
}

func (r *GuestManager) List(ctx context.Context, userTemplateID string, pageSize int, offset int) ([]domain.Guest, error) {
	var guests []model.Guests

	stmt := sqlite.SELECT(
		table.Guests.AllColumns,
	).FROM(
		table.Guests,
	).WHERE(
		table.Guests.UserTemplateID.EQ(sqlite.String(userTemplateID)),
	).ORDER_BY(
		table.Guests.CreatedAt.DESC(),
	).LIMIT(
		int64(pageSize),
	).OFFSET(
		int64(offset),
	)

	err := stmt.QueryContext(ctx, r.db.db, &guests)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return []domain.Guest{}, nil
		}
		return nil, err
	}

	result := make([]domain.Guest, 0, len(guests))
	for _, g := range guests {
		var tags []string
		if err := json.Unmarshal([]byte(g.Tags), &tags); err != nil {
			return nil, err
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

	return result, nil
}

func (r *GuestManager) Update(ctx context.Context, guestID string, guest domain.Guest) error {
	tagsJSON, err := json.Marshal(guest.Tags)
	if err != nil {
		return err
	}

	stmt := table.Guests.UPDATE().
		SET(
			table.Guests.Name.SET(sqlite.String(guest.Name)),
			table.Guests.GroupName.SET(sqlite.String(guest.Group)),
			table.Guests.Person.SET(sqlite.Int(int64(guest.Person))),
			table.Guests.Tags.SET(sqlite.String(string(tagsJSON))),
			table.Guests.Telp.SET(sqlite.String(guest.Telp)),
			table.Guests.Address.SET(sqlite.String(guest.Address)),
			table.Guests.Message.SET(sqlite.String(guest.Message)),
			table.Guests.Attend.SET(sqlite.Bool(guest.Attend)),
			table.Guests.ViewAt.SET(sqlite.DATETIME(guest.ViewAt)),
			table.Guests.CreatedAt.SET(sqlite.DATETIME(guest.CreatedAt)),
		).WHERE(
		table.Guests.ID.EQ(sqlite.String(guestID)),
	)

	result, err := stmt.ExecContext(ctx, r.db.db)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no guest found with the given ID")
	}

	return nil
}

func (r *GuestManager) Delete(ctx context.Context, guestID string) error {
	stmt := table.Guests.DELETE().
		WHERE(table.Guests.ID.EQ(sqlite.String(guestID)))

	result, err := stmt.ExecContext(ctx, r.db.db)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no guest found with the given ID")
	}

	return nil
}
