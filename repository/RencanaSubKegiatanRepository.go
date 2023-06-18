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
	SumPaguRencanaSubKegiatan(rencanaKegiatanId int) (int64, error)
	CountJumlahRencanaSubKegiatan(tahun string, tipe string) (int64, error)
	CountJumlahRencanaSubKegiatanByInstansi(tahun string, tipe string, instansis []model.Instansi) []int64
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

func (r *rencanaSubKegiatanRepository) SumPaguRencanaSubKegiatan(rencanaKegiatanId int) (int64, error) {
	var totalPaguRencanaSubKegiatan int64

	var rows, err = r.db.Table("rencana_sub_kegiatans").Where("rencanaKegiatanId = ?", rencanaKegiatanId).Select("sum(paguSubKegiatan) as totalPaguRencanaSubKegiatan").Rows()

	for rows.Next() {
		rows.Scan(&totalPaguRencanaSubKegiatan)
	}

	return totalPaguRencanaSubKegiatan, err
}

func (r *rencanaSubKegiatanRepository) CountJumlahRencanaSubKegiatan(tahun string, tipe string) (int64, error) {
	var count int64

	var err = r.db.Where("tahun = ? and tipe = ?", tahun, tipe).Table("rencana_sub_kegiatans").Select("count(*)").Count(&count).Error

	return count, err
}

func (r *rencanaSubKegiatanRepository) CountJumlahRencanaSubKegiatanByInstansi(tahun string, tipe string, instansis []model.Instansi) []int64 {
	var count int64
	var hasil []int64

	for i := 0; i < len(instansis); i++ {
		var _ = r.db.Select("count(*)").Table("rencana_programs as rp").Joins("inner join programs as p on p.id = rp.programId").Joins("inner join rencana_kegiatans as rk on rk.rencanaProgramId = rp.id").Joins("inner join rencana_sub_kegiatans as rsk on rsk.rencanaKegiatanId = rk.id").Where("p.tahun = ? and rp.tipe = ? and rp.instansiId = ?", tahun, tipe, instansis[i].Id).Scan(&count)

		hasil = append(hasil, count)
	}

	return hasil
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
		Tahun:             rencanaSubKegiatan.Tahun,
	}).Error

	return rencanaSubKegiatan, err
}

func (r *rencanaSubKegiatanRepository) Delete(rencanaSubKegiatan model.RencanaSubKegiatan) (model.RencanaSubKegiatan, error) {
	var err = r.db.Delete(&rencanaSubKegiatan).Error

	return rencanaSubKegiatan, err
}
