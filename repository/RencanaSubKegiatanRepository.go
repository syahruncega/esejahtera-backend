package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RencanaSubKegiatanRepository interface {
	FindAll() ([]model.RencanaSubKegiatan, error)
	FindById(id int) (model.RencanaSubKegiatan, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.RencanaSubKegiatan, error)
	Create(rencanaSubKegiatan model.RencanaSubKegiatan) (model.RencanaSubKegiatan, error)
	Update(rencanaSubKegiatan model.RencanaSubKegiatan) (model.RencanaSubKegiatan, error)
	Delete(rencanaSubKegiatan model.RencanaSubKegiatan) (model.RencanaSubKegiatan, error)
}

type rencanaSubKegiatanRepository struct {
	db *gorm.DB
}

func NewRencanaSubKegiatanRepository(db *gorm.DB) *rencanaSubKegiatanRepository {
	return &rencanaSubKegiatanRepository{db}
}

func (r *rencanaSubKegiatanRepository) FindAll() ([]model.RencanaSubKegiatan, error) {
	var rencanaSubKegiatans []model.RencanaSubKegiatan

	var err = r.db.Model(&rencanaSubKegiatans).Preload(clause.Associations).Preload("RencanaKegiatan." + clause.Associations).Preload("RencanaKegiatan.RencanaProgram." + clause.Associations).Preload("SubKegiatan").Find(&rencanaSubKegiatans).Error

	return rencanaSubKegiatans, err
}

func (r *rencanaSubKegiatanRepository) FindById(id int) (model.RencanaSubKegiatan, error) {
	var rencanaSubKegiatan model.RencanaSubKegiatan

	var err = r.db.Model(&rencanaSubKegiatan).Preload(clause.Associations).Preload("RencanaKegiatan."+clause.Associations).Preload("RencanaKegiatan.RencanaProgram."+clause.Associations).Preload("SubKegiatan").Take(&rencanaSubKegiatan, id).Error

	return rencanaSubKegiatan, err
}

func (r *rencanaSubKegiatanRepository) FindBySearch(whereClause map[string]interface{}) ([]model.RencanaSubKegiatan, error) {
	var rencanaSubKegiatans []model.RencanaSubKegiatan

	var err = r.db.Where(whereClause).Model(&rencanaSubKegiatans).Preload(clause.Associations).Preload("RencanaKegiatan." + clause.Associations).Preload("RencanaKegiatan.RencanaProgram." + clause.Associations).Preload("SubKegiatan").Find(&rencanaSubKegiatans).Error

	return rencanaSubKegiatans, err
}

func (r *rencanaSubKegiatanRepository) Create(rencanaSubKegiatan model.RencanaSubKegiatan) (model.RencanaSubKegiatan, error) {
	var err = r.db.Create(&rencanaSubKegiatan).Error

	return rencanaSubKegiatan, err
}

func (r *rencanaSubKegiatanRepository) Update(rencanaSubKegiatan model.RencanaSubKegiatan) (model.RencanaSubKegiatan, error) {
	var err = r.db.Model(&rencanaSubKegiatan).Updates(model.RencanaSubKegiatan{
		RencanaKegiatanId: rencanaSubKegiatan.RencanaKegiatanId,
		SubKegiatanId:     rencanaSubKegiatan.SubKegiatanId,
		PaguSubKegiatan:   rencanaSubKegiatan.PaguSubKegiatan,
		Tipe:              rencanaSubKegiatan.Tipe,
	}).Error

	return rencanaSubKegiatan, err
}

func (r *rencanaSubKegiatanRepository) Delete(rencanaSubKegiatan model.RencanaSubKegiatan) (model.RencanaSubKegiatan, error) {
	var err = r.db.Delete(&rencanaSubKegiatan).Error

	return rencanaSubKegiatan, err
}
