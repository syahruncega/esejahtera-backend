package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type InstansiOnProgramRepository interface {
	FindAll() ([]model.InstansiOnProgram, error)
	FindById(id int) (model.InstansiOnProgram, error)
	FindByInstansiId(instansiId int) ([]model.InstansiOnProgram, error)
	Create(instansiOnProgram model.InstansiOnProgram) (model.InstansiOnProgram, error)
	Update(instansiOnProgram model.InstansiOnProgram) (model.InstansiOnProgram, error)
	Delete(instansiOnProgram model.InstansiOnProgram) (model.InstansiOnProgram, error)
}

type instansiOnProgramRepository struct {
	db *gorm.DB
}

func NewInstansiOnProgramRepository(db *gorm.DB) *instansiOnProgramRepository {
	return &instansiOnProgramRepository{db}
}

func (r *instansiOnProgramRepository) FindAll() ([]model.InstansiOnProgram, error) {
	var instansiOnPrograms []model.InstansiOnProgram

	var err = r.db.Model(&instansiOnPrograms).Preload("Instansi").Preload("Program").Find(&instansiOnPrograms).Error

	return instansiOnPrograms, err
}

func (r *instansiOnProgramRepository) FindById(id int) (model.InstansiOnProgram, error) {
	var instansiOnProgram model.InstansiOnProgram

	var err = r.db.Model(&instansiOnProgram).Preload("Instansi").Preload("Program").Take(&instansiOnProgram, id).Error

	return instansiOnProgram, err
}

func (r *instansiOnProgramRepository) FindByInstansiId(instansiId int) ([]model.InstansiOnProgram, error) {
	var instansiOnPrograms []model.InstansiOnProgram

	var err = r.db.Where("instansiId = ?", instansiId).Model(&instansiOnPrograms).Preload("Instansi").Preload("Program").Find(&instansiOnPrograms).Error

	return instansiOnPrograms, err
}

func (r *instansiOnProgramRepository) Create(instansiOnProgram model.InstansiOnProgram) (model.InstansiOnProgram, error) {
	var err = r.db.Create(&instansiOnProgram).Error

	return instansiOnProgram, err
}

func (r *instansiOnProgramRepository) Update(instansiOnProgram model.InstansiOnProgram) (model.InstansiOnProgram, error) {
	var err = r.db.Model(&instansiOnProgram).Updates(model.InstansiOnProgram{
		InstansiId: instansiOnProgram.InstansiId,
		ProgramId:  instansiOnProgram.ProgramId,
	}).Error

	return instansiOnProgram, err
}

func (r *instansiOnProgramRepository) Delete(instansiOnProgram model.InstansiOnProgram) (model.InstansiOnProgram, error) {
	var err = r.db.Delete(&instansiOnProgram).Error

	return instansiOnProgram, err
}
