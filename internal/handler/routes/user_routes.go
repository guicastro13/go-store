package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/guicastro13/go-store/internal/handler/userhandler"
)

func InitUserRoutes(router chi.Router, h userhandler.UserHandler) {
  router.Route("/user", func(r chi.Router) {
  })
}
