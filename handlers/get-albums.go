package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/evamdf/api-project/database"
	"github.com/evamdf/api-project/models"
	"github.com/gin-gonic/gin"
)

// Handler function to get all albums (/api/v1/albums)
func GetAlbums(c *gin.Context) {
	var albums []models.Album
	database.DB.Find(&albums)
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data:    albums,
	})
}

// Handler function to get all albums by a specific artist ID (/api/v1/artists/12/albums)
func GetAlbumByArtistID(c *gin.Context) {
	// Parse the artist ID from the URL parameter
	artist_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Error:   "invalid artist ID",
		})
		return
	}

	// Get the artist by ID to verify it exists
	artist, err := getArtistByID(artist_id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Error:   "artist not found",
		})
		return
	}

	// Get all the albums by that artist
	var albums []models.Album
	if result := database.DB.Where(`"ArtistId" = ?`, artist_id).Find(&albums); result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Error:   "failed to retrieve albums",
		})
		return
	}

	if len(albums) == 0 {
		c.JSON(http.StatusOK, models.Response{
			Success: true,
			Message: fmt.Sprintf("%s (%d) has no albums in this database", artist.Name, artist_id),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data:    albums,
	})

}
