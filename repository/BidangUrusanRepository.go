package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type BidangUrusanRepository interface {
	FindAll() ([]model.BidangUrusan, error)
	FindById(id int) (model.BidangUrusan, error)
	Create(bidangUrusan model.BidangUrusan) (model.BidangUrusan, error)
	Update(bidangUrusan model.BidangUrusan) (model.BidangUrusan, error)
	Delete(bidangUrusan model.BidangUrusan) (model.BidangUrusan, error)
}

type bidangUrusanRepository struct {
	db *gorm.DB
}

func NewBidangUrusanRepository(db *gorm.DB) *bidangUrusanRepository {
	return &bidangUrusanRepository{db}
}

func (r *bidangUrusanRepository) FindAll() ([]model.BidangUrusan, error) {
	var bidangUrusans []model.BidangUrusan

	var err = r.db.Find(&bidangUrusans).Error

	return bidangUrusans, err
}

func (r *bidangUrusanRepository) FindById(id int) (model.BidangUrusan, error) {
	var bidangUrusan model.BidangUrusan

	var err = r.db.Take(&bidangUrusan, id).Error
	// var err = r.db.Model(model.BidangUrusan{Id: id}).First(&bidangUrusan).Error

	return bidangUrusan, err
}

func (r *bidangUrusanRepository) Create(bidangUrusan model.BidangUrusan) (model.BidangUrusan, error) {
	var err = r.db.Create(&bidangUrusan).Error

	return bidangUrusan, err
}

func (r *bidangUrusanRepository) Update(bidangUrusan model.BidangUrusan) (model.BidangUrusan, error) {
	var err = r.db.Model(&bidangUrusan).Updates(model.BidangUrusan{
		KodeBidangUrusan: bidangUrusan.KodeBidangUrusan,
		NamaBidangUrusan: bidangUrusan.NamaBidangUrusan,
	}).Error

	return bidangUrusan, err
}

func (r *bidangUrusanRepository) Delete(bidangUrusan model.BidangUrusan) (model.BidangUrusan, error) {
	var err = r.db.Delete(&bidangUrusan).Error

	return bidangUrusan, err
}
