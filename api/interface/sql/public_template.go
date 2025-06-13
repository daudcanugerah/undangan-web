package sql

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"basic-service/domain"
	"basic-service/gen/db/model"
	"basic-service/gen/db/table"

	"github.com/go-jet/jet/v2/sqlite"
)

type PublicTemplate struct {
	db *SQLite
}

func NewPublicTemplateRepository(db *SQLite) *PublicTemplate {
	return &PublicTemplate{db: db}
}

func (r *PublicTemplate) Create(ctx context.Context, template domain.PublicTemplate) error {
	tagsJSON, err := json.Marshal(template.Tags)
	if err != nil {
		return err
	}

	stmt := table.PublicTemplates.INSERT(
		table.PublicTemplates.ID,
		table.PublicTemplates.Name,
		table.PublicTemplates.Description,
		table.PublicTemplates.PriceInterval,
		table.PublicTemplates.Price,
		table.PublicTemplates.Type,
		table.PublicTemplates.Tags,
		table.PublicTemplates.CoverImage,
		table.PublicTemplates.State,
		table.PublicTemplates.CreatedAt,
		table.PublicTemplates.UpdatedAt,
	).VALUES(
		template.ID,
		template.Name,
		template.Description,
		template.PriceInterval,
		template.Price,
		template.Type,
		tagsJSON,
		template.CoverImage,
		template.State,
		template.CreatedAt,
		template.UpdatedAt,
	)

	_, err = stmt.ExecContext(ctx, r.db.db)
	return err
}

func (r *PublicTemplate) List(ctx context.Context) ([]domain.PublicTemplate, error) {
	var templates []model.PublicTemplates

	stmt := sqlite.SELECT(
		table.PublicTemplates.AllColumns,
	).FROM(
		table.PublicTemplates,
	).ORDER_BY(
		table.PublicTemplates.CreatedAt.DESC(),
	)

	err := stmt.QueryContext(ctx, r.db.db, &templates)
	if err != nil {
		return nil, err
	}

	result := make([]domain.PublicTemplate, 0, len(templates))
	for _, t := range templates {
		var tags []string
		if err := json.Unmarshal([]byte(t.Tags), &tags); err != nil {
			return nil, err
		}

		result = append(result, domain.PublicTemplate{
			ID:            t.ID,
			Name:          t.Name,
			Description:   t.Description,
			PriceInterval: t.PriceInterval,
			Price:         int(t.Price),
			Type:          t.Type,
			Tags:          tags,
			CoverImage:    t.CoverImage,
			State:         int(t.State),
			CreatedAt:     t.CreatedAt,
			UpdatedAt:     t.UpdatedAt,
		})
	}

	return result, nil
}

func (r *PublicTemplate) Update(ctx context.Context, templateID string, template domain.PublicTemplate) error {
	tagsJSON, err := json.Marshal(template.Tags)
	if err != nil {
		return err
	}

	stmt := table.PublicTemplates.UPDATE().
		SET(
			table.PublicTemplates.Name.SET(sqlite.String(template.Name)),
			table.PublicTemplates.Description.SET(sqlite.String(template.Description)),
			table.PublicTemplates.PriceInterval.SET(sqlite.String(template.PriceInterval)),
			table.PublicTemplates.Price.SET(sqlite.Int(int64(template.Price))),
			table.PublicTemplates.Type.SET(sqlite.String(template.Type)),
			table.PublicTemplates.Tags.SET(sqlite.String(string(tagsJSON))),
			table.PublicTemplates.CoverImage.SET(sqlite.String(template.CoverImage)),
			table.PublicTemplates.State.SET(sqlite.Int(int64(template.State))),
			table.PublicTemplates.UpdatedAt.SET(sqlite.DATETIME(time.Now())),
		).WHERE(
		table.PublicTemplates.ID.EQ(sqlite.String(templateID)),
	)

	_, err = stmt.ExecContext(ctx, r.db.db)
	return err
}

func (r *PublicTemplate) Delete(ctx context.Context, templateID string) error {
	stmt := table.PublicTemplates.DELETE().
		WHERE(table.PublicTemplates.ID.EQ(sqlite.String(templateID)))

	result, err := stmt.ExecContext(ctx, r.db.db)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no template found with the given ID")
	}

	return nil
}
