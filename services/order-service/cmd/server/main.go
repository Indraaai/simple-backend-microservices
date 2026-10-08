package main

import (
	"log"

	"github.com/Indraaai/backend-microservices-lab/services/order-service/internal/config"
	"github.com/Indraaai/backend-microservices-lab/services/order-service/internal/database"
	"github.com/Indraaai/backend-microservices-lab/services/order-service/internal/handler"
	"github.com/Indraaai/backend-microservices-lab/services/order-service/internal/repository"
	"github.com/Indraaai/backend-microservices-lab/services/order-service/internal/router"
	"github.com/Indraaai/backend-microservices-lab/services/order-service/internal/service"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewPostgresPool(config.OrderDatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	orderRepository := repository.NewOrderRepository(db)

	orderService := service.NewOrderService(orderRepository)

	orderHandler := handler.NewOrderHandler(orderService)

	r := router.SetupRouter(orderHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
