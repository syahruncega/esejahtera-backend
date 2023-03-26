package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type IndikatorSasaranRepository interface {
	FindAll() ([]model.IndikatorSasaran, error)
	FindById(id int) (model.IndikatorSasaran, error)
	FindByProgramId(programId int) ([]model.IndikatorSasaran, error)
	Create(indikatorSasaran model.IndikatorSasaran) (model.IndikatorSasaran, error)
	Update(indikatorSasaran model.IndikatorSasaran) (model.IndikatorSasaran, error)
	Delete(indikatorSasaran model.IndikatorSasaran) (model.IndikatorSasaran, error)
}

type indikatorSasaranRepository struct {
	db *gorm.DB
}

func NewIndikatorSasaranRepository(db *gorm.DB) *indikatorSasaranRepository {
	return &indikatorSasaranRepository{db}
}

func (r *indikatorSasaranRepository) FindAll() ([]model.IndikatorSasaran, error) {
	var indikatorSasarans []model.IndikatorSasaran

	var err = r.db.Model(&indikatorSasarans).Preload("Program").Find(&indikatorSasarans).Error

	return indikatorSasarans, err
}

func (r *indikatorSasaranRepository) FindById(id int) (model.IndikatorSasaran, error) {
	var indikatorSasaran model.IndikatorSasaran

	var err = r.db.Model(&indikatorSasaran).Preload("Program").Take(&indikatorSasaran, id).Error

	return indikatorSasaran, err
}

func (r *indikatorSasaranRepository) FindByProgramId(programId int) ([]model.IndikatorSasaran, error) {
	var indikatorSasarans []model.IndikatorSasaran

	var err = r.db.Where("programId = ?", programId).Model(&indikatorSasarans).Preload("Program").Find(&indikatorSasarans).Error

	return indikatorSasarans, err
}

func (r *indikatorSasaranRepository) Create(indikatorSasaran model.IndikatorSasaran) (model.IndikatorSasaran, error) {
	var err = r.db.Create(&indikatorSasaran).Error

	return indikatorSasaran, err
}

func (r *indikatorSasaranRepository) Update(indikatorSasaran model.IndikatorSasaran) (model.IndikatorSasaran, error) {
	var err = r.db.Model(&indikatorSasaran).Updates(model.IndikatorSasaran{
		ProgramId:            indikatorSasaran.ProgramId,
		NamaIndikatorSasaran: indikatorSasaran.NamaIndikatorSasaran,
	}).Error

	return indikatorSasaran, err
}

func (r *indikatorSasaranRepository) Delete(indikatorSasaran model.IndikatorSasaran) (model.IndikatorSasaran, error) {
	var err = r.db.Delete(&indikatorSasaran).Error

	return indikatorSasaran, err
}
