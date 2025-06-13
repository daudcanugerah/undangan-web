package sql

import (
	"context"
	"encoding/json"
	"errors"

	"basic-service/domain"
	"basic-service/gen/db/model"
	"basic-service/gen/db/table"

	"github.com/go-jet/jet/v2/sqlite"
)

var ErrUserTemplateNotFound = errors.New("user template not found")

type UserTemplateRepository struct {
	db *SQLite
}

func NewUserTemplateRepository(db *SQLite) *UserTemplateRepository {
	return &UserTemplateRepository{db: db}
}

type UserTemplateList struct {
	Total int
	Data  []domain.UserTemplate
}

func (r *UserTemplateRepository) List(ctx context.Context) (UserTemplateList, error) {
	// Get total count
	countStmt := sqlite.SELECT(
		sqlite.COUNT(table.UserTemplates.ID).AS("total"),
	).FROM(table.UserTemplates)

	var total struct {
		Total int64
	}
	if err := countStmt.QueryContext(ctx, r.db.db, &total); err != nil {
		return UserTemplateList{}, err
	}

	// Get all templates
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
	)

	var dbTemplates []model.UserTemplates
	if err := stmt.QueryContext(ctx, r.db.db, &dbTemplates); err != nil {
		return UserTemplateList{}, err
	}

	templates := make([]domain.UserTemplate, 0, len(dbTemplates))
	for _, dbTemplate := range dbTemplates {
		var msgTemplates map[string]domain.MessageTemplate
		if err := json.Unmarshal([]byte(dbTemplate.MessageTemplate), &msgTemplates); err != nil {
			return UserTemplateList{}, err
		}

		templates = append(templates, domain.UserTemplate{
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
		})
	}

	return UserTemplateList{
		Total: int(total.Total),
		Data:  templates,
	}, nil
}

func (r *UserTemplateRepository) Create(ctx context.Context, template domain.UserTemplate) error {
	msgTemplateJSON, err := json.Marshal(template.MessageTemplate)
	if err != nil {
		return err
	}
	msgTemplateStr := string(msgTemplateJSON)

	stmt := table.UserTemplates.INSERT(
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
	).VALUES(
		template.UserID,
		template.BaseTemplateID,
		template.State,
		template.Slug,
		template.URL,
		msgTemplateStr,
		template.Name,
		template.CoverImage,
		sqlite.CURRENT_TIMESTAMP(),
		sqlite.CURRENT_TIMESTAMP(),
		sqlite.DATETIME(template.ExpireAt),
	)

	_, err = stmt.ExecContext(ctx, r.db.db)
	return err
}

func (r *UserTemplateRepository) Update(ctx context.Context, templateID string, template domain.UserTemplate) error {
	// First check if template exists
	existsStmt := sqlite.SELECT(
		sqlite.EXISTS(
			table.UserTemplates.SELECT(table.UserTemplates.ID).
				WHERE(table.UserTemplates.ID.EQ(sqlite.String(templateID))),
		),
	)

	var exists bool
	if err := existsStmt.QueryContext(ctx, r.db.db, &exists); err != nil {
		return err
	}
	if !exists {
		return ErrUserTemplateNotFound
	}

	msgTemplateJSON, err := json.Marshal(template.MessageTemplate)
	if err != nil {
		return err
	}
	msgTemplateStr := string(msgTemplateJSON)

	stmt := table.UserTemplates.UPDATE().
		SET(
			table.UserTemplates.UserID.SET(sqlite.String(template.ID)),
			table.UserTemplates.BaseTemplateID.SET(sqlite.String(template.BaseTemplateID)),
			table.UserTemplates.State.SET(sqlite.Int(int64(template.State))),
			table.UserTemplates.Slug.SET(sqlite.String(template.Slug)),
			table.UserTemplates.URL.SET(sqlite.String(template.URL)),
			table.UserTemplates.MessageTemplate.SET(sqlite.String(msgTemplateStr)),
			table.UserTemplates.Name.SET(sqlite.String(template.Name)),
			table.UserTemplates.CoverImage.SET(sqlite.String(template.CoverImage)),
			table.UserTemplates.UpdatedAt.SET(sqlite.CURRENT_TIMESTAMP()),
			table.UserTemplates.ExpireAt.SET(sqlite.DATETIME(template.ExpireAt)),
		).WHERE(
		table.UserTemplates.ID.EQ(sqlite.String(templateID)),
	)

	_, err = stmt.ExecContext(ctx, r.db.db)
	return err
}
