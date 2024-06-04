package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/guicastro13/go-store/internal/dto"
	"github.com/guicastro13/go-store/internal/handler/httperr"
	"github.com/guicastro13/go-store/internal/handler/validation"
)

func (h *handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
  var req dto.CreateCategoryDto

  if r.Body == http.NoBody {
    slog.Error("body is empty", slog.String("package", "categoryhandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewNotFoundError("body is required")
    json.NewEncoder(w).Encode(msg)
    return
  }
  err := json.NewDecoder(r.Body).Decode(&req)
  if err != nil {
    slog.Error("error to decode body", "err", err, slog.String("package", "categoryhandler"))
    w.WriteHeader(http.StatusBadRequest)
    msg := httperr.NewBadRequestError("error to decode body")
    json.NewEncoder(w).Encode(msg)
    return
  }
  httpErr := validation.ValidateHttpData(req)
  if httpErr != nil {
    slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "categoryhandler"))
    w.WriteHeader(httpErr.Code)
    json.NewEncoder(w).Encode(httpErr)
    return
  }
  err = h.categoryService.CreateCategory(r.Context(), req)
  if err != nil {
    slog.Error(fmt.Sprintf("error to create category: %v", err), slog.String("package", "categoryhandler"))
    w.WriteHeader(http.StatusBadRequest)
  }
  w.WriteHeader(http.StatusCreated)
}
