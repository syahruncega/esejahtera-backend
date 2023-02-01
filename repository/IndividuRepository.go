package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type IndividuRepository interface {
	FindAll() ([]model.Individu, error)
	FindById(id int) (model.Individu, error)
	FindByIdKeluarga(idKeluarga string) ([]model.Individu, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.Individu, error)
	Update(individu model.Individu) (model.Individu, error)
}

type individuRepository struct {
	db *gorm.DB
}

func NewIndividuRepository(db *gorm.DB) *individuRepository {
	return &individuRepository{db}
}

func (r *individuRepository) FindAll() ([]model.Individu, error) {
	var individus []model.Individu

	var err = r.db.Limit(15).Model(&individus).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Preload("User").Preload("Mahasiswa").Find(&individus).Error

	return individus, err
}

func (r *individuRepository) FindById(id int) (model.Individu, error) {
	var individu model.Individu

	var err = r.db.Model(&individu).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Preload("User").Preload("Mahasiswa").Take(&individu, id).Error

	return individu, err
}

func (r *individuRepository) FindByIdKeluarga(idKeluarga string) ([]model.Individu, error) {
	var individus []model.Individu

	var err = r.db.Where("idKeluarga = ? ", idKeluarga).Model(&individus).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Preload("User").Preload("Mahasiswa").Find(&individus).Error

	return individus, err
}

func (r *individuRepository) FindBySearch(whereClause map[string]interface{}) ([]model.Individu, error) {
	var individus []model.Individu

	var err = r.db.Where(whereClause).Limit(15).Model(&individus).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Preload("User").Preload("Mahasiswa").Find(&individus).Error

	return individus, err
}

func (r *individuRepository) Update(individu model.Individu) (model.Individu, error) {
	var err = r.db.Model(&individu).Updates(model.Individu{
		UserId:           individu.UserId,
		MahasiswaId:      individu.MahasiswaId,
		StatusVerifikasi: individu.StatusVerifikasi,
	}).Error

	return individu, err
}
