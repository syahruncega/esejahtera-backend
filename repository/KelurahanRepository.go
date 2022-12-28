package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type KelurahanRepository interface {
	FindByKecamatanId(kecamatanId string) ([]model.Kelurahan, error)
	FindById(id string) (model.Kelurahan, error)
}

type kelurahanRepository struct {
	db *gorm.DB
}

func NewKelurahanRepository(db *gorm.DB) *kelurahanRepository {
	return &kelurahanRepository{db}
}

func (r *kelurahanRepository) FindByKecamatanId(kecamatanId string) ([]model.Kelurahan, error) {
	var kelurahans []model.Kelurahan

	var err = r.db.Where("kecamatanId = ? ", kecamatanId).Find(&kelurahans).Error

	return kelurahans, err
}

func (r *kelurahanRepository) FindById(id string) (model.Kelurahan, error) {
	var kelurahan model.Kelurahan

	var err = r.db.Take(&kelurahan, id).Error

	return kelurahan, err
}
