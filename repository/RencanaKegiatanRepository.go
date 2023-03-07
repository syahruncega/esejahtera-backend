package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type RencanaKegiatanRepository interface {
	FindAll() ([]model.RencanaKegiatan, error)
	FindById(id int) (model.RencanaKegiatan, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.RencanaKegiatan, error)
	Create(rencanaKegiatan model.RencanaKegiatan) (model.RencanaKegiatan, error)
	Update(rencanaKegiatan model.RencanaKegiatan) (model.RencanaKegiatan, error)
	Delete(rencanaKegiatan model.RencanaKegiatan) (model.RencanaKegiatan, error)
}

type rencanaKegiatanRepository struct {
	db *gorm.DB
}

func NewRencanaKegiatanRepository(db *gorm.DB) *rencanaKegiatanRepository {
	return &rencanaKegiatanRepository{db}
}

func (r *rencanaKegiatanRepository) FindAll() ([]model.RencanaKegiatan, error) {
	var rencanaKegiatans []model.RencanaKegiatan

	var err = r.db.Model(&rencanaKegiatans).Preload("RencanaProgram").Preload("Kegiatan").Find(&rencanaKegiatans).Error

	return rencanaKegiatans, err
}

func (r *rencanaKegiatanRepository) FindById(id int) (model.RencanaKegiatan, error) {
	var rencanaKegiatan model.RencanaKegiatan

	var err = r.db.Model(&rencanaKegiatan).Preload("RencanaProgram").Preload("Kegiatan").Find(&rencanaKegiatan, id).Error

	return rencanaKegiatan, err
}

func (r *rencanaKegiatanRepository) FindBySearch(whereClause map[string]interface{}) ([]model.RencanaKegiatan, error) {
	var rencanaKegiatans []model.RencanaKegiatan

	var err = r.db.Where(whereClause).Model(&rencanaKegiatans).Preload("RencanaProgram").Preload("Kegiatan").Find(&rencanaKegiatans).Error

	return rencanaKegiatans, err
}

func (r *rencanaKegiatanRepository) Create(rencanaKegiatan model.RencanaKegiatan) (model.RencanaKegiatan, error) {
	var err = r.db.Create(&rencanaKegiatan).Error

	return rencanaKegiatan, err
}

func (r *rencanaKegiatanRepository) Update(rencanaKegiatan model.RencanaKegiatan) (model.RencanaKegiatan, error) {
	var err = r.db.Model(&rencanaKegiatan).Updates(model.RencanaKegiatan{
		RencanaProgramId: rencanaKegiatan.RencanaProgramId,
		KegiatanId:       rencanaKegiatan.KegiatanId,
		PaguKegiatan:     rencanaKegiatan.PaguKegiatan,
		Tipe:             rencanaKegiatan.Tipe,
	}).Error

	return rencanaKegiatan, err
}

func (r *rencanaKegiatanRepository) Delete(rencanaKegiatan model.RencanaKegiatan) (model.RencanaKegiatan, error) {
	var err = r.db.Delete(&rencanaKegiatan).Error

	return rencanaKegiatan, err
}
