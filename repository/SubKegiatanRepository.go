package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type SubKegiatanRepository interface {
	FindAll() ([]model.Sub_Kegiatan, error)
	FindById(id int) (model.Sub_Kegiatan, error)
	FindAllRelation() ([]model.Sub_Kegiatan, error)
	Create(sub_kegiatan model.Sub_Kegiatan) (model.Sub_Kegiatan, error)
	Update(sub_kegiatan model.Sub_Kegiatan) (model.Sub_Kegiatan, error)
	Delete(sub_kegiatan model.Sub_Kegiatan) (model.Sub_Kegiatan, error)
}

type subKegiatanRepository struct {
	db *gorm.DB
}

func NewSubKegiatanRepository(db *gorm.DB) *subKegiatanRepository {
	return &subKegiatanRepository{db}
}

func (r *subKegiatanRepository) FindAll() ([]model.Sub_Kegiatan, error) {
	var subKegiatans []model.Sub_Kegiatan

	var err = r.db.Find(&subKegiatans).Error

	return subKegiatans, err
}

func (r *subKegiatanRepository) FindById(id int) (model.Sub_Kegiatan, error) {
	var subKegiatan model.Sub_Kegiatan

	var err = r.db.Take(&subKegiatan, id).Error

	return subKegiatan, err
}

func (r *subKegiatanRepository) FindAllRelation() ([]model.Sub_Kegiatan, error) {
	var subKegiatans []model.Sub_Kegiatan

	var err = r.db.Model(&subKegiatans).Preload("Kegiatan").Find(&subKegiatans).Error

	return subKegiatans, err
}

func (r *subKegiatanRepository) Create(subKegiatan model.Sub_Kegiatan) (model.Sub_Kegiatan, error) {
	var err = r.db.Create(&subKegiatan).Error

	return subKegiatan, err
}

func (r *subKegiatanRepository) Update(subKegiatan model.Sub_Kegiatan) (model.Sub_Kegiatan, error) {
	var err = r.db.Model(&subKegiatan).Select("NamaSubKegiatan", "IndikatorKinerjaSubKegiatan", "PaguSubKegiatan", "KegiatanId").Updates(model.Sub_Kegiatan{
		NamaSubKegiatan:             subKegiatan.NamaSubKegiatan,
		IndikatorKinerjaSubKegiatan: subKegiatan.IndikatorKinerjaSubKegiatan,
		PaguSubKegiatan:             subKegiatan.PaguSubKegiatan,
		KegiatanId:                  subKegiatan.KegiatanId,
	}).Error

	return subKegiatan, err
}

func (r *subKegiatanRepository) Delete(subKegiatan model.Sub_Kegiatan) (model.Sub_Kegiatan, error) {
	var err = r.db.Delete(&subKegiatan).Error

	return subKegiatan, err
}
