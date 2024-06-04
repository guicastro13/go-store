package handler

import (
	"net/http"

	"github.com/guicastro13/go-store/internal/service/categoryservice"
	"github.com/guicastro13/go-store/internal/service/productservice"
	"github.com/guicastro13/go-store/internal/service/userservice"
)

func NewHandler(
  userService userservice.UserService,
  categoryService categoryservice.CategoryService,
  productService productservice.ProductService,
) Handler {
	return &handler{
    userService: userService,
    categoryService: categoryService,
    productservice: productService,
	}
}

type handler struct {
	userService userservice.UserService
  categoryService categoryservice.CategoryService
  productservice productservice.ProductService
}

type Handler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	FindManyUsers(w http.ResponseWriter, r *http.Request)
	UpdateUserPassword(w http.ResponseWriter, r *http.Request)
  Login(w http.ResponseWriter, r *http.Request)

  CreateCategory(w http.ResponseWriter, r *http.Request)
  FindManyCategories(w http.ResponseWriter, r *http.Request)

  CreateProduct(w http.ResponseWriter, r *http.Request)
  UpdateProduct(w http.ResponseWriter, r *http.Request)
  DeleteProduct(w http.ResponseWriter, r *http.Request)
  FindManyProducts(w http.ResponseWriter, r *http.Request)
}
