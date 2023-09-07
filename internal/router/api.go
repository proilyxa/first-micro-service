package router

import (
	"education-project/internal/http/controllers"
	"education-project/internal/http/middlewares"
	"education-project/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

func initApiRoutes(
	r *chi.Mux,
	authController *controllers.AuthController,
) {
	r.Route("/api", func(r chi.Router) {
		r.Use(render.SetContentType(render.ContentTypeJSON))

		r.Post("/register", authController.Register)
		r.Post("/login", authController.Login)

		r.Group(func(r chi.Router) {
			userService := authController.UserService.(*services.UserServiceImpl)
			r.Use(middlewares.BearerAuth(userService.UserRepository))

			r.Get("/test", func(resp http.ResponseWriter, req *http.Request) {
				user := middlewares.GetAuthUser(req.Context())
				render.JSON(resp, req, user)
			})
		})

	})
}
