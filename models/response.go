package models

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Standard response types ------------------
type ArtistResponse struct {
	ArtistId int    `json:"artist_id"`
	Name     string `json:"name"`
}

type AlbumResponse struct {
	AlbumId int    `json:"album_id"`
	Title   string `json:"title"`
}

type PlaylistResponse struct {
	PlaylistId int    `json:"playlist_id"`
	Name       string `json:"name"`
}

type FullTrackResponse struct {
	TrackId      int     `json:"track_id"`
	Name         string  `json:"name"`
	AlbumId      int     `json:"album_id"`
	MediaTypeId  int     `json:"media_type_id"`
	GenreId      int     `json:"genre_id"`
	Composer     string  `json:"composer"`
	Milliseconds int     `json:"milliseconds"`
	Bytes        int     `json:"bytes"`
	UnitPrice    float64 `json:"unit_price"`
}

type TrackResponse struct {
	TrackId  int    `json:"track_id"`
	Name     string `json:"name"`
	Composer string `json:"composer"`
}

// Specialized response types ------------------
type ArtistAlbumsResponse struct {
	ArtistId int             `json:"artist_id"`
	Name     string          `json:"name"`
	Albums   []AlbumResponse `json:"albums"`
}

type PlaylistTracksResponse struct {
	PlaylistId int             `json:"playlist_id"`
	Name       string          `json:"name"`
	Tracks     []TrackResponse `json:"tracks"`
}
