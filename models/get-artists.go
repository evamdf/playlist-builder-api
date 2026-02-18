package models

type Artist struct {
	ArtistId int    `json:"artist_id" gorm:"column:ArtistId;primaryKey"`
	Name     string `json:"name" gorm:"column:Name"`
}

// TableName specifies the exact table name in the Chinook DB so GORM
// uses the correct (case-sensitive) identifier.
func (Artist) TableName() string {
	return "Artist"
}
