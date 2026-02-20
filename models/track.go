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
