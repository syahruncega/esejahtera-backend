package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type KabupatenKotaRepository interface {
	FindAll() ([]model.Kabupaten_Kota, error)
	FindById(id string) (model.Kabupaten_Kota, error)
}

type kabupatenKotaRepository struct {
	db *gorm.DB
}

func NewKabupatenKotaRepository(db *gorm.DB) *kabupatenKotaRepository {
	return &kabupatenKotaRepository{db}
}

func (r *kabupatenKotaRepository) FindAll() ([]model.Kabupaten_Kota, error) {
	var kabupatenKotas []model.Kabupaten_Kota

	var err = r.db.Find(&kabupatenKotas).Error

	return kabupatenKotas, err
}

func (r *kabupatenKotaRepository) FindById(id string) (model.Kabupaten_Kota, error) {
	var kabupatenKota model.Kabupaten_Kota

	var err = r.db.Take(&kabupatenKota, id).Error

	return kabupatenKota, err
}
