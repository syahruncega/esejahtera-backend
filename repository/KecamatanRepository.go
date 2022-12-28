package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type KecamatanRepository interface {
	FindByKabupatenKota(kabupatenKotaId string) ([]model.Kecamatan, error)
	FindById(id string) (model.Kecamatan, error)
}

type kecamatanRepository struct {
	db *gorm.DB
}

func NewKecamatanRepository(db *gorm.DB) *kecamatanRepository {
	return &kecamatanRepository{db}
}

func (r *kecamatanRepository) FindByKabupatenKota(kabupatenKotaId string) ([]model.Kecamatan, error) {
	var kecamatans []model.Kecamatan

	var err = r.db.Where("kabupatenKotaId = ?", kabupatenKotaId).Find(&kecamatans).Error

	return kecamatans, err
}

func (r *kecamatanRepository) FindById(id string) (model.Kecamatan, error) {
	var kecamatan model.Kecamatan

	var err = r.db.Take(&kecamatan, id).Error

	return kecamatan, err
}
