package userservice

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/guicastro13/go-store/config/env"
	"github.com/guicastro13/go-store/internal/dto"
	"github.com/guicastro13/go-store/internal/handler/response"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, u dto.LoginDto) (*response.UserAuthToken, error) {
  user, err := s.repo.FindUserByEmail(ctx, u.Email)
  if err != nil {
    slog.Error("error to search user by email", "err", err, slog.String("package", "userservice"))
    return nil, errors.New("error to search user password")
  }
  if user == nil {
    slog.Error("user not found", slog.String("package", "userservice"))
    return nil, errors.New("user not found")
  }
  userPass, err := s.repo.GetUserPassword(ctx, user.ID)
  if err != nil {
    slog.Error("error to search user password", "err", err, slog.String("package", "userservice"))
    return nil, errors.New("error to search user password")
  }
  err = bcrypt.CompareHashAndPassword([]byte(userPass.Password), []byte(u.Password))
  if err != nil {
    slog.Error("invalid password", slog.String("package", "userservice"))
    return nil, errors.New("invalid password")
  }
  _, token, _ := env.Env.TokenAuth.Encode(map[string]interface{}{
    "id": user.ID,
    "email": u.Email,
    "name": user.Name,
    "exp": time.Now().Add(time.Second * time.Duration(env.Env.JwtExpiresIn)).Unix(),
  })
  userAuthToken := response.UserAuthToken{
    AccessToken: token,
  }
  return &userAuthToken, nil
}
