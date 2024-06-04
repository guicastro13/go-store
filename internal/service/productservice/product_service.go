package productservice

import (
	"context"

	"github.com/guicastro13/go-store/internal/dto"
)

func NewProductService(repo productrepository.ProductRepository) ProductService {
  return &service{
    repo,
  }
}

type service struct {
  repo productrepository.ProductRepository
}

type ProductService interface {
  CreateProduct(ctx context.Context, u dto.CreateCategoryDto) error
  UpdateProduct(ctx context.Context, id string, u dto.UpdateProductDto) error
  DeleteProduct(ctx context.Context, id string) error
  FindManyProducts(ctx context.Context, d dto.FindProductDto) ([]response.ProductResponse, error)
}
