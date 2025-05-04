package handlers

import "github.com/gin-gonic/gin"

func StatsController(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}
