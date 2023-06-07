package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type FokusBelanjaRepository interface {
	FindAll() ([]model.FokusBelanja, error)
	FindById(id int) (model.FokusBelanja, error)
	FindByRencanaSubKegiatanId(rencanaSubKegiatanId int) ([]model.FokusBelanja, error)
	SumPaguFokusBelanja(rencanaSubKegiatanId int) (int64, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.FokusBelanja, error)
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

func (r *fokusBelanjaRepository) FindBySearch(whereClause map[string]interface{}) ([]model.FokusBelanja, error) {
	var fokusBelanjas []model.FokusBelanja

	var err = r.db.Where(whereClause).Model(&fokusBelanjas).Preload("RencanaSubKegiatan").Find(&fokusBelanjas).Error

	return fokusBelanjas, err
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
