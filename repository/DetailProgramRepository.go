package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type DetailProgramRepository interface {
	FindAll() ([]model.DetailProgram, error)
	FindById(id int) (model.DetailProgram, error)
	FindByInstansiId(instansiId string) ([]model.DetailProgram, error)
	Create(detailProgram model.DetailProgram) (model.DetailProgram, error)
	Update(detailProgram model.DetailProgram) (model.DetailProgram, error)
	Delete(detailProgram model.DetailProgram) (model.DetailProgram, error)
}

type detailProgramRepository struct {
	db *gorm.DB
}

func NewDetailProgramRepository(db *gorm.DB) *detailProgramRepository {
	return &detailProgramRepository{db}
}

func (r *detailProgramRepository) FindAll() ([]model.DetailProgram, error) {
	var detailPrograms []model.DetailProgram

	var err = r.db.Model(&detailPrograms).Preload("Instansi").Preload("Program").Find(&detailPrograms).Error

	return detailPrograms, err
}

func (r *detailProgramRepository) FindById(id int) (model.DetailProgram, error) {
	var detailProgram model.DetailProgram

	var err = r.db.Model(&detailProgram).Preload("Instansi").Preload("Program").Take(&detailProgram, id).Error

	return detailProgram, err
}

func (r *detailProgramRepository) FindByInstansiId(instansiId string) ([]model.DetailProgram, error) {
	var detailPrograms []model.DetailProgram

	var err = r.db.Where("instansiId = ?", instansiId).Model(&detailPrograms).Preload("Instansi").Preload("Program").Find(&detailPrograms).Error

	return detailPrograms, err
}

func (r *detailProgramRepository) Create(detailProgram model.DetailProgram) (model.DetailProgram, error) {
	var err = r.db.Create(&detailProgram).Error

	return detailProgram, err
}

func (r *detailProgramRepository) Update(detailProgram model.DetailProgram) (model.DetailProgram, error) {
	var err = r.db.Model(&detailProgram).Updates(model.DetailProgram{
		InstansiId:  detailProgram.InstansiId,
		ProgramId:   detailProgram.ProgramId,
		PaguProgram: detailProgram.PaguProgram,
	}).Error

	return detailProgram, err
}

func (r *detailProgramRepository) Delete(detailProgram model.DetailProgram) (model.DetailProgram, error) {
	var err = r.db.Delete(&detailProgram).Error

	return detailProgram, err
}
