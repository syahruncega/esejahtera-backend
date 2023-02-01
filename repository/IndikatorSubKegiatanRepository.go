package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type IndikatorSubKegiatanRepository interface {
	FindAll() ([]model.IndikatorSubKegiatan, error)
	FindById(id int) (model.IndikatorSubKegiatan, error)
	Create(indikatorSubKegiatan model.IndikatorSubKegiatan) (model.IndikatorSubKegiatan, error)
	Update(indikatorSubKegiatan model.IndikatorSubKegiatan) (model.IndikatorSubKegiatan, error)
	Delete(indikatorSubKegiatan model.IndikatorSubKegiatan) (model.IndikatorSubKegiatan, error)
}

type indikatorSubKegiatanRepository struct {
	db *gorm.DB
}

func NewIndikatorSubKegiatanRepository(db *gorm.DB) *indikatorSubKegiatanRepository {
	return &indikatorSubKegiatanRepository{db}
}

func (r *indikatorSubKegiatanRepository) FindAll() ([]model.IndikatorSubKegiatan, error) {
	var indikatorSubkegiatans []model.IndikatorSubKegiatan

	var err = r.db.Model(&indikatorSubkegiatans).Preload("Kegiatan").Preload("SubKegiatan").Find(&indikatorSubkegiatans).Error

	return indikatorSubkegiatans, err
}

func (r *indikatorSubKegiatanRepository) FindById(id int) (model.IndikatorSubKegiatan, error) {
	var indikatorSubKegiatan model.IndikatorSubKegiatan

	var err = r.db.Model(&indikatorSubKegiatan).Preload("Kegiatan").Preload("SubKegiatan").Take(&indikatorSubKegiatan, id).Error

	return indikatorSubKegiatan, err
}

func (r *indikatorSubKegiatanRepository) Create(indikatorSubKegiatan model.IndikatorSubKegiatan) (model.IndikatorSubKegiatan, error) {
	var err = r.db.Create(&indikatorSubKegiatan).Error

	return indikatorSubKegiatan, err
}

func (r *indikatorSubKegiatanRepository) Update(indikatorSubKegiatan model.IndikatorSubKegiatan) (model.IndikatorSubKegiatan, error) {
	var err = r.db.Model(&indikatorSubKegiatan).Select("KegiatanId", "SubKegiatanId", "IndikatorKinerjaSubKegiatan", "PaguSubKegiatan").Updates(model.IndikatorSubKegiatan{
		KegiatanId:                  indikatorSubKegiatan.KegiatanId,
		SubKegiatanId:               indikatorSubKegiatan.SubKegiatanId,
		IndikatorKinerjaSubKegiatan: indikatorSubKegiatan.IndikatorKinerjaSubKegiatan,
		PaguSubKegiatan:             indikatorSubKegiatan.PaguSubKegiatan,
	}).Error

	return indikatorSubKegiatan, err
}

func (r *indikatorSubKegiatanRepository) Delete(indikatorSubKegiatan model.IndikatorSubKegiatan) (model.IndikatorSubKegiatan, error) {
	var err = r.db.Delete(&indikatorSubKegiatan).Error

	return indikatorSubKegiatan, err
}
