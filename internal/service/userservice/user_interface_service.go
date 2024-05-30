package userservice

import (
	"context"

	"github.com/guicastro13/go-store/internal/dto"
	"github.com/guicastro13/go-store/internal/handler/response"
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
  UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error
  GetUserByID(ctx context.Context, id string) (*response.UserResponse, error)
  DeleteUser(ctx context.Context, id string) error
  FindManyUsers(ctx context.Context) (response.ManyUsersReponse, error)
}
