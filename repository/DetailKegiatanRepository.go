package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type DetailKegiatanRepository interface {
	FindAll() ([]model.DetailKegiatan, error)
	FindById(id int) (model.DetailKegiatan, error)
	FindByProgramId(programId string) ([]model.DetailKegiatan, error)
	Create(detailKegiatan model.DetailKegiatan) (model.DetailKegiatan, error)
	Update(detailKegiatan model.DetailKegiatan) (model.DetailKegiatan, error)
	Delete(detailKegiatan model.DetailKegiatan) (model.DetailKegiatan, error)
}

type detailKegiatanRepository struct {
	db *gorm.DB
}

func NewDetailKegiatanRepository(db *gorm.DB) *detailKegiatanRepository {
	return &detailKegiatanRepository{db}
}

func (r *detailKegiatanRepository) FindAll() ([]model.DetailKegiatan, error) {
	var detailKegiatans []model.DetailKegiatan

	var err = r.db.Model(&detailKegiatans).Preload("Program").Preload("Kegiatan").Find(&detailKegiatans).Error

	return detailKegiatans, err
}

func (r *detailKegiatanRepository) FindById(id int) (model.DetailKegiatan, error) {
	var detailKegiatan model.DetailKegiatan

	var err = r.db.Model(&detailKegiatan).Preload("Program").Preload("Kegiatan").Find(&detailKegiatan, id).Error

	return detailKegiatan, err
}

func (r *detailKegiatanRepository) FindByProgramId(programId string) ([]model.DetailKegiatan, error) {
	var detailKegiatans []model.DetailKegiatan

	var err = r.db.Model(&detailKegiatans).Preload("Program").Preload("Kegiatan").Find(&detailKegiatans).Error

	return detailKegiatans, err
}

func (r *detailKegiatanRepository) Create(detailKegiatan model.DetailKegiatan) (model.DetailKegiatan, error) {
	var err = r.db.Create(&detailKegiatan).Error

	return detailKegiatan, err
}

func (r *detailKegiatanRepository) Update(detailKegiatan model.DetailKegiatan) (model.DetailKegiatan, error) {
	var err = r.db.Model(&detailKegiatan).Updates(model.DetailKegiatan{
		ProgramId:  detailKegiatan.ProgramId,
		KegiatanId: detailKegiatan.KegiatanId,
	}).Error

	return detailKegiatan, err
}

func (r *detailKegiatanRepository) Delete(detailKegiatan model.DetailKegiatan) (model.DetailKegiatan, error) {
	var err = r.db.Delete(&detailKegiatan).Error

	return detailKegiatan, err
}
