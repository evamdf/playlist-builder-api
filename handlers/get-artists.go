package handlers

import (
	"net/http"

	"github.com/evamdf/api-project/database"
	"github.com/evamdf/api-project/models"
	"github.com/gin-gonic/gin"
)

func GetArtists(c *gin.Context) {
	var artist []models.Artist
	database.DB.Find(&artist)
	c.JSON(http.StatusOK, artist)
}
