package models

type Album struct {
	AlbumId  int    `json:"album_id" gorm:"column:AlbumId;primaryKey"`
	Title    string `json:"title" gorm:"column:Title"`
	ArtistId int    `json:"artist_id" gorm:"column:ArtistId"`
}

func (Album) TableName() string {
	return "Album"
}
