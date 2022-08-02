package models

type Artista struct {
	ArtistaId  uint    `gorm:"primaryKey;column:artisti_id" json:"artistaId"`
	Nome       string  `json:"nome"`
	Nazione    string  `json:"nazione"`
	AnnoInizio int     `json:"annoInizio"`
	Album      []Album `gorm:"foreignKey:ArtistaId"`
}

func (a *Artista) TableName() string {
	return "artisti"
}
