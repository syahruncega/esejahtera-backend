package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type InstansiRepository interface {
	FindAll() ([]model.Instansi, error)
	FindById(id int) (model.Instansi, error)
	Create(instansi model.Instansi) (model.Instansi, error)
	Update(instansi model.Instansi) (model.Instansi, error)
	Delete(instansi model.Instansi) (model.Instansi, error)
}

type instansiRepository struct {
	db *gorm.DB
}

func NewInstansiRepository(db *gorm.DB) *instansiRepository {
	return &instansiRepository{db}
}

func (r *instansiRepository) FindAll() ([]model.Instansi, error) {
	var instansis []model.Instansi

	var err = r.db.Find(&instansis).Error

	return instansis, err
}

func (r *instansiRepository) FindById(id int) (model.Instansi, error) {
	var instansi model.Instansi

	var err = r.db.Take(&instansi, id).Error

	return instansi, err
}

func (r *instansiRepository) Create(instansi model.Instansi) (model.Instansi, error) {
	var err = r.db.Create(&instansi).Error

	return instansi, err
}

func (r *instansiRepository) Update(instansi model.Instansi) (model.Instansi, error) {
	var err = r.db.Model(&instansi).Updates(model.Instansi{
		InstansiId:   instansi.InstansiId,
		NamaInstansi: instansi.NamaInstansi,
	}).Error

	return instansi, err
}

func (r *instansiRepository) Delete(instansi model.Instansi) (model.Instansi, error) {
	var err = r.db.Delete(&instansi).Error

	return instansi, err
}
