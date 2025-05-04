package app

import "MalDB/internal/config"

type App struct {
	Cfg config.Config
}

func Init(cfg *config.Config) *App {
	return &App{Cfg: *cfg}
}
