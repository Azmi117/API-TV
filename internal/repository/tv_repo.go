package repository

import (
	"github.com/Azmi117/API-TV.git/internal/models"
	"gorm.io/gorm"
)

type TvRepository struct {
	db *gorm.DB
}

func NewTvRepository(params *gorm.DB) *TvRepository {
	return &TvRepository{
		db: params,
	}
}

func (r *TvRepository) FindAll() ([]models.Tv, error) {
	var data []models.Tv
	err := r.db.Find(&data).Error
	return data, err
}

func (r *TvRepository) FindById(params int) (models.Tv, error) {
	var data models.Tv
	err := r.db.First(&data, params).Error
	return data, err
}

func (r *TvRepository) FindByName(params string) (models.Tv, error) {
	var data models.Tv
	err := r.db.Where("brand = ?", params).First(&data).Error
	return data, err
}

func (r *TvRepository) Create(params models.Tv) (models.Tv, error) {
	err := r.db.Create(&params).Error
	return params, err
}

func (r *TvRepository) Update(params models.Tv) error {
	return r.db.Model(&params).Updates(map[string]interface{}{
		"brand":      params.Brand,
		"price":      params.Price,
		"qty":        params.Qty,
		"updated_at": params.UpdatedAt,
	}).Error
}

func (r *TvRepository) Delete(params models.Tv) error {
	return r.db.Delete(&params).Error
}
