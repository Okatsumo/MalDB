package server

import (
	"MalDB/internal/config"
	"MalDB/internal/core/http/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine, config config.Config) {
	r.GET("/statistics", handlers.StatsController)
}
