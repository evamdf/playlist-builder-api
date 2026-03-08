package handlers

import (
	"net/http"
	"strconv"

	"github.com/evamdf/api-project/database"
	"github.com/evamdf/api-project/models"
	"github.com/gin-gonic/gin"
)

// unexported little helper function to get a track by ID
func getTrackByID(id int) (*models.Track, error) {
	var track models.Track
	result := database.DB.First(&track, id)
	return &track, result.Error
}

// Handler function to get a track by ID (/api/v1/tracks/12)
func GetFullTrackByID(c *gin.Context) {
	// Parse the track ID from the URL parameter
	artist_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Error:   "invalid track ID",
		})
		return
	}

	track, err := getTrackByID(artist_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Error:   "invalid track ID",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data:    models.ToFullTrackResponse(*track),
	})
}

// Handler function to get all the names and composers of the tracks in a playlist by playlist ID (/api/v1/playlists/:id)
func GetPlaylistTracks(c *gin.Context) {
	playlist_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Error:   "invalid playlist id",
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

	var tracks []models.Track
	database.DB.Table("\"Track\"").
		Joins("JOIN \"PlaylistTrack\" ON \"PlaylistTrack\".\"TrackId\" = \"Track\".\"TrackId\"").
		Where("\"PlaylistTrack\".\"PlaylistId\" = ?", playlist_id).
		Find(&tracks)

	tracks_response := models.ToTracksResponse(tracks)

	response := models.PlaylistTracksResponse{
		PlaylistId: playlist.PlaylistId,
		Name:       playlist.Name,
		Tracks:     tracks_response,
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data:    response,
	})
}
