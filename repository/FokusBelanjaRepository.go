package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FokusBelanjaRepository interface {
	FindAll() ([]model.FokusBelanja, error)
	FindById(id int) (model.FokusBelanja, error)
	FindByRencanaSubKegiatanId(rencanaSubKegiatanId int) ([]model.FokusBelanja, error)
	SumPaguFokusBelanja(rencanaSubKegiatanId int) (int64, error)
	SumAllPaguFokusBelanja(tahun, tipe string) int64
	SumPaguFokusBelanjaByInstansi(tahun, tipe string, instansis []model.Instansi) []int64
	FindBySearch(whereClause map[string]interface{}) ([]model.FokusBelanja, error)
	CountJumlahFokusBelanja(tahun, tipe string) (int64, error)
	CountJumlahFokusBelanjaByInstansi(tahun, tipe string, instansis []model.Instansi) []int64
	Create(detailSubKegiatan model.FokusBelanja) (model.FokusBelanja, error)
	Update(detailSubKegiatan model.FokusBelanja) (model.FokusBelanja, error)
	Delete(detailSubKegiatan model.FokusBelanja) (model.FokusBelanja, error)
}

type fokusBelanjaRepository struct {
	db *gorm.DB
}

func NewFokusBelanjaRepository(db *gorm.DB) *fokusBelanjaRepository {
	return &fokusBelanjaRepository{db}
}

func (r *fokusBelanjaRepository) FindAll() ([]model.FokusBelanja, error) {
	var fokusBelanjas []model.FokusBelanja

	var err = r.db.Model(&fokusBelanjas).Preload("RencanaSubKegiatan").Find(&fokusBelanjas).Error

	return fokusBelanjas, err
}

func (r *fokusBelanjaRepository) FindById(id int) (model.FokusBelanja, error) {
	var fokusBelanja model.FokusBelanja

	var err = r.db.Model(&fokusBelanja).Preload("RencanaSubKegiatan").Take(&fokusBelanja, id).Error

	return fokusBelanja, err
}

func (r *fokusBelanjaRepository) FindByRencanaSubKegiatanId(rencanaSubKegiatanId int) ([]model.FokusBelanja, error) {
	var fokusBelanjas []model.FokusBelanja

	var err = r.db.Where("rencanaSubKegiatanId = ?", rencanaSubKegiatanId).Model(&fokusBelanjas).Preload("RencanaSubKegiatan").Find(&fokusBelanjas).Error

	return fokusBelanjas, err
}

func (r *fokusBelanjaRepository) SumPaguFokusBelanja(rencanaSubKegiatanId int) (int64, error) {
	var totalPaguFokusBelanja int64

	var rows, err = r.db.Table("fokus_belanjas").Where("rencanaSubKegiatanId = ?", rencanaSubKegiatanId).Select("sum(paguFokusBelanja) as totalPaguFokusBelanja").Rows()

	for rows.Next() {
		rows.Scan(&totalPaguFokusBelanja)
	}

	return totalPaguFokusBelanja, err
}

func (r *fokusBelanjaRepository) SumAllPaguFokusBelanja(tahun, tipe string) int64 {
	var jumlahPaguFokusBelanja int64

	var _ = r.db.Table("fokus_belanjas as fb").Select("sum(paguFokusBelanja) as jumlahPaguFokusBelanja").Joins("inner join rencana_sub_kegiatans as rsk on rsk.id = fb.rencanaSubKegiatanId").Where("fb.tahun = ? and rsk.tipe = ?", tahun, tipe).Scan(&jumlahPaguFokusBelanja)

	return jumlahPaguFokusBelanja
}

func (r *fokusBelanjaRepository) SumPaguFokusBelanjaByInstansi(tahun, tipe string, instansis []model.Instansi) []int64 {
	var sumPaguFokusBelanja int64
	var jumlahPaguFokusBelanja []int64

	for i := 0; i < len(instansis); i++ {

		var _ = r.db.Table("rencana_programs as rp").Select("coalesce(sum(paguFokusBelanja),0) as sumPaguFokusBelanja").Joins("inner join rencana_kegiatans as rk on rp.id = rk.rencanaProgramId").Joins("inner join rencana_sub_kegiatans as rsk on rsk.rencanaKegiatanId = rk.id").Joins("inner join fokus_belanjas as fb on fb.rencanaSubKegiatanId = rsk.id").Where("rp.tahun = ? and rp.tipe = ? and rp.instansiId = ?", tahun, tipe, instansis[i].Id).Scan(&sumPaguFokusBelanja)

		jumlahPaguFokusBelanja = append(jumlahPaguFokusBelanja, sumPaguFokusBelanja)
	}

	return jumlahPaguFokusBelanja
}

func (r *fokusBelanjaRepository) FindBySearch(whereClause map[string]interface{}) ([]model.FokusBelanja, error) {
	var fokusBelanjas []model.FokusBelanja

	var err = r.db.Where(whereClause).Model(&fokusBelanjas).Preload(clause.Associations).Preload("RencanaSubKegiatan." + clause.Associations).Preload("RencanaSubKegiatan.RencanaKegiatan." + clause.Associations).Preload("RencanaSubKegiatan.RencanaKegiatan.RencanaProgram." + clause.Associations).Find(&fokusBelanjas).Error

	return fokusBelanjas, err
}

func (r *fokusBelanjaRepository) CountJumlahFokusBelanja(tahun, tipe string) (int64, error) {
	var count int64

	var err = r.db.Select("count(*)").Table("fokus_belanjas as fk").Joins("inner join rencana_sub_kegiatans as rsk on rsk.id = fk.rencanaSubKegiatanId").Where("fk.tahun = ? and rsk.tipe = ?", tahun, tipe).Count(&count).Error

	return count, err
}

func (r *fokusBelanjaRepository) CountJumlahFokusBelanjaByInstansi(tahun, tipe string, instansis []model.Instansi) []int64 {
	var count int64
	var hasil []int64

	for i := 0; i < len(instansis); i++ {
		var _ = r.db.Select("count(*)").Table("rencana_programs as rp").Joins("inner join programs as p on p.id = rp.programId").Joins("inner join rencana_kegiatans as rk on rk.rencanaProgramId = rp.id").Joins("inner join rencana_sub_kegiatans as rsk on rsk.rencanaKegiatanId = rk.id").Joins("inner join fokus_belanjas as fb on fb.rencanaSubKegiatanId = rsk.id").Where("p.tahun = ? and rp.tipe = ? and rp.instansiId = ?", tahun, tipe, instansis[i].Id).Scan(&count)

		hasil = append(hasil, count)
	}

	return hasil
}

func (r *fokusBelanjaRepository) Create(fokusBelanja model.FokusBelanja) (model.FokusBelanja, error) {
	var err = r.db.Create(&fokusBelanja).Error

	return fokusBelanja, err
}

func (r *fokusBelanjaRepository) Update(fokusBelanja model.FokusBelanja) (model.FokusBelanja, error) {
	var err = r.db.Model(&fokusBelanja).Select("RencanaSubKegiatanId", "NamaFokusBelanja", "Indikator", "Target", "Satuan", "PaguFokusBelanja", "Keterangan", "Tahun").Updates(model.FokusBelanja{
		RencanaSubKegiatanId: fokusBelanja.RencanaSubKegiatanId,
		NamaFokusBelanja:     fokusBelanja.NamaFokusBelanja,
		Indikator:            fokusBelanja.Indikator,
		Target:               fokusBelanja.Target,
		Satuan:               fokusBelanja.Satuan,
		PaguFokusBelanja:     fokusBelanja.PaguFokusBelanja,
		Keterangan:           fokusBelanja.Keterangan,
		Tahun:                fokusBelanja.Tahun,
	}).Error

	return fokusBelanja, err
}

func (r *fokusBelanjaRepository) Delete(fokusBelanja model.FokusBelanja) (model.FokusBelanja, error) {
	var err = r.db.Delete(&fokusBelanja).Error

	return fokusBelanja, err
}
