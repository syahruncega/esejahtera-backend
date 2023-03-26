package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type SasaranRepository interface {
	FindAll() ([]model.Sasaran, error)
	FindById(id int) (model.Sasaran, error)
	FindByProgramId(programId int) ([]model.Sasaran, error)
	Create(sasaran model.Sasaran) (model.Sasaran, error)
	Update(sasaran model.Sasaran) (model.Sasaran, error)
	Delete(sasaran model.Sasaran) (model.Sasaran, error)
}

type sasaranRepository struct {
	db *gorm.DB
}

func NewSasaranRepository(db *gorm.DB) *sasaranRepository {
	return &sasaranRepository{db}
}

func (r *sasaranRepository) FindAll() ([]model.Sasaran, error) {
	var sasarans []model.Sasaran

	var err = r.db.Model(&sasarans).Preload("Program").Find(&sasarans).Error

	return sasarans, err
}

func (r *sasaranRepository) FindById(id int) (model.Sasaran, error) {
	var sasaran model.Sasaran

	var err = r.db.Model(&sasaran).Preload("Program").Take(&sasaran, id).Error

	return sasaran, err
}

func (r *sasaranRepository) FindByProgramId(programId int) ([]model.Sasaran, error) {
	var sasarans []model.Sasaran

	var err = r.db.Where("programId = ?", programId).Preload("Program").Find(&sasarans).Error

	return sasarans, err
}

func (r *sasaranRepository) Create(sasaran model.Sasaran) (model.Sasaran, error) {
	var err = r.db.Create(&sasaran).Error

	return sasaran, err
}

func (r *sasaranRepository) Update(sasaran model.Sasaran) (model.Sasaran, error) {
	var err = r.db.Model(&sasaran).Updates(model.Sasaran{
		ProgramId:   sasaran.ProgramId,
		NamaSasaran: sasaran.NamaSasaran,
	}).Error

	return sasaran, err
}

func (r *sasaranRepository) Delete(sasaran model.Sasaran) (model.Sasaran, error) {
	var err = r.db.Delete(&sasaran).Error

	return sasaran, err
}
