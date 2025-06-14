package usecase

import (
	"context"

	"basic-service/domain"
	"basic-service/interface/sql"
)

type UserUsecase struct{}

func NewUserUsecase(guestManager sql.GuestManager) *UserUsecase {
	return &UserUsecase{}
}

type UserListResult struct {
	Total int
	Data  []domain.SafeUser
}

func (p *UserUsecase) List(ctx context.Context, page, limit int) (PublicTemplateList, error) {
	var templates PublicTemplateList
	return templates, nil
}
