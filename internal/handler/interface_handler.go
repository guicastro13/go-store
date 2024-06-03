package handler

import (
	"net/http"

	userservice "github.com/guicastro13/go-store/internal/service/userservice"
)

func NewHandler(service userservice.UserService) Handler {
	return &handler{
		service,
	}
}

type handler struct {
	service userservice.UserService
}

type Handler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	FindManyUsers(w http.ResponseWriter, r *http.Request)
	UpdateUserPassword(w http.ResponseWriter, r *http.Request)
  Login(w http.ResponseWriter, r *http.Request)
}
