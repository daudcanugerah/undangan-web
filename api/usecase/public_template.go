package usecase

import (
	"context"
	"time"

	"basic-service/domain"
	"basic-service/interface/sql"
)

type PublicTemplateUseCase struct {
	repo *sql.PublicTemplate
}

func NewPublicTemplateUseCase(repo *sql.PublicTemplate) *PublicTemplateUseCase {
	return &PublicTemplateUseCase{repo: repo}
}

type PublicTemplateList struct {
	Total int64
	Data  []domain.PublicTemplate
}

func (p *PublicTemplateUseCase) List(ctx context.Context, page, limit int) (PublicTemplateList, error) {
	var result PublicTemplateList

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10 // default limit
	}
	offset := (page - 1) * limit

	// Get total count of templates
	total, err := p.repo.Count(ctx)
	if err != nil {
		return result, err
	}
	result.Total = total

	// Get paginated templates
	templates, err := p.repo.List(ctx, limit, offset)
	if err != nil {
		return result, err
	}
	result.Data = templates

	return result, nil
}

func (p *PublicTemplateUseCase) Get(ctx context.Context, id string) (domain.PublicTemplate, error) {
	return domain.PublicTemplate{}, nil
}

func (p *PublicTemplateUseCase) Create(ctx context.Context, data domain.PublicTemplate) error {
	now := time.Now()
	data.CreatedAt = now
	data.UpdatedAt = now
	return p.repo.Create(ctx, data)
}
