package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type FokusBelanjaRepository interface {
	FindAll() ([]model.FokusBelanja, error)
	FindById(id int) (model.FokusBelanja, error)
	FindBySubKegiatanId(subkegiatanId int) ([]model.FokusBelanja, error)
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

func (r *fokusBelanjaRepository) FindBySubKegiatanId(subKegiatanId int) ([]model.FokusBelanja, error) {
	var fokusBelanjas []model.FokusBelanja

	var err = r.db.Where("subKegiatanId = ?", subKegiatanId).Model(&fokusBelanjas).Preload("RencanaSubKegiatan").Find(&fokusBelanjas).Error

	return fokusBelanjas, err
}

func (r *fokusBelanjaRepository) Create(fokusBelanja model.FokusBelanja) (model.FokusBelanja, error) {
	var err = r.db.Create(&fokusBelanja).Error

	return fokusBelanja, err
}

func (r *fokusBelanjaRepository) Update(fokusBelanja model.FokusBelanja) (model.FokusBelanja, error) {
	var err = r.db.Model(&fokusBelanja).Select("RencanaSubKegiatanId", "NamaFokusBelanja", "Indikator", "Target", "Satuan", "PaguFokusBelanja", "Keterangan").Updates(model.FokusBelanja{
		RencanaSubKegiatanId: fokusBelanja.RencanaSubKegiatanId,
		NamaFokusBelanja:     fokusBelanja.NamaFokusBelanja,
		Indikator:            fokusBelanja.Indikator,
		Target:               fokusBelanja.Target,
		Satuan:               fokusBelanja.Satuan,
		PaguFokusBelanja:     fokusBelanja.PaguFokusBelanja,
		Keterangan:           fokusBelanja.Keterangan,
	}).Error

	return fokusBelanja, err
}

func (r *fokusBelanjaRepository) Delete(fokusBelanja model.FokusBelanja) (model.FokusBelanja, error) {
	var err = r.db.Delete(&fokusBelanja).Error

	return fokusBelanja, err
}
