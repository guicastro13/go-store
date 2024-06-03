package userrepository

import (
	"context"

	"github.com/guicastro13/go-store/internal/entity"
)

func (r *repository) CreateUser() error {
	return nil
}

func (r *repository) FindUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
  userMock := entity.UserEntity{
    ID: "1",
    Name: "Guilherme",
    Email: "guii_1@hotmail.com",
  }
  return &userMock, nil
}

func (r *repository) FindUserByID(ctx context.Context, id string) (*entity.UserEntity, error) {
	return nil, nil
}

func (r *repository) UpdateUser(ctx context.Context, u *entity.UserEntity) error {
	return nil
}

func (r *repository) DeleteUser(ctx context.Context, id string) error {
	return nil
}

func (r *repository) FindManyUsers(ctx context.Context) ([]entity.UserEntity, error) {
	return nil, nil
}

func (r *repository) UpdatePassword(ctx context.Context, pass, id string) error {
	return nil
}

func (r *repository) GetUserPassword(ctx context.Context, id string) (*entity.UserEntity, error) {
  userMock := entity.UserEntity{
    ID: "1",
    Password: "$2y$12$CwjjXJGAkR4OKQeTvMo9suJ1s6PdKl9l4RZL9/yg.8cccDE8o/5sm",
  }
  return &userMock, nil
}
