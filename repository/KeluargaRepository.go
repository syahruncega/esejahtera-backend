package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type KeluargaRepository interface {
	FindAll() ([]model.Keluarga, error)
	FindById(id int) (model.Keluarga, error)
	FindByIdKeluarga(idKeluarga string) (model.Keluarga, error)
	CountData(whereClause map[string]interface{}) (int64, error)
	FindBySearch(whereClause map[string]interface{}, limit int, offset int) ([]model.Keluarga, error)
	Update(keluarga model.Keluarga) (model.Keluarga, error)
	CountJumlahKeluarga(places string, placesId string) (int64, error)
	CountJumlahDesil(places string, placesId string, desil string) (int64, error)
	CountJumlahPenerima(places string, placesId string, penerima string, penerimaValue string) (int64, error)
	DistinctKabupatenKota() ([]model.DistinctKabupatenKota, error)
	DistinctCountKabupatenKota() (map[string]int64, error)
	DistinctKecamatan() ([]model.DistinctKecamatan, error)
	DistinctKecamatanByKabupatenKota(kabupatenKotaId string) ([]model.DistinctKecamatan, error)
	DistinctCountKecamatan(kabupatenKotaId string) (map[string]int64, error)
	DistinctKelurahanByKecamatan(kecamatanId string) ([]model.DistinctKelurahan, error)
	DistinctCountKelurahan(kecamatanId string) (map[string]int64, error)
}

type keluargaRepository struct {
	db *gorm.DB
}

func NewKeluargaRepository(db *gorm.DB) *keluargaRepository {
	return &keluargaRepository{db}
}

func (r *keluargaRepository) FindAll() ([]model.Keluarga, error) {

	var keluargas []model.Keluarga

	var err = r.db.Limit(15).Model(&keluargas).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Preload("User").Preload("Mahasiswa").Find(&keluargas).Error

	return keluargas, err
}

func (r *keluargaRepository) FindById(id int) (model.Keluarga, error) {
	var keluarga model.Keluarga

	var err = r.db.Model(&keluarga).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Preload("User").Preload("Mahasiswa").Take(&keluarga, id).Error

	return keluarga, err
}

func (r *keluargaRepository) FindByIdKeluarga(idKeluarga string) (model.Keluarga, error) {
	var keluarga model.Keluarga

	var err = r.db.Where("idKeluarga = ?", idKeluarga).Model(&keluarga).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Preload("User").Preload("Mahasiswa").Take(&keluarga).Error

	return keluarga, err
}

func (r *keluargaRepository) CountData(whereClause map[string]interface{}) (int64, error) {
	var count int64

	var err = r.db.Where(whereClause).Table("keluargas").Select("count(*)").Count(&count).Error

	return count, err
}

func (r *keluargaRepository) FindBySearch(whereClause map[string]interface{}, limit int, offset int) ([]model.Keluarga, error) {
	var keluargas []model.Keluarga

	var err = r.db.Where(whereClause).Limit(limit).Offset(offset).Model(&keluargas).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Preload("User").Preload("Mahasiswa").Find(&keluargas).Error

	return keluargas, err
}

func (r *keluargaRepository) Update(keluarga model.Keluarga) (model.Keluarga, error) {
	var err = r.db.Model(&keluarga).Updates(model.Keluarga{
		UserId:           keluarga.UserId,
		StatusVerifikasi: keluarga.StatusVerifikasi,
		MahasiswaId:      keluarga.MahasiswaId,
	}).Error

	return keluarga, err
}

func (r *keluargaRepository) CountJumlahKeluarga(places string, placesId string) (int64, error) {
	var count int64

	var err = r.db.Where(places+"= ?", placesId).Table("keluargas").Select("count(*)").Count(&count).Error

	return count, err
}

func (r *keluargaRepository) CountJumlahDesil(places string, placesId string, desil string) (int64, error) {
	var count int64

	var err = r.db.Where(places+"= ? and desilKesejahteraan = ?", placesId, desil).Table("keluargas").Select("count(*)").Count(&count).Error

	return count, err
}

func (r *keluargaRepository) CountJumlahPenerima(places string, placesId string, penerima string, penerimaValue string) (int64, error) {
	var count int64

	var err = r.db.Where(places+"= ? and "+penerima+" = ?", placesId, penerimaValue).Table("keluargas").Select("count(*)").Count(&count).Error

	return count, err
}

func (r *keluargaRepository) DistinctKabupatenKota() ([]model.DistinctKabupatenKota, error) {
	var distinctKabupatenKota []model.DistinctKabupatenKota

	var err = r.db.Distinct("k.kabupatenKotaId, kb.nama").Table("keluargas as k").Joins("inner join kabupaten_kota as kb on kb.id = k.kabupatenKotaId").Scan(&distinctKabupatenKota).Error

	return distinctKabupatenKota, err
}

func (r *keluargaRepository) DistinctCountKabupatenKota() (map[string]int64, error) {
	var distinctsKabupatenKota, err = r.DistinctKabupatenKota()

	var jumlah = make(map[string]int64)

	for _, distinctKabupatenKota := range distinctsKabupatenKota {
		var test, _ = r.CountJumlahKeluarga("kabupatenKotaId", distinctKabupatenKota.KabupatenKotaId)

		jumlah[distinctKabupatenKota.Nama] = test
	}

	return jumlah, err
}

func (r *keluargaRepository) DistinctKecamatan() ([]model.DistinctKecamatan, error) {
	var distinctsKecamatan []model.DistinctKecamatan

	var err = r.db.Distinct("k.kecamatanId, kc.nama").Table("keluargas as k").Joins("inner join kecamatans as kc on kc.id = k.kecamatanId").Scan(&distinctsKecamatan).Error

	return distinctsKecamatan, err
}

func (r *keluargaRepository) DistinctKecamatanByKabupatenKota(kabupatenKotaId string) ([]model.DistinctKecamatan, error) {
	var distinctsKecamatan []model.DistinctKecamatan

	var err = r.db.Distinct("k.kecamatanId, kc.nama").Table("keluargas as k").Joins("inner join kecamatans as kc on kc.id = k.kecamatanId").Where("k.kabupatenKotaId = ?", kabupatenKotaId).Scan(&distinctsKecamatan).Error

	return distinctsKecamatan, err
}

func (r *keluargaRepository) DistinctCountKecamatan(kabupatenKotaId string) (map[string]int64, error) {
	var distinctsKecamatan, err = r.DistinctKecamatanByKabupatenKota(kabupatenKotaId)

	var jumlah = make(map[string]int64)

	for _, distinctKecamatan := range distinctsKecamatan {
		var test, _ = r.CountJumlahKeluarga("kecamatanId", distinctKecamatan.KecamatanId)

		jumlah[distinctKecamatan.Nama] = test
	}

	return jumlah, err
}

func (r *keluargaRepository) DistinctKelurahanByKecamatan(kecamatanId string) ([]model.DistinctKelurahan, error) {
	var distinctsKelurahan []model.DistinctKelurahan

	var err = r.db.Distinct("k.kelurahanId, kl.nama").Table("keluargas as k").Joins("inner join kelurahans as kl on kl.id = k.kelurahanId").Where("k.kecamatanId = ?", kecamatanId).Scan(&distinctsKelurahan).Error

	return distinctsKelurahan, err
}

func (r *keluargaRepository) DistinctCountKelurahan(kecamatanId string) (map[string]int64, error) {
	var distinctsKelurahan, err = r.DistinctKelurahanByKecamatan(kecamatanId)

	var jumlah = make(map[string]int64)

	for _, distinctKelurahan := range distinctsKelurahan {
		var test, _ = r.CountJumlahKeluarga("kelurahanId", distinctKelurahan.KelurahanId)

		jumlah[distinctKelurahan.Nama] = test
	}

	return jumlah, err
}
