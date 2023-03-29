package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RencanaKegiatanRepository interface {
	FindAll() ([]model.RencanaKegiatan, error)
	FindById(id int) (model.RencanaKegiatan, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.RencanaKegiatan, error)
	SumPaguRencanaKegiatan(rencanaProgramId int) (int64, error)
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

	var err = r.db.Model(&rencanaKegiatans).Preload(clause.Associations).Preload("RencanaProgram." + clause.Associations).Preload("Kegiatan").Find(&rencanaKegiatans).Error

	return rencanaKegiatans, err
}

func (r *rencanaKegiatanRepository) FindById(id int) (model.RencanaKegiatan, error) {
	var rencanaKegiatan model.RencanaKegiatan

	var err = r.db.Model(&rencanaKegiatan).Preload(clause.Associations).Preload("RencanaProgram."+clause.Associations).Preload("Kegiatan").Find(&rencanaKegiatan, id).Error

	return rencanaKegiatan, err
}

func (r *rencanaKegiatanRepository) FindBySearch(whereClause map[string]interface{}) ([]model.RencanaKegiatan, error) {
	var rencanaKegiatans []model.RencanaKegiatan

	var err = r.db.Where(whereClause).Model(&rencanaKegiatans).Preload(clause.Associations).Preload("RencanaProgram." + clause.Associations).Preload("Kegiatan").Find(&rencanaKegiatans).Error

	return rencanaKegiatans, err
}

func (r *rencanaKegiatanRepository) SumPaguRencanaKegiatan(rencanaProgramId int) (int64, error) {
	var totalPaguRencanaKegiatan int64

	var rows, err = r.db.Table("rencana_kegiatans").Where("rencanaProgramId = ?", rencanaProgramId).Select("sum(paguKegiatan) as totalPaguRencanaKegiatan").Rows()

	for rows.Next() {
		rows.Scan(&totalPaguRencanaKegiatan)
	}

	return totalPaguRencanaKegiatan, err
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
