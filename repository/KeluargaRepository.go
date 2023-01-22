package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type KeluargaRepository interface {
	FindAll(kabupatenKotaId string) ([]model.Keluarga, error)
	FindById(kabupatenKotaId string, id int) (model.Keluarga, error)
	FindByIdKeluargaByKabupatenKota(kabupatenKotaId string, idKeluarga string) (model.Keluarga, error)
	CountPenerimaByKabupatenKota(kabupatenKotaId string, penerimaParameter string, nilai string) (jumlah int64, err error)
	CountPenerimaByKecamatan(kecamatanId string, penerimaParameter string, nilai string) (jumlah int64, err error)
	CountPenerimaByKelurahan(kelurahanId string, penerimaParameter string, nilai string) (jumlah int64, err error)
	CountDesilByKabupatenKota(kabupatenKotaId string, nilaiDesil string) (jumlah int64, err error)
	CountDesilByKecamatan(kecamatanId string, nilaiDesil string) (jumlah int64, err error)
	CountDesilByKelurahan(kelurahanId string, nilaiDesil string) (jumlah int64, err error)
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

func (r *keluargaRepository) FindByIdKeluargaByKabupatenKota(kabupatenKotaId string, idKeluarga string) (model.Keluarga, error) {
	var keluarga model.Keluarga

	var err = r.db.Where("kabupatenKotaId = ? AND idKeluarga = ?", kabupatenKotaId, idKeluarga).Model(&keluarga).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Take(&keluarga).Error

	return keluarga, err
}

func (r *keluargaRepository) CountPenerimaByKabupatenKota(kabupatenKotaId string, penerimaParameter string, nilai string) (jumlah int64, err error) {
	var count int64

	var errr = r.db.Where(penerimaParameter+"=? and kabupatenKotaId=?", nilai, kabupatenKotaId).Table("keluargas").Select("count(" + penerimaParameter + ")").Count(&count).Error

	return count, errr
}

func (r *keluargaRepository) CountPenerimaByKecamatan(kecamatanId string, penerimaParameter string, nilai string) (jumlah int64, err error) {
	var count int64

	var errr = r.db.Where(penerimaParameter+"=? and kecamatanId=?", nilai, kecamatanId).Table("keluargas").Select("count(" + penerimaParameter + ")").Count(&count).Error

	return count, errr
}

func (r *keluargaRepository) CountPenerimaByKelurahan(kelurahanId string, penerimaParameter string, nilai string) (jumlah int64, err error) {
	var count int64

	var errr = r.db.Where(penerimaParameter+"=? and kelurahanId=?", nilai, kelurahanId).Table("keluargas").Select("count(" + penerimaParameter + ")").Count(&count).Error

	return count, errr
}

func (r *keluargaRepository) CountDesilByKabupatenKota(kabupatenKotaId string, nilaiDesil string) (jumlah int64, err error) {
	var count int64

	var errr = r.db.Where("kabupatenKotaId = ? and desilKesejahteraan = ?", kabupatenKotaId, nilaiDesil).Table("keluargas").Select("count(*)").Count(&count).Error

	return count, errr
}

func (r *keluargaRepository) CountDesilByKecamatan(kecamatanId string, nilaiDesil string) (jumlah int64, err error) {
	var count int64

	var errr = r.db.Where("kecamatanId = ? and desilKesejahteraan = ?", kecamatanId, nilaiDesil).Table("keluargas").Select("count(*)").Count(&count).Error

	return count, errr
}

func (r *keluargaRepository) CountDesilByKelurahan(kelurahanId string, nilaiDesil string) (jumlah int64, err error) {
	var count int64

	var errr = r.db.Where("kelurahanId = ? and desilKesejahteraan = ?", kelurahanId, nilaiDesil).Table("keluargas").Select("count(*)").Count(&count).Error

	return count, errr
}
