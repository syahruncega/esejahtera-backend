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
	SumAllPaguRencanaKegiatan(tahun, tipe string) int64
	SumPaguRencanaKegiatanByInstansi(tahun, tipe string, instansis []model.Instansi) []int64
	CountJumlahRencanaKegiatan(tahun string, tipe string) (int64, error)
	CountJumlahRencanaKegiatanAllInstansi(tahun string, tipe string, instansis []model.Instansi) []int64
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

func (r *rencanaKegiatanRepository) SumAllPaguRencanaKegiatan(tahun, tipe string) int64 {
	var totalPaguRencanaKegiatan int64

	var rows, _ = r.db.Table("rencana_kegiatans").Where("tahun = ? and tipe = ?", tahun, tipe).Select("sum(paguKegiatan) as totalPaguRencanaKegiatan").Rows()

	for rows.Next() {
		rows.Scan(&totalPaguRencanaKegiatan)
	}

	return totalPaguRencanaKegiatan
}

func (r *rencanaKegiatanRepository) SumPaguRencanaKegiatanByInstansi(tahun, tipe string, instansis []model.Instansi) []int64 {
	var sumPaguRencanaKegiatan int64
	var totalPaguRencanaKegiatan []int64

	for i := 0; i < len(instansis); i++ {

		var rows, _ = r.db.Table("rencana_programs as rp").Select("sum(paguKegiatan) as sumPaguRencanaKegiatan").Joins("inner join rencana_kegiatans as rk on rp.id = rk.rencanaProgramId").Where("rp.tahun = ? and rp.tipe = ? and rp.instansiId = ?", tahun, tipe, instansis[i].Id).Rows()

		for rows.Next() {
			rows.Scan(&sumPaguRencanaKegiatan)
		}

		totalPaguRencanaKegiatan = append(totalPaguRencanaKegiatan, sumPaguRencanaKegiatan)
	}

	return totalPaguRencanaKegiatan
}

func (r *rencanaKegiatanRepository) CountJumlahRencanaKegiatan(tahun string, tipe string) (int64, error) {
	var count int64

	var err = r.db.Where("tahun = ? and tipe = ?", tahun, tipe).Table("rencana_kegiatans").Select("count(*)").Count(&count).Error

	return count, err
}

func (r *rencanaKegiatanRepository) CountJumlahRencanaKegiatanAllInstansi(tahun string, tipe string, instansis []model.Instansi) []int64 {
	var count int64
	var hasil []int64

	for i := 0; i < len(instansis); i++ {
		var _ = r.db.Select("count(*)").Table("rencana_programs as rp").Joins("inner join programs as p on p.id = rp.programId").Joins("inner join rencana_kegiatans as rk on rk.rencanaProgramId = rp.id").Where("p.tahun = ? and rp.tipe = ? and rp.instansiId = ?", tahun, tipe, instansis[i].Id).Scan(&count)

		hasil = append(hasil, count)
	}

	return hasil
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
		Tahun:            rencanaKegiatan.Tahun,
	}).Error

	return rencanaKegiatan, err
}

func (r *rencanaKegiatanRepository) Delete(rencanaKegiatan model.RencanaKegiatan) (model.RencanaKegiatan, error) {
	var err = r.db.Delete(&rencanaKegiatan).Error

	return rencanaKegiatan, err
}
