package userrepository

import (
	"context"
	"database/sql"

	"github.com/guicastro13/go-store/internal/database/sqlc"
	"github.com/guicastro13/go-store/internal/entity"
)

func NewUserRepository(db *sql.DB, q *sqlc.Queries) UserRepository {
	return &repository{
		db,
		q,
	}
}

type repository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

type UserRepository interface {
	CreateUser(ctx context.Context, u *entity.UserEntity) error
	FindUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
	FindUserByID(ctx context.Context, id string) (*entity.UserEntity, error)
	UpdateUser(ctx context.Context, u *entity.UserEntity) error
	DeleteUser(ctx context.Context, id string) error
	FindManyUsers(ctx context.Context) ([]entity.UserEntity, error)
	UpdatePassword(ctx context.Context, pass, id string) error
  GetUserPassword(ctx context.Context, id string) (string, error)
}
