package handlers

import (
	"fmt"
	"net/http"

	"basic-service/interface/rest/model"
	"basic-service/usecase"

	"github.com/ggicci/httpin"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	validator *validator.Validate
	cs        *usecase.UserUsecase
}

func NewUserHandler(cs *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		validator: validator.New(),
		cs:        cs,
	}
}

func (h *UserHandler) ListUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// Retrieve your data in one line of code!
	input := r.Context().Value(httpin.Input).(*model.PaginationRequest)

	if err := h.validator.Struct(input); err != nil {
		renderError(w, r, http.StatusBadRequest, "Validation failed", err)
		return
	}

	u, err := h.cs.List(ctx, input.Page, input.Limit)
	if err != nil {
		renderError(w, r, http.StatusBadRequest, "Get UserHandler Message failed", err)
		return
	}

	var response []model.SafeUser

	for _, v := range u.Data {
		fmt.Printf("User: %+v\n", v)
		response = append(response, model.SafeUser{
			CreatedAt: v.CreatedAt,
			Email:     v.Email,
			Id:        v.ID,
			IsActive:  v.IsActive,
			Name:      v.Name,
			Role:      model.UserRole(v.Role),
			UpdatedAt: v.UpdatedAt,
		})
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]any{
		"total": u.Total,
		"data":  response,
	})
}
