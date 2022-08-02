package repositories

import (
	"musicSpringBootButInGo/models"
	"sort"

	"gorm.io/gorm"
)

type AlbumRepositoryInterface interface {
	GetLastId() (uint, error)
	FindAll() ([]models.Album, error)
	FindById(id uint) (*models.Album, error)
	Insert(album *models.Album) (uint, error)
	Update(album *models.Album) error
	Delete(album *models.Album) error
}

type albumRepository struct {
	DB *gorm.DB
}

func NewAlbumRepository(DB *gorm.DB) AlbumRepositoryInterface {
	return &albumRepository{
		DB: DB,
	}
}

func (a *albumRepository) GetLastId() (uint, error) {
	albums, err := a.FindAll()
	if err != nil {
		return 0, err
	}

	sort.Slice(albums, func(i, j int) bool {
		return albums[i].AlbumId > albums[j].AlbumId
	})

	return albums[0].AlbumId, nil
}

func (a *albumRepository) FindAll() ([]models.Album, error) {
	var album []models.Album
	result := a.DB.Find(&album)

	return album, result.Error
}

func (a *albumRepository) FindById(id uint) (*models.Album, error) {
	var album models.Album
	album.AlbumId = id
	result := a.DB.Find(&album)

	var err error
	if result.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}

	return &album, err
}

func (a *albumRepository) Insert(album *models.Album) (uint, error) {
	result := a.DB.Create(&album)

	return album.AlbumId, result.Error
}

func (a *albumRepository) Update(album *models.Album) error {
	albumOld, err := a.FindById(album.AlbumId)
	if err != nil {
		return err
	}
	album.AlbumId = albumOld.AlbumId

	result := a.DB.Save(&album)

	return result.Error
}

func (a *albumRepository) Delete(album *models.Album) error {
	result := a.DB.Delete(&album)

	return result.Error
}
