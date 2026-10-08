package router

import (
	"github.com/gin-gonic/gin"

	"github.com/Indraaai/backend-microservices-lab/services/order-service/internal/handler"
)

func SetupRouter(
	orderHandler *handler.OrderHandler,
) *gin.Engine {

	router := gin.Default()

	healthHandler := handler.NewHealthHandler()

	router.GET("/health", healthHandler.HealthCheck)

	orders := router.Group("/orders")
	{
		orders.POST("", orderHandler.Create)
	}

	return router
}
