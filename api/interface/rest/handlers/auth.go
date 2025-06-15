package handlers

import (
	"basic-service/domain"
	"basic-service/interface/rest/model"
	"basic-service/usecase"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/ggicci/httpin"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type AuthHandler struct {
	validator *validator.Validate
	cs        *usecase.Auth
	upload    *UploadHandler
	// Add any dependencies like userService, tokenService etc.
}

func NewAuthHandler(auth *usecase.Auth, upload *UploadHandler) *AuthHandler {
	return &AuthHandler{
		validator: validator.New(),
		cs:        auth,
		upload:    upload,
	}
}

// Register handles user registration
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// Retrieve your data in one line of code!
	input := r.Context().Value(httpin.Input).(*model.RegisterUser)
	input.Email = strings.TrimSpace(input.Email)

	imageURL, err := h.upload.UploadImage(input.Profile)
	if err != nil {
		renderError(w, r, http.StatusBadRequest, "Image upload failed", err)
		return
	}

	if err := h.validator.Struct(input); err != nil {
		renderError(w, r, http.StatusBadRequest, "Validation failed", err)
		return
	}

	if err := h.cs.Register(ctx, domain.User{
		ID:        uuid.New().String(),
		Email:     input.Email,
		Password:  input.Password,
		Name:      input.Name,
		Profile:   path.Join("uploads", imageURL),
		Role:      domain.RoleUser,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}); err != nil {
		renderError(w, r, http.StatusBadRequest, "Register failed", err)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{})
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	sf, err := h.cs.Me(r.Context())
	if err != nil {
		renderError(w, r, http.StatusBadRequest, "Get Me failed", err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, model.SafeUser{
		CreatedAt: sf.CreatedAt,
		Email:     sf.Email,
		Id:        sf.ID,
		IsActive:  sf.IsActive,
		Name:      sf.Name,
		Role:      model.UserRole(sf.Role),
		UpdatedAt: time.Time{},
	})
}

// Login handles user authentication
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		renderError(w, r, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := h.validator.Struct(req); err != nil {
		renderError(w, r, http.StatusBadRequest, "Validation failed", err)
		return
	}

	token, err := h.cs.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		renderError(w, r, http.StatusBadRequest, "Login failed", err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, model.LoginResponse{
		Token: &token,
	})
}

// Logout handles user logout
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{"message": "Successfully logged out"})
}

// renderError is a helper for consistent error responses
func renderError(w http.ResponseWriter, r *http.Request, status int, message string, err error) {
	resp := model.ErrorResponse{
		Status:  status,
		Message: message,
	}

	if err != nil {
		resp.Error = strings.ReplaceAll(fmt.Sprintf("%s: %+v", message, err), "\n", "\n")
	}

	render.Status(r, status)
	render.JSON(w, r, resp)
}
