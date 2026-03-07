package models

type Track struct {
	TrackId      int     `json:"track_id" gorm:"column:TrackId;primaryKey"`
	Name         string  `json:"name" gorm:"column:Name"`
	AlbumId      int     `json:"album_id" gorm:"column:AlbumId"`
	MediaTypeId  int     `json:"media_type_id" gorm:"column:MediaTypeId"`
	GenreId      int     `json:"genre_id" gorm:"column:GenreId"`
	Composer     string  `json:"composer" gorm:"column:Composer"`
	Milliseconds int     `json:"milliseconds" gorm:"column:Milliseconds"`
	Bytes        int     `json:"bytes" gorm:"column:Bytes"`
	UnitPrice    float64 `json:"unit_price" gorm:"column:UnitPrice"`
}

func (Track) TableName() string {
	return "Track"
}

func ToFullTrackResponse(t Track) FullTrackResponse {
	return FullTrackResponse{
		TrackId:      t.TrackId,
		Name:         t.Name,
		AlbumId:      t.AlbumId,
		MediaTypeId:  t.MediaTypeId,
		GenreId:      t.GenreId,
		Composer:     t.Composer,
		Milliseconds: t.Milliseconds,
		Bytes:        t.Bytes,
		UnitPrice:    t.UnitPrice,
	}
}

func ToTrackResponse(t Track) TrackResponse {
	return TrackResponse{
		TrackId:  t.TrackId,
		Name:     t.Name,
		Composer: t.Composer,
	}
}

func ToTracksResponse(tracks []Track) []TrackResponse {
	var trackResponses []TrackResponse
	for _, track := range tracks {
		trackResponses = append(trackResponses, ToTrackResponse(track))
	}
	return trackResponses
}
