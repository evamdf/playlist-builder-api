package handlers

import (
	"net/http"
	"time"

	"github.com/evamdf/api-project/database"
	"github.com/evamdf/api-project/models"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	dbStatus := "connected"
	if database.DB == nil {
		dbStatus = "not connected"
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "API is running",
		Data: map[string]interface{}{
			"database": dbStatus,
			"time":     time.Now().Format(time.RFC3339),
		},
	})
}
