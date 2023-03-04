package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type BidangUrusanOnInstansiRepository interface {
	FindAll() ([]model.BidangUrusanOnInstansi, error)
	FindById(id int) (model.BidangUrusanOnInstansi, error)
	Create(bidangUrusanOnInstansi model.BidangUrusanOnInstansi) (model.BidangUrusanOnInstansi, error)
	Update(bidangUrusanOnInstansi model.BidangUrusanOnInstansi) (model.BidangUrusanOnInstansi, error)
	Delete(bidangUrusanOnInstansi model.BidangUrusanOnInstansi) (model.BidangUrusanOnInstansi, error)
}

type bidangUrusanOnInstansiRepository struct {
	db *gorm.DB
}

func NewBidangUrusanOnInstansiRepository(db *gorm.DB) *bidangUrusanOnInstansiRepository {
	return &bidangUrusanOnInstansiRepository{db}
}

func (r *bidangUrusanOnInstansiRepository) FindAll() ([]model.BidangUrusanOnInstansi, error) {
	var bidangUrusanOnInstansi []model.BidangUrusanOnInstansi

	var err = r.db.Model(&bidangUrusanOnInstansi).Preload("Instansi").Preload("BidangUrusan").Find(&bidangUrusanOnInstansi).Error

	return bidangUrusanOnInstansi, err
}

func (r *bidangUrusanOnInstansiRepository) FindById(id int) (model.BidangUrusanOnInstansi, error) {
	var bidangUrusanOnInstansi model.BidangUrusanOnInstansi

	var err = r.db.Model(&bidangUrusanOnInstansi).Preload("Instansi").Preload("BidangUrusan").Take(&bidangUrusanOnInstansi, id).Error

	return bidangUrusanOnInstansi, err
}

func (r *bidangUrusanOnInstansiRepository) Create(bidangUrusanOnInstansi model.BidangUrusanOnInstansi) (model.BidangUrusanOnInstansi, error) {
	var err = r.db.Create(&bidangUrusanOnInstansi).Error

	return bidangUrusanOnInstansi, err
}

func (r *bidangUrusanOnInstansiRepository) Update(bidangUrusanOnInstansi model.BidangUrusanOnInstansi) (model.BidangUrusanOnInstansi, error) {
	var err = r.db.Model(&bidangUrusanOnInstansi).Updates(model.BidangUrusanOnInstansi{
		InstansiId:     bidangUrusanOnInstansi.InstansiId,
		BidangUrusanId: bidangUrusanOnInstansi.BidangUrusanId,
	}).Error

	return bidangUrusanOnInstansi, err
}

func (r *bidangUrusanOnInstansiRepository) Delete(bidangUrusanOnInstansi model.BidangUrusanOnInstansi) (model.BidangUrusanOnInstansi, error) {
	var err = r.db.Delete(&bidangUrusanOnInstansi).Error

	return bidangUrusanOnInstansi, err
}
