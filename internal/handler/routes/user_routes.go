package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/guicastro13/go-store/config/env"
	"github.com/guicastro13/go-store/internal/handler"
	"github.com/guicastro13/go-store/internal/handler/middleware"
)

func InitRoutes(router chi.Router, h handler.Handler) {
  router.Use(middleware.LoggerData)


	router.Route("/user", func(r chi.Router) {
    r.Use(jwtauth.Verifier(env.Env.TokenAuth))
    r.Use(jwtauth.Authenticator)
    
    r.Get("/all", h.FindManyUsers)
    r.Patch("/me", h.UpdateUser)
	  r.Patch("/password", h.UpdateUserPassword)
	  r.Delete("/me", h.DeleteUser)
	  r.Get("/me", h.GetUserByID)
  })

  router.Post("/user", h.CreateUser)
  router.Post("/auth/login", h.Login)
}
