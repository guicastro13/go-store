package userservice

import (
	"context"
	"time"

	"github.com/guicastro13/go-store/internal/dto"
	"github.com/guicastro13/go-store/internal/handler/response"
)

func (s *service) CreateUser(ctx context.Context, u dto.CreateUserDto) error {
  return nil
}

func (s *service) UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error {
  return nil
}

func (s *service) GetUserByID(ctx context.Context, id string) (*response.UserResponse, error) {
  userFake := response.UserResponse{
    ID: "123",
    Name: "Guilherme",
    Email: "guii_1@hotmail.com",
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }
  return &userFake, nil
}

func (s *service) DeleteUser(ctx context.Context, id string) error {
  return nil
}
