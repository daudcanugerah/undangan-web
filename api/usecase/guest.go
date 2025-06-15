package usecase

import (
	"context"
	"errors"
	"time"

	"basic-service/domain"
	"basic-service/interface/sql"

	"braces.dev/errtrace"
)

type GuestUsecase struct {
	guestRepo *sql.GuestManager
}

type GuestListResult struct {
	Total int64
	Data  []domain.Guest
}

func NewGuestUsecase(guestRepo *sql.GuestManager) *GuestUsecase {
	return &GuestUsecase{guestRepo: guestRepo}
}

func (g *GuestUsecase) GetGuest(ctx context.Context, id string) (*domain.Guest, error) {
	guest, err := g.guestRepo.Get(ctx, id)
	if err != nil {
		return guest, err
	}

	return guest, nil
}

func (g *GuestUsecase) Create(ctx context.Context, data domain.Guest) error {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	return g.guestRepo.Create(ctx, data)
}

func (g *GuestUsecase) List(ctx context.Context, userTemplateID string, page int, limit int) (GuestListResult, error) {
	claims, err := GetClaimFromContext(ctx)
	if err != nil {
		return GuestListResult{}, errtrace.Wrap(errors.New("invalid token claims"))
	}

	guests, total, err := g.guestRepo.List(ctx, claims.UserID, userTemplateID, page, limit)
	if err != nil {
		return GuestListResult{}, err
	}

	return GuestListResult{
		Total: total,
		Data:  guests,
	}, nil
}

func (g *GuestUsecase) UpdateLastView(ctx context.Context, id string) error {
	// Get the existing guest first
	guest, err := g.guestRepo.Get(ctx, id)
	if err != nil {
		return err
	}

	return g.guestRepo.UpdateMessageAndLastView(ctx, id, guest.Message, guest.Attend)
}

func (g *GuestUsecase) UpdateMessageAndLastView(ctx context.Context, id, message string, attend bool) error {
	_, err := g.guestRepo.Get(ctx, id)
	if err != nil {
		return err
	}

	return g.guestRepo.UpdateMessageAndLastView(ctx, id, message, &attend)
}
