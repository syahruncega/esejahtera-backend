package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type IndividuRepository interface {
	FindAll(kabupatenKotaId string) ([]model.Individu, error)
	FindById(kabupatenKotaId string, id int) (model.Individu, error)
	FindByIdKeluarga(kabupatenKotaId string, idKeluarga string) ([]model.Individu, error)
	Update(individu model.Individu) (model.Individu, error)
}

type individuRepository struct {
	db *gorm.DB
}

func NewIndividuRepository(db *gorm.DB) *individuRepository {
	return &individuRepository{db}
}

func (r *individuRepository) FindAll(kabupatenKotaId string) ([]model.Individu, error) {
	var individus []model.Individu

	var err = r.db.Where("kabupatenKotaId = ?", kabupatenKotaId).Limit(15).Model(&individus).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Preload("User").Preload("Mahasiswa").Find(&individus).Error

	return individus, err
}

func (r *individuRepository) FindById(kabupatenKotaId string, id int) (model.Individu, error) {
	var individu model.Individu

	var err = r.db.Where("kabupatenKotaId = ?", kabupatenKotaId).Model(&individu).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Preload("User").Preload("Mahasiswa").Take(&individu, id).Error

	return individu, err
}

func (r *individuRepository) FindByIdKeluarga(kabupatenKotaId string, idKeluarga string) ([]model.Individu, error) {
	var individus []model.Individu

	var err = r.db.Where("kabupatenKotaId = ? and idKeluarga = ? ", kabupatenKotaId, idKeluarga).Model(&individus).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Preload("User").Preload("Mahasiswa").Find(&individus).Error

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
