package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type DetailSubKegiatanRepository interface {
	FindAll() ([]model.Detail_Sub_Kegiatan, error)
	FindById(id int) (model.Detail_Sub_Kegiatan, error)
	FindAllRelation() ([]model.Detail_Sub_Kegiatan, error)
	Create(detailSubKegiatan model.Detail_Sub_Kegiatan) (model.Detail_Sub_Kegiatan, error)
	Update(detailSubKegiatan model.Detail_Sub_Kegiatan) (model.Detail_Sub_Kegiatan, error)
	Delete(detailSubKegiatan model.Detail_Sub_Kegiatan) (model.Detail_Sub_Kegiatan, error)
}

type detailSubKegiatanRepository struct {
	db *gorm.DB
}

func NewDetailSubKegiatanRepository(db *gorm.DB) *detailSubKegiatanRepository {
	return &detailSubKegiatanRepository{db}
}

func (r *detailSubKegiatanRepository) FindAll() ([]model.Detail_Sub_Kegiatan, error) {
	var detailSubKegiatans []model.Detail_Sub_Kegiatan

	var err = r.db.Find(&detailSubKegiatans).Error

	return detailSubKegiatans, err
}

func (r *detailSubKegiatanRepository) FindById(id int) (model.Detail_Sub_Kegiatan, error) {
	var detailSubKegiatan model.Detail_Sub_Kegiatan

	var err = r.db.Take(&detailSubKegiatan, id).Error

	return detailSubKegiatan, err
}

func (r *detailSubKegiatanRepository) FindAllRelation() ([]model.Detail_Sub_Kegiatan, error) {
	var detailSubKegiatans []model.Detail_Sub_Kegiatan

	var err = r.db.Model(&detailSubKegiatans).Preload("SubKegiatan").Find(&detailSubKegiatans).Error

	return detailSubKegiatans, err
}

func (r *detailSubKegiatanRepository) Create(detailSubKegiatan model.Detail_Sub_Kegiatan) (model.Detail_Sub_Kegiatan, error) {
	var err = r.db.Create(&detailSubKegiatan).Error

	return detailSubKegiatan, err
}

func (r *detailSubKegiatanRepository) Update(detailSubKegiatan model.Detail_Sub_Kegiatan) (model.Detail_Sub_Kegiatan, error) {
	var err = r.db.Model(&detailSubKegiatan).Updates(model.Detail_Sub_Kegiatan{
		FokusBelanja:     detailSubKegiatan.FokusBelanja,
		Indikator:        detailSubKegiatan.Indikator,
		Target:           detailSubKegiatan.Target,
		Satuan:           detailSubKegiatan.Satuan,
		PaguFokusBelanja: detailSubKegiatan.PaguFokusBelanja,
		Keterangan:       detailSubKegiatan.Keterangan,
		SubKegiatanId:    detailSubKegiatan.SubKegiatanId,
	}).Error

	return detailSubKegiatan, err
}

func (r *detailSubKegiatanRepository) Delete(detailSubKegiatan model.Detail_Sub_Kegiatan) (model.Detail_Sub_Kegiatan, error) {
	var err = r.db.Delete(&detailSubKegiatan).Error

	return detailSubKegiatan, err
}
