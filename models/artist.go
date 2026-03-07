package models

type Artist struct {
	ArtistId int    `json:"artist_id" gorm:"column:ArtistId;primaryKey"`
	Name     string `json:"name" gorm:"column:Name"`
}

func (Artist) TableName() string {
	return "Artist"
}

func ToArtistResponse(a Artist) ArtistResponse {
	return ArtistResponse{
		ArtistId: a.ArtistId,
		Name:     a.Name,
	}
}

func ToArtistsResponse(artists []Artist) []ArtistResponse {
	var artistResponses []ArtistResponse
	for _, artist := range artists {
		artistResponses = append(artistResponses, ToArtistResponse(artist))
	}
	return artistResponses
}
