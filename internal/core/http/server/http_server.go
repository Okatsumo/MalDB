package server

import (
	"MalDB/internal/core/app"
	"MalDB/internal/core/http/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func Run(app *app.App) {
	if !app.Cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())

	//RateLimiter
	clientLimiter := middleware.NewClientLimiter(app.Cfg.HttpRateLimit, app.Cfg.HttpRateLimitBurst)
	router.Use(middleware.RateLimitMiddleware(clientLimiter))

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))
	RegisterRouter(router)

	server := &http.Server{
		Addr:           app.Cfg.URL,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("API fatal error: ", err)
	}
}
