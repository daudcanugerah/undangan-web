package handlers

import (
	"net/http"
	"path"

	"basic-service/domain"
	"basic-service/interface/rest/model"
	"basic-service/usecase"

	"github.com/ggicci/httpin"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type PublicTemplate struct {
	validator *validator.Validate
	cs        *usecase.PublicTemplateUseCase
	upload    *UploadHandler
	// Add any dependencies like userService, tokenService etc.
}

func NewPublicTemplate(cs *usecase.PublicTemplateUseCase, upload *UploadHandler) *PublicTemplate {
	return &PublicTemplate{
		validator: validator.New(),
		cs:        cs,
		upload:    upload,
	}
}

func (h *PublicTemplate) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// Retrieve your data in one line of code!
	input := r.Context().Value(httpin.Input).(*model.PublicTemplateCreateRequest)
	templateSlug := uuid.New().String()
	err := h.upload.UploadTemplate(input.ZipFile, templateSlug)
	if err != nil {
		renderError(w, r, http.StatusBadRequest, "Template upload failed", err)
		return
	}

	// input.Email = strings.TrimSpace(input.Email)
	//
	coverURL, err := h.upload.UploadImage(input.CoverImage)
	if err != nil {
		renderError(w, r, http.StatusBadRequest, "Cover Image upload failed", err)
		return
	}

	if err := h.validator.Struct(input); err != nil {
		renderError(w, r, http.StatusBadRequest, "Validation failed", err)
		return
	}

	if err := h.cs.Create(ctx, domain.PublicTemplate{
		ID:            uuid.New().String(),
		Name:          input.Name,
		Description:   input.Description,
		PriceInterval: input.PriceInterval,
		Price:         input.Price,
		Type:          input.Type,
		Tags:          input.Tags,
		CoverImage:    path.Join("uploads", coverURL),
		State:         input.State,
		Slug:          templateSlug,
	}); err != nil {
		renderError(w, r, http.StatusBadRequest, "Create Public Template failed", err)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{})
}

func (h *PublicTemplate) List(w http.ResponseWriter, r *http.Request) {
	input := r.Context().Value(httpin.Input).(*model.PaginationRequest)

	data, err := h.cs.List(r.Context(), input.Page, input.Limit)
	if err != nil {
		renderError(w, r, http.StatusBadRequest, "get template list error", err)
		return
	}

	result := model.PublicTemplateListResult{Total: int(data.Total)}
	for _, v := range data.Data {
		result.PublicTemplates = append(result.PublicTemplates, model.PublicTemplate{
			CoverImage:    v.CoverImage,
			Description:   v.Description,
			Id:            v.ID,
			Name:          v.Name,
			Price:         v.Price,
			PriceInterval: v.PriceInterval,
			State:         v.State,
			Tags:          v.Tags,
			Type:          v.Type,
			CreatedAt:     v.CreatedAt,
			Slug:          v.Slug,
			UpdatedAt:     v.UpdatedAt,
		})
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, result)
}
