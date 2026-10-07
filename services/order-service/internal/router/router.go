package router

import (
	"github.com/Indraaai/backend-microservices-lab/services/order-service/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	healthHandler := handler.NewHealthHandler()
	router.GET("/health", healthHandler.HealthCheck)
	return router
}
