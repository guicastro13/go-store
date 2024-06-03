package userservice

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/google/uuid"
	viacep "github.com/guicastro13/go-store/api"
	"github.com/guicastro13/go-store/internal/dto"
	"github.com/guicastro13/go-store/internal/entity"
	"github.com/guicastro13/go-store/internal/handler/response"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) CreateUser(ctx context.Context, u dto.CreateUserDto) error {
  userExists, err := s.repo.FindUserByEmail(ctx, u.Email)
  if err != nil {
    slog.Error("error to search user by email", "err", err, slog.String("package", "userservice"))
    return err
  }
  if userExists != nil {
    slog.Error("user already exists", slog.String("package", "userservice"))
    return errors.New("user already exists")
  }
  passwordEcrypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), 11)
  if err != nil {
    slog.Error("error to encrypt password", "err", err, slog.String("package", "userservice"))
    return errors.New("error to escrypt password")
  }
  cep, err := viacep.GetCep(u.CEP)
  if err != nil {
    slog.Error("error to get cep", "err", err, slog.String("package", "userservice"))
    return err
  }
  newUser := entity.UserEntity{
    ID: uuid.New().String(),
    Name: u.Name,
    Email: u.Email,
    Password: string(passwordEcrypted),
    Address: entity.UserAddress{
      CEP: cep.CEP,
      IBGE: cep.IBGE,
      UF: cep.UF,
      City: cep.Localidade,
      Complement: cep.Complemento,
      Street: cep.Logradouro,
    },
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }
  err = s.repo.CreateUser(ctx, &newUser)
  if err != nil {
    slog.Error("error to create user", "err", err, slog.String("package", "userservice"))
    return err
  }
  return nil
}

func (s *service) UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error {
  var updateUser entity.UserEntity
  if u.Email != "" {
    verifyUserEmail, err := s.repo.FindUserByEmail(ctx, u.Email)
    if err != nil {
      slog.Error("error to search user by email", "err", err, slog.String("package", "userservice"))
      return err
    }
    if verifyUserEmail != nil {
      slog.Error("user already exists", slog.String("package", "userservice"))
      return errors.New("user already exists")
    }
    updateUser.Email = u.Email
  }
  if u.CEP != "" {
    cep, err := viacep.GetCep(u.CEP)
    if err != nil {
      slog.Error("error to get cep", "err", err, slog.String("package", "userservice"))
      return err
    }
    updateUser.Address = entity.UserAddress{
      CEP:        cep.CEP,
      IBGE:       cep.IBGE,
      UF:         cep.UF,
      City:       cep.Localidade,
      Complement: cep.Complemento,
      Street:     cep.Logradouro,
    }
  }
  return nil
}

func (s *service) GetUserByID(ctx context.Context, id string) (*response.UserResponse, error) {
  userExists, err := s.repo.FindUserByID(ctx, id)
  if err != nil {
    slog.Error("error to search user by id", "err", err, slog.String("package", "userservice"))
    return nil, err
  }
  if userExists == nil {
    slog.Error("user not found", slog.String("package", "userservice"))
    return nil, errors.New("user no found")
  }
  user := response.UserResponse{
    ID: userExists.ID,
    Name: userExists.Name,
    Email: userExists.Email,
    CreatedAt: userExists.CreatedAt,
    UpdatedAt: userExists.UpdatedAt,
  }
  return &user, nil
}

func (s *service) DeleteUser(ctx context.Context, id string) error {
  userExists, err := s.repo.FindUserByID(ctx, id)
  if err != nil {
    slog.Error("error to seach user by id", "err", err, slog.String("package", "userservice"))
    return err
  }
  if userExists == nil {
    slog.Error("user not found", slog.String("package", "userservice"))
    return errors.New("user not found")
  }
  err = s.repo.DeleteUser(ctx, id)
  if err != nil {
    slog.Error("error to delete user", "err", err, slog.String("package", "userservice"))
    return err
  }
  return nil
}

func (s *service) FindManyUsers(ctx context.Context) (*response.ManyUsersReponse, error) {
  findManyUsers, err := s.repo.FindManyUsers(ctx)
  if err != nil {
    slog.Error("error to find many users", "err", err, slog.String("package", "userservice"))
    return nil, err
  }
  users := response.ManyUsersReponse{}
  for _, user := range findManyUsers{
    userResponse := response.UserResponse{
      ID: user.ID,
      Name: user.Name,
      Email: user.Email,
      CreatedAt: user.CreatedAt,
      UpdatedAt: user.UpdatedAt,
    }
    users.Users = append(users.Users, userResponse)
  }
  return &users, nil
}

func (s *service) UpdateUserPassword(ctx context.Context, u *dto.UpdateUserPasswordDto, id string) error {
  userExists, err := s.repo.FindUserByID(ctx, id)
  if err != nil {
    slog.Error("error to search user by id", "err", err, slog.String("package", "userservice"))
    return err
  }
  if userExists == nil {
    slog.Error("user not found", slog.String("package", "userservice"))
    return errors.New("user not found")
  }
  err = bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(u.OldPassword))
  if err != nil {
    slog.Error("invalid password", slog.String("package", "userservice"))
    return errors.New("invalid password")
  }
  err = bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(u.Password))
  if err == nil {
    slog.Error("new password is equal to old password", slog.String("package", "userservice"))
    return errors.New("new password is equal to old password")
  }
  passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
  if err != nil {
    slog.Error("error to encrypt password", "err", err, slog.String("package", "userservice"))
    return errors.New("error to encrypt password")
  }
  err = s.repo.UpdatePassword(ctx, string(passwordEncrypted), id)
  if err != nil {
    slog.Error("error to update password", "err", err, slog.String("package", "userservice"))
    return err
  }
  return nil
}

