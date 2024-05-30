package main

import (
  "log/slog"

  "github.com/guicastro13/go-store/config/logger"
)

type user struct {
  Name string `json:"name"`
  Age int `json:"age"`
  Password string `json:"password"`
}

func (u user) LogUser() slog.Value {
  return slog.GroupValue(
    slog.String("name", u.Name),
    slog.Int("age", u.Age),
    slog.String("password", "HIDDEN"),
  )
}

func main() {
  logger.InitLogger()
  user := user{
    Name: "Guilherme",
    Age: 29,
    Password: "123456",
  }

  slog.Info("starting api")
  slog.Info("creating user", "user", user.LogUser())
}
