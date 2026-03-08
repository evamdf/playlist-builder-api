package handlers

import (
	"net/http"
	"strconv"

	"github.com/evamdf/api-project/database"
	"github.com/evamdf/api-project/models"
	"github.com/gin-gonic/gin"
)

// Handler function to get all playlists (/api/v1/playlists)
func GetPlaylists(c *gin.Context) {
	var playlists []models.Playlist = []models.Playlist{}
	result := database.DB.Find(&playlists)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Error:   "failed to fetch playlists",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data:    models.ToPlaylistsResponse(playlists),
	})
}

// unexported little helper function to get a playlist by ID
func getPlaylistByID(id int) (*models.Playlist, error) {
	var playlist models.Playlist
	result := database.DB.First(&playlist, id)
	return &playlist, result.Error
}

// Handler function to get a playlist by ID (/api/v1/playlists/12)
func GetPlaylistByID(c *gin.Context) {
	// Parse the playlist ID from the URL parameter
	playlist_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Error:   "invalid playlist ID",
		})
		return
	}

	playlist, err := getPlaylistByID(playlist_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Error:   "invalid playlist ID",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data:    models.ToPlaylistResponse(*playlist),
	})
}
