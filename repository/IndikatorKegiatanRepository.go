package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type IndikatorKegiatanRepository interface {
	FindAll() ([]model.IndikatorKegiatan, error)
	FindById(id int) (model.IndikatorKegiatan, error)
	Create(indikatorKegiatan model.IndikatorKegiatan) (model.IndikatorKegiatan, error)
	Update(indikatorKegiatan model.IndikatorKegiatan) (model.IndikatorKegiatan, error)
	Delete(indikatorKegiatan model.IndikatorKegiatan) (model.IndikatorKegiatan, error)
}

type indikatorKegiatanRepository struct {
	db *gorm.DB
}

func NewIndikatorKegiatanRepository(db *gorm.DB) *indikatorKegiatanRepository {
	return &indikatorKegiatanRepository{db}
}

func (r *indikatorKegiatanRepository) FindAll() ([]model.IndikatorKegiatan, error) {
	var indikatorKegiatans []model.IndikatorKegiatan

	var err = r.db.Model(&indikatorKegiatans).Preload("Program").Preload("Kegiatan").Find(&indikatorKegiatans).Error

	return indikatorKegiatans, err
}

func (r *indikatorKegiatanRepository) FindById(id int) (model.IndikatorKegiatan, error) {
	var indikatorKegiatan model.IndikatorKegiatan

	var err = r.db.Model(&indikatorKegiatan).Preload("Program").Preload("Kegiatan").Take(&indikatorKegiatan, id).Error

	return indikatorKegiatan, err
}

func (r *indikatorKegiatanRepository) Create(indikatorKegiatan model.IndikatorKegiatan) (model.IndikatorKegiatan, error) {
	var err = r.db.Create(&indikatorKegiatan).Error

	return indikatorKegiatan, err
}

func (r *indikatorKegiatanRepository) Update(indikatorKegiatan model.IndikatorKegiatan) (model.IndikatorKegiatan, error) {
	var err = r.db.Model(&indikatorKegiatan).Select("ProgramId", "KegiatanId", "IndikatorKinerjaKegiatan", "PaguKegiatan").Updates(model.IndikatorKegiatan{
		ProgramId:                indikatorKegiatan.ProgramId,
		KegiatanId:               indikatorKegiatan.KegiatanId,
		IndikatorKinerjaKegiatan: indikatorKegiatan.IndikatorKinerjaKegiatan,
		PaguKegiatan:             indikatorKegiatan.PaguKegiatan,
	}).Error

	return indikatorKegiatan, err
}

func (r *indikatorKegiatanRepository) Delete(indikatorKegiatan model.IndikatorKegiatan) (model.IndikatorKegiatan, error) {
	var err = r.db.Delete(&indikatorKegiatan).Error

	return indikatorKegiatan, err
}
