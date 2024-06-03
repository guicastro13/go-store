package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/guicastro13/go-store/config/env"
	"github.com/guicastro13/go-store/internal/handler/userhandler"
  "github.com/guicastro13/go-store/internal/handler/middleware"
)

func InitUserRoutes(router chi.Router, h userhandler.UserHandler) {
  router.Use(middleware.LoggerData)

	router.Post("/user", h.CreateUser)
	router.Route("/user", func(r chi.Router) {
    r.Use(jwtauth.Verifier(env.Env.TokenAuth))
    r.Use(jwtauth.Authenticator)
    
    r.Get("/list-all", h.FindManyUsers)
    r.Patch("/", h.UpdateUser)
	  r.Patch("/password", h.UpdateUserPassword)
	  r.Delete("/", h.DeleteUser)
	  r.Get("/", h.GetUserByID)
  })

  router.Post("/auth/login", h.Login)
}
