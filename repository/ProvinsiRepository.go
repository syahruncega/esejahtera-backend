package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type ProvinsiRepository interface {
	FindAll() ([]model.Provinsi, error)
	FindById(id string) (model.Provinsi, error)
}

type provinsiRepository struct {
	db *gorm.DB
}

func NewProvinsiRepository(db *gorm.DB) *provinsiRepository {
	return &provinsiRepository{db}
}

func (r *provinsiRepository) FindAll() ([]model.Provinsi, error) {
	var provinsis []model.Provinsi

	var err = r.db.Find(&provinsis).Error

	return provinsis, err
}

func (r *provinsiRepository) FindById(id string) (model.Provinsi, error) {
	var provinsi model.Provinsi

	var err = r.db.Take(&provinsi, id).Error

	return provinsi, err
}
