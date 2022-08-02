package repositories

import (
	"musicSpringBootButInGo/models"

	"gorm.io/gorm"
)

type TracciaRepositoryInterface interface {
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

func (a *tracciaRepository) FindAll() ([]models.Traccia, error) {
	var tracce []models.Traccia
	result := a.DB.Find(&tracce)

	return tracce, result.Error
}

func (a *tracciaRepository) FindById(id uint) (*models.Traccia, error) {
	var traccia models.Traccia
	result := a.DB.Find(&traccia)

	var err error
	if result.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}

	return &traccia, err
}

func (a *tracciaRepository) Insert(traccia *models.Traccia) (uint, error) {
	result := a.DB.Create(&traccia)

	return traccia.TracciaId, result.Error
}

func (a *tracciaRepository) Update(traccia *models.Traccia) error {
	result := a.DB.Save(&traccia)

	return result.Error
}

func (a *tracciaRepository) Delete(traccia *models.Traccia) error {
	result := a.DB.Delete(&traccia)

	return result.Error
}
