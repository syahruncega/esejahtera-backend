package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type KeluargaRepository interface {
	FindAll(kabupatenKotaIds string) ([]model.Keluarga, error)
	FindById(kabupatenKotaId string, id int) (model.Keluarga, error)
	FindByIdKeluarga(kabupatenKotaId string, idKeluarga string) (model.Keluarga, error)
	FindPenerimaByKelurahan(kelurahanId string, penerimaParameter string, nilai string) (jumlah int64, err error)
}

type keluargaRepository struct {
	db *gorm.DB
}

func NewKeluargaRepository(db *gorm.DB) *keluargaRepository {
	return &keluargaRepository{db}
}

func (r *keluargaRepository) FindAll(kabupatenKotaId string) ([]model.Keluarga, error) {

	var keluargas []model.Keluarga

	var err = r.db.Where("kabupatenKotaId = ?", kabupatenKotaId).Limit(15).Model(&keluargas).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Find(&keluargas).Error

	return keluargas, err
}

func (r *keluargaRepository) FindById(kabupatenkotaId string, id int) (model.Keluarga, error) {
	var keluarga model.Keluarga

	var err = r.db.Where("kabupatenKotaId = ?", kabupatenkotaId).Model(&keluarga).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Take(&keluarga, id).Error

	return keluarga, err
}

func (r *keluargaRepository) FindByIdKeluarga(kabupatenKotaId string, idKeluarga string) (model.Keluarga, error) {
	var keluarga model.Keluarga

	var err = r.db.Where("kabupatenKotaId = ? AND idKeluarga = ?", kabupatenKotaId, idKeluarga).Model(&keluarga).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Take(&keluarga).Error

	return keluarga, err
}

func (r *keluargaRepository) FindPenerimaByKelurahan(kelurahanId string, penerimaParameter string, nilai string) (jumlah int64, err error) {
	var count int64

	var errr = r.db.Where(penerimaParameter+"=? and kelurahanId=?", nilai, kelurahanId).Table("keluargas").Select("count(" + penerimaParameter + ")").Count(&count).Error

	return count, errr

}
