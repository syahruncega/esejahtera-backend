package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type DetailInstansiRepository interface {
	FindAll() ([]model.DetailInstansi, error)
	FindById(id int) (model.DetailInstansi, error)
	FindByInstansiId(instansiId string) ([]model.DetailInstansi, error)
	Create(detailInstansi model.DetailInstansi) (model.DetailInstansi, error)
	Update(detailInstansi model.DetailInstansi) (model.DetailInstansi, error)
	Delete(detailInstansi model.DetailInstansi) (model.DetailInstansi, error)
}

type detailInstansiRepository struct {
	db *gorm.DB
}

func NewDetailInstansiRepository(db *gorm.DB) *detailInstansiRepository {
	return &detailInstansiRepository{db}
}

func (r *detailInstansiRepository) FindAll() ([]model.DetailInstansi, error) {
	var detailInstansis []model.DetailInstansi

	var err = r.db.Model(&detailInstansis).Preload("Instansi").Preload("BidangUrusan").Find(&detailInstansis).Error

	return detailInstansis, err
}

func (r *detailInstansiRepository) FindById(id int) (model.DetailInstansi, error) {
	var detailInstansi model.DetailInstansi

	var err = r.db.Model(&detailInstansi).Preload("Instansi").Preload("BidangUrusan").Take(&detailInstansi, id).Error

	return detailInstansi, err
}

func (r *detailInstansiRepository) FindByInstansiId(instansiId string) ([]model.DetailInstansi, error) {
	var detailInstansis []model.DetailInstansi

	var err = r.db.Where("instansiId = ? ", instansiId).Model(&detailInstansis).Preload("Instansi").Preload("BidangUrusan").Find(&detailInstansis).Error

	return detailInstansis, err
}

func (r *detailInstansiRepository) Create(detailInstansi model.DetailInstansi) (model.DetailInstansi, error) {
	var err = r.db.Create(&detailInstansi).Error

	return detailInstansi, err
}

func (r *detailInstansiRepository) Update(detailInstansi model.DetailInstansi) (model.DetailInstansi, error) {
	var err = r.db.Model(&detailInstansi).Updates(model.DetailInstansi{
		InstansiId:     detailInstansi.InstansiId,
		BidangUrusanId: detailInstansi.BidangUrusanId,
	}).Error

	return detailInstansi, err
}

func (r *detailInstansiRepository) Delete(detailInstansi model.DetailInstansi) (model.DetailInstansi, error) {
	var err = r.db.Delete(&detailInstansi).Error

	return detailInstansi, err
}
