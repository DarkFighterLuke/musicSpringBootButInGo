package repositories

import (
	"musicSpringBootButInGo/models"
	"sort"

	"gorm.io/gorm"
)

type ArtistaRepositoryInterface interface {
	GetLastId() (uint, error)
	FindAll() ([]models.Artista, error)
	FindById(id uint) (*models.Artista, error)
	Insert(artista *models.Artista) (uint, error)
	Update(artista *models.Artista) error
	Delete(artista *models.Artista) error
}

type artistaRepository struct {
	DB *gorm.DB
}

func NewArtistaRepository(DB *gorm.DB) ArtistaRepositoryInterface {
	return &artistaRepository{
		DB: DB,
	}
}

func (a *artistaRepository) GetLastId() (uint, error) {
	artisti, err := a.FindAll()
	if err != nil {
		return 0, err
	}

	sort.Slice(artisti, func(i, j int) bool {
		return artisti[i].ArtistaId > artisti[j].ArtistaId
	})

	return artisti[0].ArtistaId, nil
}

func (a *artistaRepository) FindAll() ([]models.Artista, error) {
	var artisti []models.Artista
	result := a.DB.Find(&artisti)

	return artisti, result.Error
}

func (a *artistaRepository) FindById(id uint) (*models.Artista, error) {
	var artista models.Artista
	artista.ArtistaId = id
	result := a.DB.Find(&artista)

	var err error
	if result.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}

	return &artista, err
}

func (a *artistaRepository) Insert(artista *models.Artista) (uint, error) {
	result := a.DB.Create(&artista)

	return artista.ArtistaId, result.Error
}

func (a *artistaRepository) Update(artista *models.Artista) error {
	result := a.DB.Save(&artista)

	return result.Error
}

func (a *artistaRepository) Delete(artista *models.Artista) error {
	result := a.DB.Delete(&artista)

	return result.Error
}
