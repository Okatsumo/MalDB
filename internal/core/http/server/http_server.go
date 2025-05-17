package server

import (
	"MalDB/internal/core/app"
	"MalDB/internal/core/http/middleware"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func Run(app *app.App) {
	if !app.Config().Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())

	//Rate limiter
	clientLimiter := middleware.NewClientLimiter(app.Config().HttpRateLimit, app.Config().HttpRateLimitBurst)
	router.Use(middleware.RateLimitMiddleware(clientLimiter))

	// Request logger
	router.Use(middleware.RequestLogger())

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOrigins:     []string{app.Config().URL},
		MaxAge:           12 * time.Hour,
	}))
	RegisterRouter(router, app.Config())

	server := &http.Server{
		Addr:           fmt.Sprintf("0.0.0.0:%d", app.Config().Port),
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
