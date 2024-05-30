package userservice

import (
	"context"

	"github.com/guicastro13/go-store/internal/dto"
	userrepository "github.com/guicastro13/go-store/internal/repository/userrepository"
)

func NewUserService(repo userrepository.UserRepository) UserService {
  return &service{
    repo,
  }
}

type service struct {
  repo userrepository.UserRepository
}

type UserService interface {
  CreateUser(ctx context.Context, u dto.CreateUserDto) error
}
