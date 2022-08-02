package repositories

import (
	"musicSpringBootButInGo/models"
	"sort"

	"gorm.io/gorm"
)

type TracciaRepositoryInterface interface {
	GetLastId() (uint, error)
	FindAll() ([]models.Traccia, error)
	FindById(id uint) (*models.Traccia, error)
	Insert(traccia *models.Traccia) (uint, error)
	Update(traccia *models.Traccia) error
	Delete(traccia *models.Traccia) error
}

type tracciaRepository struct {
	DB *gorm.DB
}

func NewTracciaRepository(DB *gorm.DB) TracciaRepositoryInterface {
	return &tracciaRepository{
		DB: DB,
	}
}

func (t *tracciaRepository) GetLastId() (uint, error) {
	tracce, err := t.FindAll()
	if err != nil {
		return 0, err
	}

	sort.Slice(tracce, func(i, j int) bool {
		return tracce[i].TracciaId > tracce[j].TracciaId
	})

	return tracce[0].TracciaId, nil
}

func (t *tracciaRepository) FindAll() ([]models.Traccia, error) {
	var tracce []models.Traccia
	result := t.DB.Find(&tracce)

	return tracce, result.Error
}

func (t *tracciaRepository) FindById(id uint) (*models.Traccia, error) {
	var traccia models.Traccia
	traccia.TracciaId = id
	result := t.DB.Debug().Find(&traccia)

	var err error
	if result.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}

	return &traccia, err
}

func (t *tracciaRepository) Insert(traccia *models.Traccia) (uint, error) {
	result := t.DB.Create(&traccia)

	return traccia.TracciaId, result.Error
}

func (t *tracciaRepository) Update(traccia *models.Traccia) error {
	result := t.DB.Save(&traccia)

	return result.Error
}

func (t *tracciaRepository) Delete(traccia *models.Traccia) error {
	result := t.DB.Delete(&traccia)

	return result.Error
}
