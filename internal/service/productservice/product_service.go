package productservice

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

func (s *service) CreateProduct(ctx context.Context, u dto.CreateProductDto) error {
  productID := uuid.New().String()
  productEntity := entity.ProductEntity{
    ID: productID,
    Title: u.Title,
    Price: u.Price,
    Categories: u.Categories,
    Description: u.Description,
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }
  var categories []entity.ProductCategoryEntity
  for _, categoryID := range u.Categories {
  exists, err := s.repo.GetCategoryByID(ctx, categoryID)
  if err != nil || !exists {
      slog.Error("category not found", slog.String("category_id", categoryID), slog.String("package", "productservice"))
      return errors.New("category not found")
    }
    categories = append(categories, entity.ProductCategoryEntity{
      ID: uuid.New().String(),
      ProductID: productID,
      CategoryID: categoryID,
      CreatedAt: time.Now(),
      UpdatedAt: time.Now(),
    })
  }
  err := s.repo.CreateProduct(ctx, &productEntity, categories)
  if err != nil {
    return err
  }
  return nil
}

func (s *service) UpdateProduct(ctx context.Context, id string, u dto.UpdateProductDto) error {
  exists, err := s.repo.GetProductByID(ctx, id)
  if err != nil || !exists {
    slog.Error("product not found", slog.String("product_id", id), slog.String("package", "productservice"))
    return errors.New("product not found")
  }

  var categories []entity.ProductCategoryEntity
  if len(u.Categories) > 0 {
    for _, categoryID := range u.Categories {
      exists, err := s.repo.GetCategoryByID(ctx, categoryID)
      if err != nil || !exists {
        slog.Error("category not found", slog.String("category_id", categoryID), slog.String("package", "productservice"))
        return errors.New("category not found")
      }
    }

    productCategories, err := s.repo.GetCategoriesByProductID(ctx, id)
    if err != nil {
      return errors.New("error getting categories by product id")
    }

    for _, productCategory := range productCategories {
      found := false
      for _, categoryID := range u.Categories {
        if productCategory == categoryID {
          found = true
          break
        }
      }

      if !found {
        err = s.repo.DeleteProductCategory(ctx, id, productCategory)
        if err != nil {
          return errors.New("error deleting product category")
        }
      }
    }

    for _, categoryID := range u.Categories {
      found := false
      for _, productCategory := range productCategories {
        if productCategory == categoryID {
          found = true
          break
        }
      }
      if !found {
        categories = append(categories, entity.ProductCategoryEntity{
          ID: uuid.New().String(),
          ProductID: id,
          CategoryID: categoryID,
          CreatedAt: time.Now(),
          UpdatedAt: time.Now(),
        })
      }
    }
  }
  productEntity := entity.ProductEntity{
    ID: id,
    Title: u.Title,
    Price: u.Price,
    Description: u.Description,
    Categories: u.Categories,
    UpdatedAt: time.Now(),
  }
  err = s.repo.UpdateProduct(ctx, &productEntity, categories)
  if err != nil {
    return err
  }
  return nil
}

func (s *service) DeleteProduct(ctx context.Context, id string) error {
  exists, err := s.repo.GetProductByID(ctx, id)
  if err != nil || !exists {
    slog.Error("product not found", slog.String("product_id", id), slog.String("package", "productservice"))
    return errors.New("product not found")
  }
  err = s.repo.DeleteProduct(ctx, id)
  if err != nil {
    return err
  }
  return nil
}

func (s *service) FindManyProducts(ctx context.Context, d dto.FindProductDto) ([]response.ProductResponse, error) {
  products, err := s.repo.FindManyProducts(ctx, d)
  if err != nil {
    return nil, err
  }

  var productResponse []response.ProductResponse
  for _, p := range products {
    var categories []response.CategoryResponse
    for _, c := range p.Categories {
      categories = append(categories, response.CategoryResponse{
        ID: c.ID,
        Title: c.Title,
      })
    }
    productResponse = append(productResponse, response.ProductResponse{
      ID: p.ID,
      Title: p.Title,
      Description: p.Description,
      Price: p.Price,
      Categories: categories,
      CreatedAt: p.CreatedAt,
    })
  }
  if len(productResponse) == 0 {
    return []response.ProductResponse{}, nil
  }
  return productResponse, nil
}
