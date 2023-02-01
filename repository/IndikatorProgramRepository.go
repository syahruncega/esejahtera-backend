package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type IndikatorProgramRepository interface {
	FindAll() ([]model.IndikatorProgram, error)
	FindById(id int) (model.IndikatorProgram, error)
	Create(indikatorProgram model.IndikatorProgram) (model.IndikatorProgram, error)
	Update(indikatorProgram model.IndikatorProgram) (model.IndikatorProgram, error)
	Delete(indikatorProgram model.IndikatorProgram) (model.IndikatorProgram, error)
}

type indikatorProgramRepository struct {
	db *gorm.DB
}

func NewIndikatorProgramRepository(db *gorm.DB) *indikatorProgramRepository {
	return &indikatorProgramRepository{db}
}

func (r *indikatorProgramRepository) FindAll() ([]model.IndikatorProgram, error) {
	var indikatorPrograms []model.IndikatorProgram

	var err = r.db.Model(&indikatorPrograms).Preload("Instansi").Preload("Program").Find(&indikatorPrograms).Error

	return indikatorPrograms, err
}

func (r *indikatorProgramRepository) FindById(id int) (model.IndikatorProgram, error) {
	var indikatorProgram model.IndikatorProgram

	var err = r.db.Model(&indikatorProgram).Preload("Instansi").Preload("Program").Take(&indikatorProgram, id).Error

	return indikatorProgram, err
}

func (r *indikatorProgramRepository) Create(indikatorProgram model.IndikatorProgram) (model.IndikatorProgram, error) {
	var err = r.db.Create(&indikatorProgram).Error

	return indikatorProgram, err
}

func (r *indikatorProgramRepository) Update(indikatorProgram model.IndikatorProgram) (model.IndikatorProgram, error) {
	var err = r.db.Model(&indikatorProgram).Select("Sasaran", "IndikatorSasaran", "Kebijakan", "IndikatorKinerjaProgram", "PaguProgram", "InstansiId", "ProgramId").Updates(model.IndikatorProgram{
		Sasaran:                 indikatorProgram.Sasaran,
		IndikatorSasaran:        indikatorProgram.IndikatorSasaran,
		Kebijakan:               indikatorProgram.Kebijakan,
		IndikatorKinerjaProgram: indikatorProgram.IndikatorKinerjaProgram,
		PaguProgram:             indikatorProgram.PaguProgram,
		InstansiId:              indikatorProgram.InstansiId,
		ProgramId:               indikatorProgram.ProgramId,
	}).Error
	// var err = r.db.Model(&indikatorProgram).Updates(model.IndikatorProgram{
	// 	Sasaran:                 indikatorProgram.Sasaran,
	// 	IndikatorSasaran:        indikatorProgram.IndikatorSasaran,
	// 	Kebijakan:               indikatorProgram.Kebijakan,
	// 	IndikatorKinerjaProgram: indikatorProgram.IndikatorKinerjaProgram,
	// 	PaguProgram:             indikatorProgram.PaguProgram,
	// 	InstansiId:              indikatorProgram.InstansiId,
	// 	ProgramId:               indikatorProgram.ProgramId,
	// }).Error

	return indikatorProgram, err
}

func (r *indikatorProgramRepository) Delete(indikatorProgram model.IndikatorProgram) (model.IndikatorProgram, error) {
	var err = r.db.Delete(&indikatorProgram).Error

	return indikatorProgram, err
}
