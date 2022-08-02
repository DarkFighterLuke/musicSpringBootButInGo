package models

type Traccia struct {
	TracciaId     uint   `gorm:"primaryKey" json:"tracciaId"`
	NumeroTraccia int    `json:"numeroTraccia"`
	Titolo        string `json:"titolo"`
	Durata        int    `json:"durata"`
	AlbumId       uint   `json:"albumId"`
}

func (t *Traccia) TableName() string {
	return "tracce"
}
