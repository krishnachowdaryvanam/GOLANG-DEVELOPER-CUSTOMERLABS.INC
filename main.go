package main

import (
	"golang_test_v_1/handler"
	"golang_test_v_1/worker"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize worker pool
	worker.InitializeWorkerPool()

	router.POST("/event", handler.RequestHandler)

	router.Run(":8080")
}
