package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type MonevRepository interface {
	FindAll(kabupatenKotaId string) ([]model.Monev, error)
	FindById(kabupatenKotaId string, id int) (model.Monev, error)
}

type monevRepository struct {
	db *gorm.DB
}

func NewMonevRepository(db *gorm.DB) *monevRepository {
	return &monevRepository{db}
}

func (r *monevRepository) FindAll(kabupatenKotaId string) ([]model.Monev, error) {
	var monevs []model.Monev

	var err = r.db.Where("kabupatenKotaId = ?", kabupatenKotaId).Model(&monevs).Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Find(&monevs).Error

	return monevs, err
}

func (r *monevRepository) FindById(kabupatenKotaId string, id int) (model.Monev, error) {
	var monev model.Monev

	var err = r.db.Where("kabupatenKotaId = ?", kabupatenKotaId).Model(&monev).Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Find(&monev, id).Error

	return monev, err
}
