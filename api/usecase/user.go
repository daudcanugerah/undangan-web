package usecase

import (
	"context"

	"basic-service/domain"
	"basic-service/interface/sql"
)

type UserUsecase struct {
	uc *sql.UserRepository
}

func NewUserUsecase(uc *sql.UserRepository) *UserUsecase {
	return &UserUsecase{
		uc: uc,
	}
}

type UserListResult struct {
	Total int
	Data  []domain.SafeUser
}

func (p *UserUsecase) List(ctx context.Context, page, limit int) (UserListResult, error) {
	var result UserListResult

	resp, err := p.uc.ListByRole(ctx, domain.RoleUser, page, limit)
	if err != nil {
		return result, err
	}

	result.Total = resp.Total
	result.Data = make([]domain.SafeUser, 0, len(resp.Data))
	for _, v := range resp.Data {
		result.Data = append(result.Data, domain.SafeUser{
			ID:        v.ID,
			Name:      v.Name,
			Email:     v.Email,
			Profile:   v.Profile,
			IsActive:  v.IsActive,
			Role:      v.Role,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return result, nil
}
