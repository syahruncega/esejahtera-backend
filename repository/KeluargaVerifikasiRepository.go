package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type KeluargaVerifikasiRepository interface {
	FindAll() ([]model.KeluargaVerifikasi, error)
	FindById(id int) (model.KeluargaVerifikasi, error)
	FindByIdKeluarga(idKeluarga string) (model.KeluargaVerifikasi, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.KeluargaVerifikasi, error)
	Create(keluargaVerifikasi model.KeluargaVerifikasi) (model.KeluargaVerifikasi, error)
	Update(keluargaVerifikasi model.KeluargaVerifikasi) (model.KeluargaVerifikasi, error)
	Delete(keluargaVerifikasi model.KeluargaVerifikasi) (model.KeluargaVerifikasi, error)
}

type keluargaVerifikasiRepository struct {
	db *gorm.DB
}

func NewKeluargaVerifikasiRepository(db *gorm.DB) *keluargaVerifikasiRepository {
	return &keluargaVerifikasiRepository{db}
}

func (r *keluargaVerifikasiRepository) FindAll() ([]model.KeluargaVerifikasi, error) {
	var keluargaVerifikasis []model.KeluargaVerifikasi

	var err = r.db.Model(&keluargaVerifikasis).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("User").Preload("Mahasiswa").Preload("Kelurahan").Find(&keluargaVerifikasis).Error

	return keluargaVerifikasis, err
}

func (r *keluargaVerifikasiRepository) FindById(id int) (model.KeluargaVerifikasi, error) {
	var keluargaVerifikasi model.KeluargaVerifikasi

	var err = r.db.Model(&keluargaVerifikasi).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Preload("User").Take(&keluargaVerifikasi, id).Error

	return keluargaVerifikasi, err
}

func (r *keluargaVerifikasiRepository) FindByIdKeluarga(idKeluarga string) (model.KeluargaVerifikasi, error) {
	var keluargaVerifikasi model.KeluargaVerifikasi

	var err = r.db.Where("idKeluarga = ? ", idKeluarga).Model(&keluargaVerifikasi).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Preload("User").Take(&keluargaVerifikasi).Error

	return keluargaVerifikasi, err
}

func (r *keluargaVerifikasiRepository) FindBySearch(whereClause map[string]interface{}) ([]model.KeluargaVerifikasi, error) {
	var keluargaVerifikasis []model.KeluargaVerifikasi

	var err = r.db.Where(whereClause).Model(&keluargaVerifikasis).Preload("Provinsi").Preload("KabupatenKota").Preload("Kecamatan").Preload("User").Preload("Mahasiswa").Preload("Kelurahan").Find(&keluargaVerifikasis).Error

	return keluargaVerifikasis, err
}

func (r *keluargaVerifikasiRepository) Create(keluargaVerifikasi model.KeluargaVerifikasi) (model.KeluargaVerifikasi, error) {
	var err = r.db.Create(&keluargaVerifikasi).Error

	return keluargaVerifikasi, err
}

func (r *keluargaVerifikasiRepository) Update(keluargaVerifikasi model.KeluargaVerifikasi) (model.KeluargaVerifikasi, error) {
	var err = r.db.Model(&keluargaVerifikasi).Updates(model.KeluargaVerifikasi{
		IdKeluarga:             keluargaVerifikasi.IdKeluarga,
		ProvinsiId:             keluargaVerifikasi.ProvinsiId,
		KabupatenKotaId:        keluargaVerifikasi.KabupatenKotaId,
		KecamatanId:            keluargaVerifikasi.KecamatanId,
		KelurahanId:            keluargaVerifikasi.KelurahanId,
		DesilKesejahteraan:     keluargaVerifikasi.DesilKesejahteraan,
		Alamat:                 keluargaVerifikasi.Alamat,
		KepalaKeluarga:         keluargaVerifikasi.KepalaKeluarga,
		Nik:                    keluargaVerifikasi.Nik,
		PadanDukcapil:          keluargaVerifikasi.PadanDukcapil,
		JenisKelamin:           keluargaVerifikasi.JenisKelamin,
		TanggalLahir:           keluargaVerifikasi.TanggalLahir,
		Pekerjaan:              keluargaVerifikasi.Pekerjaan,
		Pendidikan:             keluargaVerifikasi.Pendidikan,
		KepemilikanRumah:       keluargaVerifikasi.KepemilikanRumah,
		Simpanan:               keluargaVerifikasi.Simpanan,
		JenisAtap:              keluargaVerifikasi.JenisAtap,
		JenisDinding:           keluargaVerifikasi.JenisDinding,
		JenisLantai:            keluargaVerifikasi.JenisLantai,
		SumberPenerangan:       keluargaVerifikasi.SumberPenerangan,
		BahanBakarMemasak:      keluargaVerifikasi.BahanBakarMemasak,
		SumberAirMinum:         keluargaVerifikasi.SumberAirMinum,
		FasilitasBuangAirBesar: keluargaVerifikasi.FasilitasBuangAirBesar,
		PenerimaBPNT:           keluargaVerifikasi.PenerimaBPNT,
		PenerimaBPUM:           keluargaVerifikasi.PenerimaBPUM,
		PenerimaBST:            keluargaVerifikasi.PenerimaBST,
		PenerimaPKH:            keluargaVerifikasi.PenerimaPKH,
		PenerimaSembako:        keluargaVerifikasi.PenerimaSembako,
		PenerimaLainnya:        keluargaVerifikasi.PenerimaLainnya,
		StatusResponden:        keluargaVerifikasi.StatusResponden,
		UserId:                 keluargaVerifikasi.UserId,
		MahasiswaId:            keluargaVerifikasi.MahasiswaId,
		UrlBukti:               keluargaVerifikasi.UrlBukti,
	}).Error

	return keluargaVerifikasi, err
}

func (r *keluargaVerifikasiRepository) Delete(keluargaVerifikasi model.KeluargaVerifikasi) (model.KeluargaVerifikasi, error) {
	var err = r.db.Delete(&keluargaVerifikasi).Error

	return keluargaVerifikasi, err
}
