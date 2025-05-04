package main

import (
	"MalDB/internal/config"
	"MalDB/internal/core/app"
	"MalDB/internal/core/http/server"
	"fmt"
)

func main() {
	cfg := config.Load()
	application := app.Init(cfg)

	server.Run(application)

	//Другие части аппы наверное запускать здесь, либо в слое APP...

	fmt.Print(cfg)
}
