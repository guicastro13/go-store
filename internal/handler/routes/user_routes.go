package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/guicastro13/go-store/internal/handler/userhandler"
)

func InitUserRoutes(router chi.Router, h userhandler.UserHandler) {
  router.Post("/user", h.CreateUser)
  router.Patch("/user/{id}", h.UpdateUser)
  router.Get("/user/{id}", h.GetUserByID)
  router.Delete("/user/{id}", h.DeleteUser)
  router.Get("/users", h.FindManyUsers)
  router.Patch("/password/{id}", h.UpdateUserPassword)
}
