package server

import (
	"MalDB/internal/core/app"
	"MalDB/internal/core/http/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine, app app.App) {
	r.GET("/statistics", handlers.StatsController)
}
