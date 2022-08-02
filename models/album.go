package models

import (
	"musicSpringBootButInGo/utils"
)

type Album struct {
	AlbumId    uint              `gorm:"primaryKey" json:"albumId"`
	Titolo     string            `json:"titolo"`
	DataUscita utils.CustomeTime `gorm:"embedded" json:"dataUscita"`
	Genere     string            `json:"genere"`
	ArtistaId  uint              `gorm:"column:artisti_id" json:"artistaId"`
	Tracce     []Traccia         `gorm:"foreignKey:TracceId"`
}

func (a *Album) TableName() string {
	return "album"
}
