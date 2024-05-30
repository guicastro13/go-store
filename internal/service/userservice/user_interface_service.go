package userservice

import userrepository "github.com/guicastro13/go-store/internal/repository"

func NewUserService(repo userrepository.UserRepository) UserService {
  return &service{
    repo,
  }
}

type service struct {
  repo userrepository.UserRepository
}

type UserService interface {
  CreateUser() error
}
