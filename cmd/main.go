package main

import (
  "log/slog"

  "github.com/guicastro13/go-store/config/logger"
)

func main() {
  logger.InitLogger()

  slog.Info("starting api")
}
