package categoryservice

import (
	"context"

	"github.com/guicastro13/go-store/internal/dto"
)


func NewCategoryService(repo categoryrepository.CategoryRepository) CategoryService {
  return &service{
    repo,
  }
}

type service struct {
  repo categoryrepository.CategoryRepository
}

type CategoryService interface {
  CreateCategory(ctx context.Context, u dto.CreateCategoryDto) error
}
