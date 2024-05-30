package userhandler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/guicastro13/go-store/internal/dto"
	"github.com/guicastro13/go-store/internal/handler/httperr"
	"github.com/guicastro13/go-store/internal/handler/validation"
)

func(h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
  var req dto.CreateUserDto
  if r.Body == http.NoBody {
    slog.Error("body is empty", slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("body is required")
    json.NewEncoder(w).Encode(msg)
    return
  }
  err := json.NewDecoder(r.Body).Decode(&req)
  if err != nil {
    slog.Error("error to decode body", "err", err, slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("error to decode body")
    json.NewEncoder(w).Encode(msg)
    return
  }
  httpErr := validation.ValidateHttpData(req)
  if httpErr != nil {
    slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "userhandler"))
    w.WriteHeader(httpErr.Code)
    json.NewEncoder(w).Encode(httpErr)
    return
  }
  err = h.service.CreateUser(r.Context(), req)
  if err != nil {
    slog.Error(fmt.Sprintf("error to create user: %v", err), slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusInternalServerError)
    msg := httperr.NewBadRequestError("error to create user")
    json.NewEncoder(w).Encode(msg)
    return 
  }
}

func(h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
  var req dto.UpdateUserDto
  id := chi.URLParam(r, "id")
  if id == "" {
    slog.Error("id is empty", slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("id is required")
    json.NewEncoder(w).Encode(msg)
    return
  }
  _, err := uuid.Parse(id)
  if err != nil {
    slog.Error(fmt.Sprintf("error to parse id: %v", err), slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("error to parse id")
    json.NewEncoder(w).Encode(msg)
    return
  }
  if r.Body == http.NoBody {
    slog.Error("body is empty", slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("body is required")
    json.NewEncoder(w).Encode(msg)
    return
  }
  err = json.NewDecoder(r.Body).Decode(&req)
  if err != nil {
    slog.Error("error to decode body", "err", err, slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("error to decode body")
    json.NewEncoder(w).Encode(msg)
    return
  }
  httpErr := validation.ValidateHttpData(req)
  if httpErr != nil {
    slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "userhandler"))
    w.WriteHeader(httpErr.Code)
    json.NewEncoder(w).Encode(httpErr)
    return
  }
  err = h.service.UpdateUser(r.Context(), req, id)
  if err != nil {
    slog.Error(fmt.Sprintf("error to update user: %v", err), slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusInternalServerError)
    msg := httperr.NewBadRequestError("error to update user")
    json.NewEncoder(w).Encode(msg)
    return
  }
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
  id := chi.URLParam(r, "id")
  if id == "" {
    slog.Error("id is empty", slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("id is required")
    json.NewEncoder(w).Encode(msg)
    return
  }
  _, err := uuid.Parse(id)
  if err != nil {
    slog.Error(fmt.Sprintf("error to parse id: %v", err), slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("error to parse id")
    json.NewEncoder(w).Encode(msg)
    return
  }
  res, err := h.service.GetUserByID(r.Context(), id)
  if err != nil {
    slog.Error(fmt.Sprintf("error to get user: %v", err), slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusInternalServerError)
    msg := httperr.NewBadRequestError("error to get user")
    json.NewEncoder(w).Encode(msg)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(res)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
  id := chi.URLParam(r, "id")
  if id == "" {
    slog.Error("id is empty", slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("id is required")
    json.NewEncoder(w).Encode(msg)
    return
  }
  _, err := uuid.Parse(id)
  if err != nil {
    slog.Error(fmt.Sprintf("error to parse id: %v", err), slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("error to parse id")
    json.NewEncoder(w).Encode(msg)
    return
  }
  err = h.service.DeleteUser(r.Context(), id)
  if err != nil {
    slog.Error(fmt.Sprintf("error to delete user: %v", err), slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusInternalServerError)
    msg := httperr.NewBadRequestError("error to delete user")
    json.NewEncoder(w).Encode(msg)
    return
  }
  w.WriteHeader(http.StatusNoContent)
}

func (h *handler) FindManyUsers(w http.ResponseWriter, r *http.Request) {
  res, err := h.service.FindManyUsers(r.Context())
  if err != nil {
    slog.Error(fmt.Sprintf("error to find many users: %v", err), slog.String("package", "userhandler"))
    w.WriteHeader(http.StatusInternalServerError)
    msg := httperr.NewBadRequestError("error to find many users")
    json.NewEncoder(w).Encode(msg)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(res)
}
