package server

import (
	"MalDB/internal/core/http/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	r.GET("/statistics", handlers.StatsController)
}
