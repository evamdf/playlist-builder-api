package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/evamdf/api-project/database"
	"github.com/evamdf/api-project/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPlaylistTracks(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid playlist id"})
		return
	}

	var playlist models.Playlist
	result := database.DB.First(&playlist, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "playlist not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	var tracks []models.Track
	database.DB.Table("\"Track\"").
		Joins("JOIN \"PlaylistTrack\" ON \"PlaylistTrack\".\"TrackId\" = \"Track\".\"TrackId\"").
		Where("\"PlaylistTrack\".\"PlaylistId\" = ?", id).
		Find(&tracks)

	playlist.Tracks = tracks

	response := models.PlaylistTracksResponse{
		Name: playlist.Name,
	}
	for _, track := range playlist.Tracks {
		response.Tracks = append(response.Tracks, track.Name, track.Composer)
	}

	c.JSON(http.StatusOK, response)
}
