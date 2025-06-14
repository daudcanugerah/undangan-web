package usecase

import (
	"context"
	"errors"
	"time"

	"basic-service/domain"
	"basic-service/interface/sql"

	"braces.dev/errtrace"
)

type UserTemplate struct {
	repo *sql.UserTemplateRepository
}

func NewUserTemplate(repo *sql.UserTemplateRepository) *UserTemplate {
	return &UserTemplate{repo: repo}
}

type UserTemplateList struct {
	Total int64
	Data  []domain.UserTemplate
}

func (p *UserTemplate) List(ctx context.Context, page, limit int) (UserTemplateList, error) {
	var result UserTemplateList

	offset := (page - 1) * limit
	// Get total count of templates
	total, err := p.repo.Count(ctx)
	if err != nil {
		return result, errtrace.Wrap(err)
	}
	result.Total = int64(total)

	// Get paginated templates
	templates, err := p.repo.List(ctx, offset, limit)
	if err != nil {
		return result, errtrace.Wrap(err)
	}
	result.Data = templates.Data

	return result, nil
}

func (p *UserTemplate) Get(ctx context.Context, id string) (domain.UserTemplate, error) {
	// First check if template exists
	exists, err := p.repo.Exists(ctx, id)
	if err != nil {
		return domain.UserTemplate{}, errtrace.Wrap(err)
	}
	if !exists {
		return domain.UserTemplate{}, errtrace.Wrap(sql.ErrUserTemplateNotFound)
	}

	// Get the template
	template, err := p.repo.Get(ctx, id)
	if err != nil {
		return domain.UserTemplate{}, errtrace.Wrap(err)
	}

	return template, nil
}

func (p *UserTemplate) Create(ctx context.Context, data domain.UserTemplate) error {
	claims, err := GetClaimFromContext(ctx)
	if err != nil {
		return errtrace.Wrap(errors.New("invalid token claims"))
	}
	now := time.Now()
	data.UserID = claims.UserID
	data.CreatedAt = now
	data.UpdatedAt = now

	return errtrace.Wrap(p.repo.Create(ctx, data))
}

func (p *UserTemplate) Update(ctx context.Context, id string, data domain.UserTemplate) error {
	// First check if template exists
	exists, err := p.repo.Exists(ctx, id)
	if err != nil {
		return errtrace.Wrap(err)
	}
	if !exists {
		return errtrace.Wrap(sql.ErrUserTemplateNotFound)
	}

	// Set updated timestamp
	data.UpdatedAt = time.Now()

	return errtrace.Wrap(p.repo.Update(ctx, id, data))
}

func (p *UserTemplate) Delete(ctx context.Context, id string) error {
	// First check if template exists
	exists, err := p.repo.Exists(ctx, id)
	if err != nil {
		return errtrace.Wrap(err)
	}
	if !exists {
		return errtrace.Wrap(sql.ErrUserTemplateNotFound)
	}

	return errtrace.Wrap(p.repo.Delete(ctx, id))
}
