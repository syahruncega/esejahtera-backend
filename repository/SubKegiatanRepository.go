package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type SubKegiatanRepository interface {
	FindAll() ([]model.SubKegiatan, error)
	FindById(id int) (model.SubKegiatan, error)
	Create(subKegiatan model.SubKegiatan) (model.SubKegiatan, error)
	Update(subKegiatan model.SubKegiatan) (model.SubKegiatan, error)
	Delete(subKegiatan model.SubKegiatan) (model.SubKegiatan, error)
}

type subKegiatanRepository struct {
	db *gorm.DB
}

func NewSubKegiatanRepository(db *gorm.DB) *subKegiatanRepository {
	return &subKegiatanRepository{db}
}

func (r *subKegiatanRepository) FindAll() ([]model.SubKegiatan, error) {
	var subKegiatans []model.SubKegiatan

	var err = r.db.Find(&subKegiatans).Error

	return subKegiatans, err
}

func (r *subKegiatanRepository) FindById(id int) (model.SubKegiatan, error) {
	var subKegiatan model.SubKegiatan

	var err = r.db.Take(&subKegiatan, id).Error

	return subKegiatan, err
}

func (r *subKegiatanRepository) Create(subKegiatan model.SubKegiatan) (model.SubKegiatan, error) {
	var err = r.db.Create(&subKegiatan).Error

	return subKegiatan, err
}

func (r *subKegiatanRepository) Update(subKegiatan model.SubKegiatan) (model.SubKegiatan, error) {
	var err = r.db.Model(&subKegiatan).Updates(model.SubKegiatan{
		Tahun:           subKegiatan.Tahun,
		KodeSubKegiatan: subKegiatan.KodeSubKegiatan,
		NamaSubKegiatan: subKegiatan.NamaSubKegiatan,
	}).Error

	return subKegiatan, err
}

func (r *subKegiatanRepository) Delete(subKegiatan model.SubKegiatan) (model.SubKegiatan, error) {
	var err = r.db.Delete(&subKegiatan).Error

	return subKegiatan, err
}
