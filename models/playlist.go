package models

type Playlist struct {
	PlaylistId int     `json:"playlist_id" gorm:"column:PlaylistId;primaryKey"`
	Name       string  `json:"name" gorm:"column:Name"`
	Tracks     []Track `json:"tracks" gorm:"many2many:PlaylistTrack;foreignKey:PlaylistId;joinForeignKey:PlaylistId;References:TrackId;joinReferences:TrackId"`
}

// TableName specifies the exact table name in the Chinook DB so GORM
// uses the correct (case-sensitive) identifier.
func (Playlist) TableName() string {
	return "Playlist"
}
