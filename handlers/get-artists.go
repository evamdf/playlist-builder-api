package handlers

import (
	"net/http"
	"strconv"

	"github.com/evamdf/api-project/database"
	"github.com/evamdf/api-project/models"
	"github.com/gin-gonic/gin"
)

// Handler function to get all artists (/api/v1/artists)
func GetArtists(c *gin.Context) {
	var artists []models.Artist
	database.DB.Find(&artists)
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data:    artists,
	})
}

// unexported little helper function to get an artist by ID
func getArtistByID(id int) (*models.Artist, error) {
	var artist models.Artist
	result := database.DB.First(&artist, id)
	return &artist, result.Error
}

// Handler function to get an artist by ID (/api/v1/artists/12)
func GetArtistByID(c *gin.Context) {
	// Parse the artist ID from the URL parameter
	artist_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Error:   "invalid artist ID",
		})
		return
	}

	artist, err := getArtistByID(artist_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Error:   "invalid artist ID",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data:    artist,
	})
}
