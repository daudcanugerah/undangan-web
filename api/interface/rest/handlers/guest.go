package handlers

import (
	"basic-service/domain"
	"basic-service/interface/rest/model"
	"basic-service/usecase"
	"net/http"

	"github.com/ggicci/httpin"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Guest struct {
	validator *validator.Validate
	cs        *usecase.GuestUsecase
}

func NewGuest(cs *usecase.GuestUsecase) *Guest {
	return &Guest{
		validator: validator.New(),
		cs:        cs,
	}
}

func (h *Guest) UpdateLastView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// Retrieve your data in one line of code!
	input := r.Context().Value(httpin.Input).(*model.IdentityRequest)

	if err := h.validator.Struct(input); err != nil {
		renderError(w, r, http.StatusBadRequest, "Validation failed", err)
		return
	}

	if err := h.cs.UpdateLastView(ctx, input.ID); err != nil {
		renderError(w, r, http.StatusBadRequest, "Update Guest Last View failed", err)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{})
}

func (h *Guest) GetGuest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// Retrieve your data in one line of code!
	input := r.Context().Value(httpin.Input).(*model.IdentityRequest)

	if err := h.validator.Struct(input); err != nil {
		renderError(w, r, http.StatusBadRequest, "Validation failed", err)
		return
	}

	guest, err := h.cs.GetGuest(ctx, input.ID)
	if err != nil {
		renderError(w, r, http.StatusBadRequest, "Get Guest Message failed", err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, model.SafeGuest{
		Name:    guest.Name,
		Group:   guest.Group,
		Person:  guest.Person,
		Address: guest.Address,
		Message: guest.Message,
		ViewAt:  guest.ViewAt,
		Attend:  guest.Attend,
	})
}

func (h *Guest) UpdateMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// Retrieve your data in one line of code!
	input := r.Context().Value(httpin.Input).(*model.GuestUpdateMessageRequest)

	if err := h.validator.Struct(input); err != nil {
		renderError(w, r, http.StatusBadRequest, "Validation failed", err)
		return
	}

	if err := h.cs.UpdateMessageAndLastView(ctx, input.Payload.ID, input.Payload.Message, input.Payload.Attend); err != nil {
		renderError(w, r, http.StatusBadRequest, "update Guest Message failed", err)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{})
}

func (h *Guest) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// Retrieve your data in one line of code!
	input := r.Context().Value(httpin.Input).(*model.GuestCreateRequest)

	if err := h.validator.Struct(input); err != nil {
		renderError(w, r, http.StatusBadRequest, "Validation failed", err)
		return
	}

	if err := h.cs.Create(ctx, domain.Guest{
		ID:             uuid.New().String(),
		UserTemplateID: input.Payload.UserTemplateId,
		Name:           input.Payload.Name,
		Group:          input.Payload.Group,
		Person:         input.Payload.Person,
		Tags:           input.Payload.Tags,
		Telp:           input.Payload.Telp,
		Address:        input.Payload.Address,
	}); err != nil {
		renderError(w, r, http.StatusBadRequest, "Create Guest failed", err)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{})
}

func (h *Guest) List(w http.ResponseWriter, r *http.Request) {
	input := r.Context().Value(httpin.Input).(*model.GuestListRequest)

	data, err := h.cs.List(r.Context(), input.UserTemplateID, input.Page, input.Limit)
	if err != nil {
		renderError(w, r, http.StatusBadRequest, "get template list error", err)
		return
	}

	result := model.GuestListResult{Total: int(data.Total)}
	for _, v := range data.Data {
		k := model.Guest{
			Address:        v.Address,
			Attend:         v.Attend,
			CreatedAt:      v.CreatedAt,
			Group:          v.Group,
			Id:             v.ID,
			Message:        v.Message,
			Name:           v.Name,
			Person:         v.Person,
			Tags:           v.Tags,
			Telp:           v.Telp,
			UpdatedAt:      v.UpdatedAt,
			UserTemplateId: v.UserTemplateID,
		}

		if v.ViewAt != nil && !v.ViewAt.IsZero() {
			k.ViewAt = v.ViewAt
		}
		result.Data = append(result.Data, k)
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, result)
}
