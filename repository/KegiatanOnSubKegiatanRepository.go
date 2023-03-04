package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type KegiatanOnSubKegiatanRepository interface {
	FindAll() ([]model.KegiatanOnSubKegiatan, error)
	FindById(id int) (model.KegiatanOnSubKegiatan, error)
	Create(kegiatanOnSubKegiatan model.KegiatanOnSubKegiatan) (model.KegiatanOnSubKegiatan, error)
	Update(kegiatanOnSubKegiatan model.KegiatanOnSubKegiatan) (model.KegiatanOnSubKegiatan, error)
	Delete(kegiatanOnSubKegiatan model.KegiatanOnSubKegiatan) (model.KegiatanOnSubKegiatan, error)
}

type kegiatanOnSubKegiatanRepository struct {
	db *gorm.DB
}

func NewKegiatanOnSubKegiatanRepository(db *gorm.DB) *kegiatanOnSubKegiatanRepository {
	return &kegiatanOnSubKegiatanRepository{db}
}

func (r *kegiatanOnSubKegiatanRepository) FindAll() ([]model.KegiatanOnSubKegiatan, error) {
	var kegiatanOnSubKegiatan []model.KegiatanOnSubKegiatan

	var err = r.db.Model(&kegiatanOnSubKegiatan).Preload("Kegiatan").Preload("SubKegiatan").Find(&kegiatanOnSubKegiatan).Error

	return kegiatanOnSubKegiatan, err
}

func (r *kegiatanOnSubKegiatanRepository) FindById(id int) (model.KegiatanOnSubKegiatan, error) {
	var kegiatanOnSubKegiatan model.KegiatanOnSubKegiatan

	var err = r.db.Model(&kegiatanOnSubKegiatan).Preload("Kegiatan").Preload("SubKegiatan").Take(&kegiatanOnSubKegiatan, id).Error

	return kegiatanOnSubKegiatan, err
}

func (r *kegiatanOnSubKegiatanRepository) Create(kegiatanOnSubKegiatan model.KegiatanOnSubKegiatan) (model.KegiatanOnSubKegiatan, error) {
	var err = r.db.Create(&kegiatanOnSubKegiatan).Error

	return kegiatanOnSubKegiatan, err
}

func (r *kegiatanOnSubKegiatanRepository) Update(kegiatanOnSubKegiatan model.KegiatanOnSubKegiatan) (model.KegiatanOnSubKegiatan, error) {
	var err = r.db.Model(&kegiatanOnSubKegiatan).Updates(model.KegiatanOnSubKegiatan{
		KegiatanId:    kegiatanOnSubKegiatan.KegiatanId,
		SubKegiatanId: kegiatanOnSubKegiatan.SubKegiatanId,
	}).Error

	return kegiatanOnSubKegiatan, err
}

func (r *kegiatanOnSubKegiatanRepository) Delete(kegiatanOnSubKegiatan model.KegiatanOnSubKegiatan) (model.KegiatanOnSubKegiatan, error) {
	var err = r.db.Delete(&kegiatanOnSubKegiatan).Error

	return kegiatanOnSubKegiatan, err
}
