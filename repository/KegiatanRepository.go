package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type KegiatanRepository interface {
	FindAll() ([]model.Kegiatan, error)
	FindById(id int) (model.Kegiatan, error)
	Create(kegiatan model.Kegiatan) (model.Kegiatan, error)
	Update(kegiatan model.Kegiatan) (model.Kegiatan, error)
	Delete(kegiatan model.Kegiatan) (model.Kegiatan, error)
}

type kegiatanRepository struct {
	db *gorm.DB
}

func NewKegiatanRepository(db *gorm.DB) *kegiatanRepository {
	return &kegiatanRepository{db}
}

func (r *kegiatanRepository) FindAll() ([]model.Kegiatan, error) {
	var kegiatans []model.Kegiatan

	var err = r.db.Find(&kegiatans).Error

	return kegiatans, err
}

func (r *kegiatanRepository) FindById(id int) (model.Kegiatan, error) {
	var kegiatan model.Kegiatan

	var err = r.db.Take(&kegiatan, id).Error

	return kegiatan, err
}

func (r *kegiatanRepository) Create(kegiatan model.Kegiatan) (model.Kegiatan, error) {
	var err = r.db.Create(&kegiatan).Error

	return kegiatan, err
}

func (r *kegiatanRepository) Update(kegiatan model.Kegiatan) (model.Kegiatan, error) {
	var err = r.db.Model(&kegiatan).Updates(model.Kegiatan{
		KegiatanId:   kegiatan.KegiatanId,
		NamaKegiatan: kegiatan.NamaKegiatan,
	}).Error

	return kegiatan, err
}

func (r *kegiatanRepository) Delete(kegiatan model.Kegiatan) (model.Kegiatan, error) {
	var err = r.db.Delete(&kegiatan).Error

	return kegiatan, err
}
