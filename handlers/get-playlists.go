package handlers

import (
	"net/http"

	"github.com/evamdf/api-project/database"
	"github.com/evamdf/api-project/models"
	"github.com/gin-gonic/gin"
)

func GetPlaylists(c *gin.Context) {
	var playlists []models.Playlist = []models.Playlist{}
	result := database.DB.Find(&playlists)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch playlists"})
		return
	}
	c.JSON(http.StatusOK, playlists)
}
