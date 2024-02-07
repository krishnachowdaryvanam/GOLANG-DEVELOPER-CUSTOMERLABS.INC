package handler

import (
	"golang_test_v_1/model"
	"golang_test_v_1/worker"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequestHandler(c *gin.Context) {
	var req model.OriginalEvent
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Send request to worker pool
	worker.WorkerPool <- req

	c.JSON(http.StatusOK, gin.H{"message": "Request received and processing started"})
}
