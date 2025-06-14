package sql

import (
	"context"
	"encoding/json"
	"errors"

	"basic-service/domain"
	"basic-service/gen/db/model"
	"basic-service/gen/db/table"

	"braces.dev/errtrace"
	"github.com/go-jet/jet/v2/sqlite"
)

var (
	ErrUserTemplateNotFound = errors.New("user template not found")
	ErrInvalidTemplateData  = errors.New("invalid template data")
)

type UserTemplateRepository struct {
	db *SQLite
}

func NewUserTemplateRepository(db *SQLite) *UserTemplateRepository {
	return &UserTemplateRepository{db: db}
}

type UserTemplateList struct {
	Total int64
	Data  []domain.UserTemplate
}

// Count returns the total number of user templates
func (r *UserTemplateRepository) Count(ctx context.Context) (int64, error) {
	countStmt := sqlite.SELECT(
		sqlite.COUNT(table.UserTemplates.ID).AS("total"),
	).FROM(table.UserTemplates)

	var total struct {
		Total int64
	}
	if err := countStmt.QueryContext(ctx, r.db.db, &total); err != nil {
		return 0, errtrace.Wrap(err)
	}

	return total.Total, nil
}

// List returns a paginated list of user templates
func (r *UserTemplateRepository) List(ctx context.Context, offset, limit int) (UserTemplateList, error) {
	// Get total count
	total, err := r.Count(ctx)
	if err != nil {
		return UserTemplateList{}, errtrace.Wrap(err)
	}

	// Get paginated templates
	stmt := table.UserTemplates.SELECT(
		table.UserTemplates.ID,
		table.UserTemplates.UserID,
		table.UserTemplates.BaseTemplateID,
		table.UserTemplates.State,
		table.UserTemplates.Slug,
		table.UserTemplates.URL,
		table.UserTemplates.MessageTemplate,
		table.UserTemplates.Name,
		table.UserTemplates.CoverImage,
		table.UserTemplates.CreatedAt,
		table.UserTemplates.UpdatedAt,
		table.UserTemplates.ExpireAt,
	).ORDER_BY(
		table.UserTemplates.CreatedAt.DESC(),
	).OFFSET(
		int64(offset),
	).LIMIT(
		int64(limit),
	)

	var dbTemplates []model.UserTemplates
	if err := stmt.QueryContext(ctx, r.db.db, &dbTemplates); err != nil {
		return UserTemplateList{}, errtrace.Wrap(err)
	}

	templates := make([]domain.UserTemplate, 0, len(dbTemplates))
	for _, dbTemplate := range dbTemplates {
		template, err := r.mapToDomain(dbTemplate)
		if err != nil {
			return UserTemplateList{}, errtrace.Wrap(err)
		}
		templates = append(templates, template)
	}

	return UserTemplateList{
		Total: total,
		Data:  templates,
	}, nil
}

// Get retrieves a single user template by ID
func (r *UserTemplateRepository) Get(ctx context.Context, id string) (domain.UserTemplate, error) {
	stmt := table.UserTemplates.SELECT(
		table.UserTemplates.AllColumns,
	).WHERE(
		table.UserTemplates.ID.EQ(sqlite.String(id)),
	).LIMIT(1)

	var dbTemplate model.UserTemplates
	if err := stmt.QueryContext(ctx, r.db.db, &dbTemplate); err != nil {
		return domain.UserTemplate{}, errtrace.Wrap(err)
	}

	if dbTemplate.ID == "" {
		return domain.UserTemplate{}, errtrace.Wrap(ErrUserTemplateNotFound)
	}

	return r.mapToDomain(dbTemplate)
}

// Exists checks if a template with the given ID exists
func (r *UserTemplateRepository) Exists(ctx context.Context, id string) (bool, error) {
	stmt := sqlite.SELECT(
		sqlite.EXISTS(
			table.UserTemplates.SELECT(table.UserTemplates.ID).
				WHERE(table.UserTemplates.ID.EQ(sqlite.String(id))),
		),
	)

	var exists bool
	if err := stmt.QueryContext(ctx, r.db.db, &exists); err != nil {
		return false, errtrace.Wrap(err)
	}
	return exists, nil
}

// Create adds a new user template
func (r *UserTemplateRepository) Create(ctx context.Context, template domain.UserTemplate) error {
	msgTemplateStr, err := r.marshalMessageTemplate(template.MessageTemplate)
	if err != nil {
		return errtrace.Wrap(err)
	}

	stmt := table.UserTemplates.INSERT(
		table.UserTemplates.AllColumns,
	).VALUES(
		sqlite.String(template.ID),
		sqlite.String(template.UserID),
		sqlite.String(template.BaseTemplateID),
		sqlite.Int(int64(template.State)),
		sqlite.String(template.Slug),
		sqlite.String(template.URL),
		sqlite.String(msgTemplateStr),
		sqlite.String(template.Name),
		sqlite.String(template.CoverImage),
		sqlite.DATETIME(template.CreatedAt),
		sqlite.DATETIME(template.UpdatedAt),
		sqlite.DATETIME(template.ExpireAt),
	)

	_, err = stmt.ExecContext(ctx, r.db.db)
	return errtrace.Wrap(err)
}

// Update modifies an existing user template
func (r *UserTemplateRepository) Update(ctx context.Context, templateID string, template domain.UserTemplate) error {
	exists, err := r.Exists(ctx, templateID)
	if err != nil {
		return errtrace.Wrap(err)
	}
	if !exists {
		return errtrace.Wrap(ErrUserTemplateNotFound)
	}

	msgTemplateStr, err := r.marshalMessageTemplate(template.MessageTemplate)
	if err != nil {
		return errtrace.Wrap(err)
	}

	stmt := table.UserTemplates.UPDATE().
		SET(
			table.UserTemplates.UserID.SET(sqlite.String(template.UserID)),
			table.UserTemplates.BaseTemplateID.SET(sqlite.String(template.BaseTemplateID)),
			table.UserTemplates.State.SET(sqlite.Int(int64(template.State))),
			table.UserTemplates.Slug.SET(sqlite.String(template.Slug)),
			table.UserTemplates.URL.SET(sqlite.String(template.URL)),
			table.UserTemplates.MessageTemplate.SET(sqlite.String(msgTemplateStr)),
			table.UserTemplates.Name.SET(sqlite.String(template.Name)),
			table.UserTemplates.CoverImage.SET(sqlite.String(template.CoverImage)),
			table.UserTemplates.UpdatedAt.SET(sqlite.DATETIME(template.UpdatedAt)),
			table.UserTemplates.ExpireAt.SET(sqlite.DATETIME(template.ExpireAt)),
		).WHERE(
		table.UserTemplates.ID.EQ(sqlite.String(templateID)),
	)

	_, err = stmt.ExecContext(ctx, r.db.db)
	return errtrace.Wrap(err)
}

// Delete removes a user template
func (r *UserTemplateRepository) Delete(ctx context.Context, id string) error {
	exists, err := r.Exists(ctx, id)
	if err != nil {
		return errtrace.Wrap(err)
	}
	if !exists {
		return errtrace.Wrap(ErrUserTemplateNotFound)
	}

	stmt := table.UserTemplates.DELETE().
		WHERE(table.UserTemplates.ID.EQ(sqlite.String(id)))

	_, err = stmt.ExecContext(ctx, r.db.db)
	return errtrace.Wrap(err)
}

// Helper function to map database model to domain model
func (r *UserTemplateRepository) mapToDomain(dbTemplate model.UserTemplates) (domain.UserTemplate, error) {
	var msgTemplates []domain.MessageTemplate
	if dbTemplate.MessageTemplate != "" {
		if err := json.Unmarshal([]byte(dbTemplate.MessageTemplate), &msgTemplates); err != nil {
			return domain.UserTemplate{}, errtrace.Wrap(err)
		}
	}

	return domain.UserTemplate{
		ID:              dbTemplate.ID,
		UserID:          dbTemplate.UserID,
		BaseTemplateID:  dbTemplate.BaseTemplateID,
		State:           int(dbTemplate.State),
		Slug:            dbTemplate.Slug,
		URL:             dbTemplate.URL,
		MessageTemplate: msgTemplates,
		Name:            dbTemplate.Name,
		CoverImage:      dbTemplate.CoverImage,
		CreatedAt:       dbTemplate.CreatedAt,
		UpdatedAt:       dbTemplate.UpdatedAt,
		ExpireAt:        dbTemplate.ExpireAt,
	}, nil
}

// Helper function to marshal message template
func (r *UserTemplateRepository) marshalMessageTemplate(msgTemplates []domain.MessageTemplate) (string, error) {
	if msgTemplates == nil {
		return "[]", nil
	}
	msgTemplateJSON, err := json.Marshal(msgTemplates)
	if err != nil {
		return "", errtrace.Wrap(err)
	}
	return string(msgTemplateJSON), nil
}
