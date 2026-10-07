package main

import "github.com/Indraaai/backend-microservices-lab/services/order-service/internal/router"

func main() {
	server := router.SetupRouter()
	server.Run(":8080")
}
