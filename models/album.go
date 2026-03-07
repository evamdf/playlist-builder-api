package models

type Album struct {
	AlbumId  int    `json:"album_id" gorm:"column:AlbumId;primaryKey"`
	Title    string `json:"title" gorm:"column:Title"`
	ArtistId int    `json:"artist_id" gorm:"column:ArtistId"`
}

func (Album) TableName() string {
	return "Album"
}

func ToAlbumResponse(a Album) AlbumResponse {
	return AlbumResponse{
		AlbumId: a.AlbumId,
		Title:   a.Title,
	}
}

func ToAlbumsResponse(albums []Album) []AlbumResponse {
	var albumResponses []AlbumResponse
	for _, album := range albums {
		albumResponses = append(albumResponses, ToAlbumResponse(album))
	}
	return albumResponses
}
