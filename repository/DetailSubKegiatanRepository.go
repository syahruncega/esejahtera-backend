package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type DetailSubKegiatanRepository interface {
	FindAll() ([]model.DetailSubKegiatan, error)
	FindById(id int) (model.DetailSubKegiatan, error)
	FindByKegiatanId(kegiatanId string) ([]model.DetailSubKegiatan, error)
	Create(detailSubKegiatan model.DetailSubKegiatan) (model.DetailSubKegiatan, error)
	Update(detailSubKegiatan model.DetailSubKegiatan) (model.DetailSubKegiatan, error)
	Delete(detailSubKegiatan model.DetailSubKegiatan) (model.DetailSubKegiatan, error)
}

type detailSubKegiatanRepository struct {
	db *gorm.DB
}

func NewDetailSubKegiatanRepository(db *gorm.DB) *detailSubKegiatanRepository {
	return &detailSubKegiatanRepository{db}
}

func (r *detailSubKegiatanRepository) FindAll() ([]model.DetailSubKegiatan, error) {
	var detailSubKegiatans []model.DetailSubKegiatan

	var err = r.db.Model(&detailSubKegiatans).Preload("Kegiatan").Preload("SubKegiatan").Find(&detailSubKegiatans).Error

	return detailSubKegiatans, err
}

func (r *detailSubKegiatanRepository) FindById(id int) (model.DetailSubKegiatan, error) {
	var detailSubKegiatan model.DetailSubKegiatan

	var err = r.db.Model(&detailSubKegiatan).Preload("Kegiatan").Preload("SubKegiatan").Take(&detailSubKegiatan, id).Error

	return detailSubKegiatan, err
}

func (r *detailSubKegiatanRepository) FindByKegiatanId(kegiatanId string) ([]model.DetailSubKegiatan, error) {
	var detailSubKegiatans []model.DetailSubKegiatan

	var err = r.db.Where("kegiatanId = ? ", kegiatanId).Model(&detailSubKegiatans).Preload("Kegiatan").Preload("SubKegiatan").Find(&detailSubKegiatans).Error

	return detailSubKegiatans, err
}

func (r *detailSubKegiatanRepository) Create(detailSubKegiatan model.DetailSubKegiatan) (model.DetailSubKegiatan, error) {
	var err = r.db.Create(&detailSubKegiatan).Error

	return detailSubKegiatan, err
}

func (r *detailSubKegiatanRepository) Update(detailSubKegiatan model.DetailSubKegiatan) (model.DetailSubKegiatan, error) {
	var err = r.db.Model(&detailSubKegiatan).Updates(model.DetailSubKegiatan{
		KegiatanId:    detailSubKegiatan.KegiatanId,
		SubKegiatanId: detailSubKegiatan.SubKegiatanId,
	}).Error

	return detailSubKegiatan, err
}

func (r *detailSubKegiatanRepository) Delete(detailSubKegiatan model.DetailSubKegiatan) (model.DetailSubKegiatan, error) {
	var err = r.db.Delete(&detailSubKegiatan).Error

	return detailSubKegiatan, err
}
