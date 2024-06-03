package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/guicastro13/go-store/internal/common/utils"
)

func LoggerData(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    var requestData map[string]interface{}
    if r.Body != http.NoBody {
      CopyBody, _ := io.ReadAll(r.Body)
      r.Body = io.NopCloser(bytes.NewBuffer(CopyBody))
      if err := json.Unmarshal(CopyBody, &requestData); err != nil {
        slog.Error("error unmarshalling request data", err, slog.String("func", "LoggerData"))
      }
    } else {
      r.Body = http.NoBody
    }

    var userID string
    var userEmail string
    user, err := utils.DecodeJwt(r)
    if err != nil {
      userID = "no token"
      userEmail = "no token"
    } else {
      userID = user.ID
      userEmail = user.Email
    }
    slog.Info("request_data",
      slog.Any("url", r.URL.Path),
      slog.Any("method", r.Method),
      slog.Any("query", r.URL.Query()),
      slog.Any("body", requestData),
      slog.Any("id", userID),
      slog.Any("email", userEmail),
      )
    next.ServeHTTP(w,r)
  })
}
