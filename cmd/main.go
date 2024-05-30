package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/guicastro13/go-store/config"
	"github.com/guicastro13/go-store/config/env"
	"github.com/guicastro13/go-store/internal/database"
	"github.com/guicastro13/go-store/internal/database/sqlc"
	"github.com/guicastro13/go-store/internal/handler/userhandler"
	"github.com/guicastro13/go-store/internal/handler/routes"
	"github.com/guicastro13/go-store/internal/repository/userrepository"
	"github.com/guicastro13/go-store/internal/service/userservice"
)

func main() {
  logger.InitLogger()
  slog.Info("starting api")

  _, err := env.LoadingConfig(".")
  if err != nil {
    slog.Error("faild to load enviroment variable", err, slog.String("package", "main"))
    return
  }
  dbConnection, err := database.NewDBConnection()
  if err != nil {
    slog.Error("error to connect to database", "err", err, slog.String("package", "main"))
    return
  }

  router := chi.NewRouter()
  queries := sqlc.New(dbConnection)

  //user
  userRepo := userrepository.NewUserRepository(dbConnection, queries)
  newUserService := userservice.NewUserService(userRepo)
  newUserHandler := userhandler.NewUserHandler(newUserService)

  //init routes
  routes.InitUserRoutes(router, newUserHandler)

  port := fmt.Sprintf(":%s", env.Env.GoPort)
  slog.Info(fmt.Sprintf("server is running at port %s", port))
  err = http.ListenAndServe(port, router)
  if err != nil {
    slog.Error("error to start server", err, slog.String("package", "main"))
  }
}
