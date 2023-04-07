package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type RencanaProgramRepository interface {
	FindAll() ([]model.RencanaProgram, error)
	FindById(id int) (model.RencanaProgram, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.RencanaProgram, error)
	Create(rencanaProgram model.RencanaProgram) (model.RencanaProgram, error)
	Update(rencanaProgram model.RencanaProgram) (model.RencanaProgram, error)
	Delete(rencanaProgram model.RencanaProgram) (model.RencanaProgram, error)
}

type rencanaProgramRepository struct {
	db *gorm.DB
}

func NewRencanaProgramRepository(db *gorm.DB) *rencanaProgramRepository {
	return &rencanaProgramRepository{db}
}

func (r *rencanaProgramRepository) FindAll() ([]model.RencanaProgram, error) {
	var rencanaPrograms []model.RencanaProgram

	var err = r.db.Model(&rencanaPrograms).Preload("Instansi").Preload("Program").Find(&rencanaPrograms).Error

	return rencanaPrograms, err
}

func (r *rencanaProgramRepository) FindById(id int) (model.RencanaProgram, error) {
	var rencanaProgram model.RencanaProgram

	var err = r.db.Model(&rencanaProgram).Preload("Instansi").Preload("Program").Take(&rencanaProgram, id).Error

	return rencanaProgram, err
}

func (r *rencanaProgramRepository) FindBySearch(whereClause map[string]interface{}) ([]model.RencanaProgram, error) {
	var rencanaPrograms []model.RencanaProgram

	var err = r.db.Where(whereClause).Model(&rencanaPrograms).Preload("Instansi").Preload("Program").Find(&rencanaPrograms).Error

	return rencanaPrograms, err
}

func (r *rencanaProgramRepository) Create(rencanaProgram model.RencanaProgram) (model.RencanaProgram, error) {
	var err = r.db.Create(&rencanaProgram).Error

	return rencanaProgram, err
}

func (r *rencanaProgramRepository) Update(rencanaProgram model.RencanaProgram) (model.RencanaProgram, error) {
	var err = r.db.Model(&rencanaProgram).Updates(model.RencanaProgram{
		InstansiId:  rencanaProgram.InstansiId,
		ProgramId:   rencanaProgram.ProgramId,
		PaguProgram: rencanaProgram.PaguProgram,
		Tipe:        rencanaProgram.Tipe,
		Tahun:       rencanaProgram.Tahun,
	}).Error

	return rencanaProgram, err
}

func (r *rencanaProgramRepository) Delete(rencanaProgram model.RencanaProgram) (model.RencanaProgram, error) {
	var err = r.db.Delete(&rencanaProgram).Error

	return rencanaProgram, err
}
