package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/guicastro13/go-store/config/env"
	"github.com/guicastro13/go-store/internal/handler"
	"github.com/guicastro13/go-store/internal/handler/middleware"
)

func InitRoutes(router chi.Router, h handler.Handler) {
  router.Use(middleware.LoggerData)


	router.Route("/", func(r chi.Router) {
    r.Use(jwtauth.Verifier(env.Env.TokenAuth))
    r.Use(jwtauth.Authenticator)
    
    r.Get("/user/all", h.FindManyUsers)
    r.Patch("/user/me", h.UpdateUser)
	  r.Patch("/user/password", h.UpdateUserPassword)
	  r.Delete("/user/me", h.DeleteUser)
	  r.Get("/user/me", h.GetUserByID)

    //category
    r.Post("/category", h.CreateCategory)
    r.Get("/categories", h.FindManyCategories)

    //products
    r.Post("/product", h.CreateProduct)
 
    r.Get("/products", h.FindManyProducts)
    r.Patch("/product/{id}", h.UpdateProduct)
    r.Delete("/product/{id}", h.DeleteProduct)
  })

  router.Post("/user", h.CreateUser)
  router.Post("/auth/login", h.Login)
}
