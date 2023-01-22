package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type LokasiDosenRepository interface {
	FindAll() ([]model.Lokasi_Dosen, error)
	FindById(id int) (model.Lokasi_Dosen, error)
	FindAllRelation() ([]model.Lokasi_Dosen, error)
	FindRelationByDosenId(dosenId int) ([]model.Lokasi_Dosen, error)
	Create(lokasiDosen model.Lokasi_Dosen) (model.Lokasi_Dosen, error)
	Update(lokasiDosen model.Lokasi_Dosen) (model.Lokasi_Dosen, error)
	Delete(lokasiDosen model.Lokasi_Dosen) (model.Lokasi_Dosen, error)
}

type lokasiDosenRepository struct {
	db *gorm.DB
}

func NewLokasiDosenRepository(db *gorm.DB) *lokasiDosenRepository {
	return &lokasiDosenRepository{db}
}

func (r *lokasiDosenRepository) FindAll() ([]model.Lokasi_Dosen, error) {
	var lokasiDosens []model.Lokasi_Dosen

	var err = r.db.Find(&lokasiDosens).Error

	return lokasiDosens, err
}

func (r *lokasiDosenRepository) FindById(id int) (model.Lokasi_Dosen, error) {
	var lokasiDosen model.Lokasi_Dosen

	var err = r.db.Take(&lokasiDosen, id).Error

	return lokasiDosen, err
}

func (r *lokasiDosenRepository) FindAllRelation() ([]model.Lokasi_Dosen, error) {
	var lokasiDosens []model.Lokasi_Dosen

	var err = r.db.Model(&lokasiDosens).Preload("Dosen").Preload("Kelurahan").Find(&lokasiDosens).Error

	return lokasiDosens, err
}

func (r *lokasiDosenRepository) FindRelationByDosenId(dosenId int) ([]model.Lokasi_Dosen, error) {
	var lokasiDosens []model.Lokasi_Dosen

	var err = r.db.Where("dosenId = ? ", dosenId).Model(&lokasiDosens).Preload("Dosen").Preload("Kelurahan").Find(&lokasiDosens).Error

	return lokasiDosens, err
}

func (r *lokasiDosenRepository) Create(lokasiDosen model.Lokasi_Dosen) (model.Lokasi_Dosen, error) {
	var err = r.db.Create(&lokasiDosen).Error

	return lokasiDosen, err
}

func (r *lokasiDosenRepository) Update(lokasiDosen model.Lokasi_Dosen) (model.Lokasi_Dosen, error) {
	var err = r.db.Model(&lokasiDosen).Updates(model.Lokasi_Dosen{
		DosenId:     lokasiDosen.DosenId,
		KelurahanId: lokasiDosen.KelurahanId,
	}).Error

	return lokasiDosen, err
}

func (r *lokasiDosenRepository) Delete(lokasiDosen model.Lokasi_Dosen) (model.Lokasi_Dosen, error) {
	var err = r.db.Delete(&lokasiDosen).Error

	return lokasiDosen, err
}
