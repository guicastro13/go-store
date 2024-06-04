package categoryservice

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/guicastro13/go-store/internal/dto"
	"github.com/guicastro13/go-store/internal/entity"
	"github.com/guicastro13/go-store/internal/handler/response"
)

func (s *service) CreateCategory(ctx context.Context, u dto.CreateCategoryDto) error {
  categoryEntity := entity.CategoryEntity{
    ID: uuid.New().String(),
    Title: u.Title,
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }
  err := s.repo.CreateCategory(ctx, &categoryEntity)
  if err != nil {
    return errors.New("error to create category")
  }
  return nil
}

func (s *service) FindManyCategories(ctx context.Context) (*response.ManyCategoryResponse, error) {
  findManyCategories, err := s.repo.FindManyCategories(ctx)
  if err != nil {
    slog.Error("error to find many categories", "err", err, slog.String("package", "categoryservice"))
    return nil, err
  }
  categories := response.ManyCategoryResponse{}
  for _, category := range findManyCategories {
    categoryResponse := response.CategoryResponse{
      ID: category.ID,
      Title: category.Title,
    }
    categories.Categories = append(categories.Categories, categoryResponse)
  }
  return &categories, nil
}
