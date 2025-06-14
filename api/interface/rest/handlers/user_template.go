package handlers

import (
	"net/http"

	"basic-service/domain"
	"basic-service/interface/rest/model"
	"basic-service/usecase"

	"github.com/ggicci/httpin"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserTemplate struct {
	validator *validator.Validate
	cs        *usecase.UserTemplate
	upload    *UploadHandler
	// Add any dependencies like userService, tokenService etc.
}

func NewUserTemplate(cs *usecase.UserTemplate, upload *UploadHandler) *UserTemplate {
	return &UserTemplate{
		validator: validator.New(),
		cs:        cs,
		upload:    upload,
	}
}

func (h *UserTemplate) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// Retrieve your data in one line of code!
	input := r.Context().Value(httpin.Input).(*model.UserTemplateCreateRequest)

	err := h.upload.UploadTemplate(input.ZipFile, input.Slug)
	if err != nil {
		renderError(w, r, http.StatusBadRequest, "Template upload failed", err)
		return
	}

	coverURL, err := h.upload.UploadImage(input.CoverImage)
	if err != nil {
		renderError(w, r, http.StatusBadRequest, "Cover Image upload failed", err)
		return
	}

	if err := h.validator.Struct(input); err != nil {
		renderError(w, r, http.StatusBadRequest, "Validation failed", err)
		return
	}

	msgTemplate, err := input.GetMessageTemplate()
	if err != nil {
		renderError(w, r, http.StatusBadRequest, "Validation failed", err)
		return
	}

	messageTemplate := make([]domain.MessageTemplate, 0)
	for _, v := range msgTemplate {
		messageTemplate = append(messageTemplate, domain.MessageTemplate{
			Text:     v.Text,
			Provider: v.Provider,
		})
	}

	if err := h.cs.Create(ctx, domain.UserTemplate{
		ID:              uuid.New().String(),
		Name:            input.Name,
		CoverImage:      coverURL,
		State:           1,
		Slug:            input.Slug,
		BaseTemplateID:  input.BaseTemplateId,
		URL:             input.URL,
		MessageTemplate: messageTemplate,
		ExpireAt:        input.ExpireAt,
	}); err != nil {
		renderError(w, r, http.StatusBadRequest, "Create User Template failed", err)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{})
}

func (h *UserTemplate) List(w http.ResponseWriter, r *http.Request) {
	input := r.Context().Value(httpin.Input).(*model.PaginationRequest)

	data, err := h.cs.List(r.Context(), input.Page, input.Limit)
	if err != nil {
		renderError(w, r, http.StatusBadRequest, "get template list error", err)
		return
	}

	result := model.UserTemplateListResult{Total: int(data.Total)}
	for _, v := range data.Data {
		msgTemplate := make(map[string]model.MessageTemplate, 0)
		for _, x := range v.MessageTemplate {
			msgTemplate[x.Provider] = model.MessageTemplate{
				Text:     x.Text,
				Provider: x.Provider,
			}
		}
		result.Data = append(result.Data, model.UserTemplate{
			BaseTemplateId:  v.BaseTemplateID,
			CoverImage:      v.CoverImage,
			CreatedAt:       v.CreatedAt,
			ExpireAt:        v.ExpireAt,
			Id:              v.ID,
			MessageTemplate: msgTemplate,
			Name:            v.Name,
			Slug:            v.Slug,
			State:           v.State,
			UpdatedAt:       v.UpdatedAt,
			Url:             v.URL,
			UserId:          v.UserID,
		})
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, result)
}
