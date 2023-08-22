package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type RealisasiRepository interface {
	FindAll() ([]model.Realisasi, error)
	FindById(id int) (model.Realisasi, error)
	FindByFokusBelanjaId(fokusBelanjaId int) ([]model.Realisasi, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.Realisasi, error)
	Create(realisasi model.Realisasi) (model.Realisasi, error)
	Update(realisasi model.Realisasi) (model.Realisasi, error)
	Delete(realisasi model.Realisasi) (model.Realisasi, error)
}

type realisasiRepository struct {
	db *gorm.DB
}

func NewRealisasiRepository(db *gorm.DB) *realisasiRepository {
	return &realisasiRepository{db}
}

func (r *realisasiRepository) FindAll() ([]model.Realisasi, error) {
	var realisasis []model.Realisasi

	var err = r.db.Model(&realisasis).Preload("FokusBelanja").Find(&realisasis).Error

	return realisasis, err
}

func (r *realisasiRepository) FindById(id int) (model.Realisasi, error) {
	var realisasi model.Realisasi

	var err = r.db.Model(&realisasi).Preload("FokusBelanja").Take(&realisasi, id).Error

	return realisasi, err
}

func (r *realisasiRepository) FindByFokusBelanjaId(fokusBelanjaId int) ([]model.Realisasi, error) {
	var realisasis []model.Realisasi

	var err = r.db.Where("fokusBelanjaId = ? ", fokusBelanjaId).Model(&realisasis).Preload("FokusBelanja").Find(&realisasis).Error

	return realisasis, err
}

func (r *realisasiRepository) FindBySearch(whereClause map[string]interface{}) ([]model.Realisasi, error) {
	var realisasis []model.Realisasi

	var err = r.db.Where(whereClause).Model(&realisasis).Preload("FokusBelanja").Find(&realisasis).Error

	return realisasis, err
}

func (r *realisasiRepository) Create(realisasi model.Realisasi) (model.Realisasi, error) {
	var err = r.db.Create(&realisasi).Error

	return realisasi, err
}

func (r *realisasiRepository) Update(realisasi model.Realisasi) (model.Realisasi, error) {
	var err = r.db.Model(&realisasi).Updates(model.Realisasi{
		FokusBelanjaId:  realisasi.FokusBelanjaId,
		Tanggal:         realisasi.Tanggal,
		NomorSp2d:       realisasi.NomorSp2d,
		RealisasiPagu:   realisasi.RealisasiPagu,
		RealisasiTarget: realisasi.RealisasiTarget,
		Bulan:           realisasi.Bulan,
		Keterangan:      realisasi.Keterangan,
	}).Error

	return realisasi, err
}

func (r *realisasiRepository) Delete(realisasi model.Realisasi) (model.Realisasi, error) {
	var err = r.db.Delete(&realisasi).Error

	return realisasi, err
}
