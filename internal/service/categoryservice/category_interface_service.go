package categoryservice

import (
	"context"

	"github.com/guicastro13/go-store/internal/dto"
	"github.com/guicastro13/go-store/internal/handler/response"
	"github.com/guicastro13/go-store/internal/repository/categoryrepository"
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
  FindManyCategories(ctx context.Context) (*response.ManyCategoryResponse, error)
}
