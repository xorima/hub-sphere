package app

import (
	"log/slog"

	"github.com/xorima/hub-sphere/internal/config"
)

type App struct {
	config *config.HubSphere
	log    *slog.Logger
}

func NewApp(log *slog.Logger, cfg *config.HubSphere) *App {
	return &App{log: log, config: cfg}
}
