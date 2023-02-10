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
	CountJumlahIndividu(places string, placesId string) (int64, error)
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

func (r *individuRepository) CountJumlahIndividu(places string, placesId string) (int64, error) {
	var count int64

	var err = r.db.Where(places+"= ?", placesId).Table("individus").Select("count(*)").Count(&count).Error

	return count, err
}

func (r *individuRepository) CountJumlahDesil(places string, placesId string, desil string) (int64, error) {
	var count int64

	var err = r.db.Where(places+"= ? and desilKesejahteraan = ?", placesId, desil).Table("individus").Select("count(*)").Count(&count).Error

	return count, err
}

func (r *individuRepository) CountJumlahPenerima(places string, placesId string, penerima string, penerimaValue string) (int64, error) {
	var count int64

	var err = r.db.Where(places+"= ? and "+penerima+" = ?", placesId, penerimaValue).Table("individus").Select("count(*)").Count(&count).Error

	return count, err
}

func (r *individuRepository) DistinctKabupatenKota() ([]model.DistinctKabupatenKota, error) {
	var distinctKabupatenKota []model.DistinctKabupatenKota

	var err = r.db.Distinct("i.kabupatenKotaId, kb.nama").Table("individus as i").Joins("inner join kabupaten_kota as kb on kb.id = i.kabupatenKotaId").Scan(&distinctKabupatenKota).Error

	return distinctKabupatenKota, err
}

func (r *individuRepository) DistinctCountKabupatenKota() (map[string]int64, error) {
	var distinctsKabupatenKota, err = r.DistinctKabupatenKota()

	var jumlah = make(map[string]int64)

	for _, distinctKabupatenKota := range distinctsKabupatenKota {
		var test, _ = r.CountJumlahIndividu("kabupatenKotaId", distinctKabupatenKota.KabupatenKotaId)

		jumlah[distinctKabupatenKota.Nama] = test
	}

	return jumlah, err
}

func (r *individuRepository) DistinctKecamatan() ([]model.DistinctKecamatan, error) {
	var distinctsKecamatan []model.DistinctKecamatan

	var err = r.db.Distinct("i.kecamatanId, kc.nama").Table("individus as i").Joins("inner join kecamatans as kc on kc.id = i.kecamatanId").Scan(&distinctsKecamatan).Error

	return distinctsKecamatan, err
}

func (r *individuRepository) DistinctKecamatanByKabupatenKota(kabupatenKotaId string) ([]model.DistinctKecamatan, error) {
	var distinctsKecamatan []model.DistinctKecamatan

	var err = r.db.Distinct("i.kecamatanId, kc.nama").Table("individus as i").Joins("inner join kecamatans as kc on kc.id = i.kecamatanId").Where("i.kabupatenKotaId = ?", kabupatenKotaId).Scan(&distinctsKecamatan).Error

	return distinctsKecamatan, err
}

func (r *individuRepository) DistinctCountKecamatan(kabupatenKotaId string) (map[string]int64, error) {
	var distinctsKecamatan, err = r.DistinctKecamatanByKabupatenKota(kabupatenKotaId)

	var jumlah = make(map[string]int64)

	for _, distinctKecamatan := range distinctsKecamatan {
		var test, _ = r.CountJumlahIndividu("kecamatanId", distinctKecamatan.KecamatanId)

		jumlah[distinctKecamatan.Nama] = test
	}

	return jumlah, err
}

func (r *individuRepository) DistinctKelurahanByKecamatan(kecamatanId string) ([]model.DistinctKelurahan, error) {
	var distinctsKelurahan []model.DistinctKelurahan

	var err = r.db.Distinct("i.kelurahanId, kl.nama").Table("individus as i").Joins("inner join kelurahans as kl on kl.id = i.kelurahanId").Where("i.kecamatanId = ?", kecamatanId).Scan(&distinctsKelurahan).Error

	return distinctsKelurahan, err
}

func (r *individuRepository) DistinctCountKelurahan(kecamatanId string) (map[string]int64, error) {
	var distinctsKelurahan, err = r.DistinctKelurahanByKecamatan(kecamatanId)

	var jumlah = make(map[string]int64)

	for _, distinctKelurahan := range distinctsKelurahan {
		var test, _ = r.CountJumlahIndividu("kelurahanId", distinctKelurahan.KelurahanId)

		jumlah[distinctKelurahan.Nama] = test
	}

	return jumlah, err
}
