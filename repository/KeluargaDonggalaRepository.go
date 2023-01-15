package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type KeluargaDonggalaRepository interface {
	FindAll() ([]model.KeluargaDonggala, error)
	FindById(id int) (model.KeluargaDonggala, error)
	FindByIdKeluarga(idKeluarga string) (model.KeluargaDonggala, error)
}

type keluargaDonggalaRepository struct {
	db *gorm.DB
}

func NewKeluargaDonggalaRepository(db *gorm.DB) *keluargaDonggalaRepository {
	return &keluargaDonggalaRepository{db}
}

func (r *keluargaDonggalaRepository) FindAll() ([]model.KeluargaDonggala, error) {
	var keluargaDonggalas []model.KeluargaDonggala

	var err = r.db.Limit(15).Model(&keluargaDonggalas).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Find(&keluargaDonggalas).Error

	return keluargaDonggalas, err
}

func (r *keluargaDonggalaRepository) FindById(id int) (model.KeluargaDonggala, error) {
	var keluargaDonggala model.KeluargaDonggala

	var err = r.db.Model(&keluargaDonggala).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Take(&keluargaDonggala, id).Error

	return keluargaDonggala, err
}

func (r *keluargaDonggalaRepository) FindByIdKeluarga(idKeluarga string) (model.KeluargaDonggala, error) {
	var keluargaDonggala model.KeluargaDonggala

	var err = r.db.Where("idKeluarga = ?", idKeluarga).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Take(&keluargaDonggala).Error

	return keluargaDonggala, err
}
