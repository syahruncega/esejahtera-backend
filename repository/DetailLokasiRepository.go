package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type DetailLokasiRepository interface {
	FindAll() ([]model.Detail_Lokasi, error)
	FindById(id int) (model.Detail_Lokasi, error)
	FindAllRelation() ([]model.Detail_Lokasi, error)
	Create(detailLokasi model.Detail_Lokasi) (model.Detail_Lokasi, error)
	Update(detailLokasi model.Detail_Lokasi) (model.Detail_Lokasi, error)
	Delete(detailLokasi model.Detail_Lokasi) (model.Detail_Lokasi, error)
}

type detailLokasiRepository struct {
	db *gorm.DB
}

func NewDetailLokasiRepository(db *gorm.DB) *detailLokasiRepository {
	return &detailLokasiRepository{db}
}

func (r *detailLokasiRepository) FindAll() ([]model.Detail_Lokasi, error) {
	var detailLokasis []model.Detail_Lokasi

	var err = r.db.Find(&detailLokasis).Error

	return detailLokasis, err
}

func (r *detailLokasiRepository) FindById(id int) (model.Detail_Lokasi, error) {
	var detailLokasi model.Detail_Lokasi

	var err = r.db.Take(&detailLokasi, id).Error

	return detailLokasi, err
}

func (r *detailLokasiRepository) FindAllRelation() ([]model.Detail_Lokasi, error) {
	var detailLokasis []model.Detail_Lokasi

	var err = r.db.Model(&detailLokasis).Preload("DetailSubKegiatan").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Find(&detailLokasis).Error

	return detailLokasis, err
}

func (r *detailLokasiRepository) Create(detailLokasi model.Detail_Lokasi) (model.Detail_Lokasi, error) {
	var err = r.db.Create(&detailLokasi).Error

	return detailLokasi, err
}

func (r *detailLokasiRepository) Update(detailLokasi model.Detail_Lokasi) (model.Detail_Lokasi, error) {
	var err = r.db.Model(&detailLokasi).Updates(model.Detail_Lokasi{
		KabupatenKotaId:     detailLokasi.KabupatenKotaId,
		KecamatanId:         detailLokasi.KecamatanId,
		KelurahanId:         detailLokasi.KelurahanId,
		DetailSubKegiatanId: detailLokasi.DetailSubKegiatanId,
	}).Error

	return detailLokasi, err
}

func (r *detailLokasiRepository) Delete(detailLokasi model.Detail_Lokasi) (model.Detail_Lokasi, error) {
	var err = r.db.Delete(&detailLokasi).Error

	return detailLokasi, err
}
