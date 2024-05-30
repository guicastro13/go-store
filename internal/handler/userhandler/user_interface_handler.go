package userhandler

import (
	"net/http"

	userservice "github.com/guicastro13/go-store/internal/service/userservice"
)

func NewUserHandler(service userservice.UserService) UserHandler {
  return &handler{
    service,
  }
}

type handler struct {
  service userservice.UserService
}

type UserHandler interface {
  CreateUser(w http.ResponseWriter, r *http.Request)
  UpdateUser(w http.ResponseWriter, r *http.Request)
  GetUserByID(w http.ResponseWriter, r *http.Request)
}

