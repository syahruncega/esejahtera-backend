package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type LokasiDosenRepository interface {
	FindAll() ([]model.LokasiDosen, error)
	FindById(id int) (model.LokasiDosen, error)
	FindAllRelation() ([]model.LokasiDosen, error)
	FindRelationByDosenId(dosenId int) ([]model.LokasiDosen, error)
	Create(lokasiDosen model.LokasiDosen) (model.LokasiDosen, error)
	Update(lokasiDosen model.LokasiDosen) (model.LokasiDosen, error)
	Delete(lokasiDosen model.LokasiDosen) (model.LokasiDosen, error)
}

type lokasiDosenRepository struct {
	db *gorm.DB
}

func NewLokasiDosenRepository(db *gorm.DB) *lokasiDosenRepository {
	return &lokasiDosenRepository{db}
}

func (r *lokasiDosenRepository) FindAll() ([]model.LokasiDosen, error) {
	var lokasiDosens []model.LokasiDosen

	var err = r.db.Find(&lokasiDosens).Error

	return lokasiDosens, err
}

func (r *lokasiDosenRepository) FindById(id int) (model.LokasiDosen, error) {
	var lokasiDosen model.LokasiDosen

	var err = r.db.Take(&lokasiDosen, id).Error

	return lokasiDosen, err
}

func (r *lokasiDosenRepository) FindAllRelation() ([]model.LokasiDosen, error) {
	var lokasiDosens []model.LokasiDosen

	var err = r.db.Model(&lokasiDosens).Preload("Dosen").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Find(&lokasiDosens).Error

	return lokasiDosens, err
}

func (r *lokasiDosenRepository) FindRelationByDosenId(dosenId int) ([]model.LokasiDosen, error) {
	var lokasiDosens []model.LokasiDosen

	var err = r.db.Where("dosenId = ? ", dosenId).Model(&lokasiDosens).Preload("Dosen").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Find(&lokasiDosens).Error

	return lokasiDosens, err
}

func (r *lokasiDosenRepository) Create(lokasiDosen model.LokasiDosen) (model.LokasiDosen, error) {
	var err = r.db.Create(&lokasiDosen).Error

	return lokasiDosen, err
}

func (r *lokasiDosenRepository) Update(lokasiDosen model.LokasiDosen) (model.LokasiDosen, error) {
	var err = r.db.Model(&lokasiDosen).Updates(model.LokasiDosen{
		DosenId:         lokasiDosen.DosenId,
		KabupatenKotaId: lokasiDosen.KabupatenKotaId,
		KecamatanId:     lokasiDosen.KecamatanId,
		KelurahanId:     lokasiDosen.KelurahanId,
	}).Error

	return lokasiDosen, err
}

func (r *lokasiDosenRepository) Delete(lokasiDosen model.LokasiDosen) (model.LokasiDosen, error) {
	var err = r.db.Delete(&lokasiDosen).Error

	return lokasiDosen, err
}
