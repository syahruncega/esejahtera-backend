package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type KeluargaSigiRepository interface {
	FindAll() ([]model.KeluargaSigi, error)
	FindById(id int) (model.KeluargaSigi, error)
	FindByIdKeluarga(idKeluarga string) (model.KeluargaSigi, error)
}

type keluargaSigiRepository struct {
	db *gorm.DB
}

func NewKeluargaSigiRepository(db *gorm.DB) *keluargaSigiRepository {
	return &keluargaSigiRepository{db}
}

func (r *keluargaSigiRepository) FindAll() ([]model.KeluargaSigi, error) {

	var keluargaSigis []model.KeluargaSigi

	var err = r.db.Limit(15).Model(&keluargaSigis).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Find(&keluargaSigis).Error

	return keluargaSigis, err
}

func (r *keluargaSigiRepository) FindById(id int) (model.KeluargaSigi, error) {
	var keluargaSigi model.KeluargaSigi

	var err = r.db.Model(&keluargaSigi).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Take(&keluargaSigi, id).Error

	return keluargaSigi, err
}

func (r *keluargaSigiRepository) FindByIdKeluarga(idKeluarga string) (model.KeluargaSigi, error) {
	var keluargaSigi model.KeluargaSigi

	var err = r.db.Where("idKeluarga = ?", idKeluarga).Model(&keluargaSigi).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Take(&keluargaSigi).Error

	return keluargaSigi, err
}
