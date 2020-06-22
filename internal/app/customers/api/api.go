package api

import (
	_ "github.com/MarianoArias/ApiGo/cmd/customers/docs"
	"github.com/MarianoArias/ApiGo/internal/app/customers/controller"
	"github.com/MarianoArias/ApiGo/pkg/health-handler"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/zsais/go-gin-prometheus"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/customers/", controller.CgetHandler)

	// Api Doc Endpoint => /doc/index.html
	router.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Health Endpoint => /health/
	router.GET("/health/", healthhandler.HealthHandler)
	// Metrics Endpoint => /metrics/
	prometheus := ginprometheus.NewPrometheus("gin")
	prometheus.Use(router)

	return router
}
