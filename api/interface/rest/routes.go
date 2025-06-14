package rest

import (
	"net/http"

	"basic-service/interface/rest/handlers"
	"basic-service/interface/rest/model"
	"basic-service/usecase"

	httpin_integration "github.com/ggicci/httpin/integration"

	appMiddleware "basic-service/interface/rest/middleware"

	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(
	authCase *usecase.Auth,
	publicTemplateCase *usecase.PublicTemplateUseCase,
	userTemplateCase *usecase.UserTemplate,
	guestCase *usecase.GuestUsecase,
) *chi.Mux {
	r := chi.NewRouter()
	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	uploadHandler := &handlers.UploadHandler{
		UploadDir:   "./public/uploads",
		TemplateDir: "./public/template",
	}
	authHandler := handlers.NewAuthHandler(authCase, uploadHandler)
	publicTemplateHandler := handlers.NewPublicTemplate(publicTemplateCase, uploadHandler)
	userTemplateHandler := handlers.NewUserTemplate(userTemplateCase, uploadHandler)
	guestHandler := handlers.NewGuest(guestCase)

	r.Handle("/uploads/*", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./public/uploads"))))

	// Public routes
	r.Group(func(r chi.Router) {
		r.Post("/auth/login", authHandler.Login)
		r.With(httpin.NewInput(model.IdentityRequest{})).Get("/public/guest/{id}", guestHandler.GetGuest)
		r.With(httpin.NewInput(model.GuestUpdateMessageRequest{})).Post("/public/guest/message", guestHandler.UpdateMessage)
		r.With(httpin.NewInput(model.IdentityRequest{})).Put("/public/guest/{id}", guestHandler.UpdateLastView)
		r.With(httpin.NewInput(model.RegisterUser{})).Post("/auth/register", authHandler.Register)
	})

	// // Protected routes
	r.Group(func(r chi.Router) {
		// JWT verification
		r.Use(appMiddleware.AuthMiddleware(authCase))
		//
		r.Get("/auth/me", authHandler.Me)

		r.Route("/private/", func(r chi.Router) {
			// // Public Template Manager
			r.With(httpin.NewInput(model.PaginationRequest{})).Get("/public-templates", publicTemplateHandler.List)
			r.With(httpin.NewInput(model.PublicTemplateCreateRequest{})).Post("/public-templates", publicTemplateHandler.Create)

			r.With(httpin.NewInput(model.PaginationRequest{})).Get("/user-templates", userTemplateHandler.List)
			r.With(httpin.NewInput(model.UserTemplateCreateRequest{})).Post("/user-templates", userTemplateHandler.Create)

			r.Get("/guests", guestHandler.List)
			r.Post("/guests", guestHandler.Create)

			// r.Delete("/guests/{id}", guestHandler.Delete)
			// // User Manager
			// r.Get("/users", handlers.ListUsers)
			// r.Patch("/users", handlers.ChangeUserState)
			// r.Get("/users/{id}", handlers.GetUser)
			//
			//r.Get("/public-templates/{id}", handlers.GetPublicTemplate)
			// r.Put("/public-templates/{id}", handlers.UpdatePublicTemplate)
			// r.Delete("/public-templates/{id}", handlers.DeletePublicTemplate)
			//
			// // User Template Manager
			// r.Get("/user-templates/{id}", handlers.GetUserTemplate)
			// r.Put("/user-templates/{id}", handlers.UpdateUserTemplate)
			// r.Delete("/user-templates/{id}", handlers.DeleteUserTemplate)
			//
			// // Guest Manager
			// r.Put("/guests/{id}", handlers.UpdateGuest)
		})
	})

	return r
}

func init() {
	httpin_integration.UseGochiURLParam("path", chi.URLParam)
}
