package handlers

import (
	"net/http"

	"strconv"

	"github.com/evamdf/api-project/database"
	"github.com/evamdf/api-project/models"
	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	var albums []models.Album
	database.DB.Find(&albums)
	c.JSON(http.StatusOK, albums)
}

func GetAlbumByArtistID(c *gin.Context) {
	artist_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid artist id"})
		return
	}
	var albums []models.Album
	database.DB.Where("ArtistId = ?", artist_id).Find(&albums)
	c.JSON(http.StatusOK, albums)
}
