package productrepository

import (
	"context"
	"database/sql"

	"github.com/guicastro13/go-store/internal/database/sqlc"
	"github.com/guicastro13/go-store/internal/dto"
	"github.com/guicastro13/go-store/internal/entity"
)

func NewProductRepository(db *sql.DB, q *sqlc.Queries) ProductRepository {
  return &repository{
    db,
    q,
  }
}

type repository struct {
  db *sql.DB
  queries *sqlc.Queries
}

type ProductRepository interface {
  CreateProduct(ctx context.Context, p *entity.ProductEntity, c []entity.ProductCategoryEntity) error
  GetCategoryByID(ctx context.Context, id string) (bool, error)
  GetProductByID(ctx context.Context, id string) (bool, error)
  UpdateProduct(ctx context.Context, p *entity.ProductEntity, c []entity.ProductCategoryEntity) error
  GetCategoriesByProductID(ctx context.Context, id string) ([]string, error)
  DeleteProductCategory(ctx context.Context, productID, categodyID string) error
  DeleteProduct(ctx context.Context, id string) error
  FindManyProducts(ctx context.Context, d dto.FindProductDto) ([]entity.ProductWithCategoryEntity, error)
}
