package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type IndividuVerifikasiRepository interface {
	FindAll() ([]model.IndividuVerifikasi, error)
	FindById(id int) (model.IndividuVerifikasi, error)
	FindByIdKeluarga(idKeluarga string) ([]model.IndividuVerifikasi, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.IndividuVerifikasi, error)
	Create(individuVerifikasi model.IndividuVerifikasi) (model.IndividuVerifikasi, error)
	Update(individuVerifikasi model.IndividuVerifikasi) (model.IndividuVerifikasi, error)
	Delete(individuVerifikasi model.IndividuVerifikasi) (model.IndividuVerifikasi, error)
}

type individuVerifikasiRepository struct {
	db *gorm.DB
}

func NewIndividuVerifikasiRepository(db *gorm.DB) *individuVerifikasiRepository {
	return &individuVerifikasiRepository{db}
}

func (r *individuVerifikasiRepository) FindAll() ([]model.IndividuVerifikasi, error) {
	var individuVerifikasis []model.IndividuVerifikasi

	var err = r.db.Limit(15).Model(&individuVerifikasis).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Preload("User").Preload("Mahasiswa").Find(&individuVerifikasis).Error

	return individuVerifikasis, err
}

func (r *individuVerifikasiRepository) FindById(id int) (model.IndividuVerifikasi, error) {
	var individuVerifikasi model.IndividuVerifikasi

	var err = r.db.Limit(15).Model(&individuVerifikasi).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Preload("User").Preload("Mahasiswa").Take(&individuVerifikasi, id).Error

	return individuVerifikasi, err
}

func (r *individuVerifikasiRepository) FindByIdKeluarga(idKeluarga string) ([]model.IndividuVerifikasi, error) {
	var individuVerifikasis []model.IndividuVerifikasi

	var err = r.db.Where("idKeluarga = ? ", idKeluarga).Limit(15).Model(&individuVerifikasis).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Preload("User").Preload("Mahasiswa").Find(&individuVerifikasis).Error

	return individuVerifikasis, err
}

func (r *individuVerifikasiRepository) FindBySearch(whereClause map[string]interface{}) ([]model.IndividuVerifikasi, error) {
	var individuVerifikasis []model.IndividuVerifikasi

	var err = r.db.Where(whereClause).Limit(15).Model(&individuVerifikasis).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Preload("User").Preload("Mahasiswa").Find(&individuVerifikasis).Error

	return individuVerifikasis, err
}

func (r *individuVerifikasiRepository) Create(individuVerifikasi model.IndividuVerifikasi) (model.IndividuVerifikasi, error) {
	var err = r.db.Create(&individuVerifikasi).Error

	return individuVerifikasi, err
}

func (r *individuVerifikasiRepository) Update(individuVerifikasi model.IndividuVerifikasi) (model.IndividuVerifikasi, error) {
	var err = r.db.Model(&individuVerifikasi).Updates(model.IndividuVerifikasi{
		IdKeluarga:         individuVerifikasi.IdKeluarga,
		ProvinsiId:         individuVerifikasi.ProvinsiId,
		KabupatenKotaId:    individuVerifikasi.KabupatenKotaId,
		KecamatanId:        individuVerifikasi.KecamatanId,
		KelurahanId:        individuVerifikasi.KelurahanId,
		DesilKesejahteraan: individuVerifikasi.DesilKesejahteraan,
		Alamat:             individuVerifikasi.Alamat,
		IdIndividu:         individuVerifikasi.IdIndividu,
		Nama:               individuVerifikasi.Nama,
		Nik:                individuVerifikasi.Nik,
		PadanDukcapil:      individuVerifikasi.PadanDukcapil,
		JenisKelamin:       individuVerifikasi.JenisKelamin,
		Hubungan:           individuVerifikasi.Hubungan,
		TanggalLahir:       individuVerifikasi.TanggalLahir,
		StatusKawin:        individuVerifikasi.StatusKawin,
		Pekerjaan:          individuVerifikasi.Pekerjaan,
		Pendidikan:         individuVerifikasi.Pendidikan,
		KategoriUsia:       individuVerifikasi.KategoriUsia,
		PenerimaBPNT:       individuVerifikasi.PenerimaBPNT,
		PenerimaBPUM:       individuVerifikasi.PenerimaBPUM,
		PenerimaBST:        individuVerifikasi.PenerimaBST,
		PenerimaPKH:        individuVerifikasi.PenerimaPKH,
		PenerimaSembako:    individuVerifikasi.PenerimaSembako,
		PenerimaLainnya:    individuVerifikasi.PenerimaLainnya,
		StatusResponden:    individuVerifikasi.StatusResponden,
		UserId:             individuVerifikasi.UserId,
		MahasiswaId:        individuVerifikasi.MahasiswaId,
		UrlBukti:           individuVerifikasi.UrlBukti,
	}).Error

	return individuVerifikasi, err
}

func (r *individuVerifikasiRepository) Delete(individuVerifikasi model.IndividuVerifikasi) (model.IndividuVerifikasi, error) {
	var err = r.db.Delete(&individuVerifikasi).Error

	return individuVerifikasi, err
}
