package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type KebijakanRepository interface {
	FindAll() ([]model.Kebijakan, error)
	FindById(id int) (model.Kebijakan, error)
	FindByProgramId(programId int) ([]model.Kebijakan, error)
	Create(kebijakan model.Kebijakan) (model.Kebijakan, error)
	Update(kebijakan model.Kebijakan) (model.Kebijakan, error)
	Delete(kebijakan model.Kebijakan) (model.Kebijakan, error)
}

type kebijakanRepository struct {
	db *gorm.DB
}

func NewKebijakanRepository(db *gorm.DB) *kebijakanRepository {
	return &kebijakanRepository{db}
}

func (r *kebijakanRepository) FindAll() ([]model.Kebijakan, error) {
	var kebijakans []model.Kebijakan

	var err = r.db.Model(&kebijakans).Preload("Program").Find(&kebijakans).Error

	return kebijakans, err
}

func (r *kebijakanRepository) FindById(id int) (model.Kebijakan, error) {
	var kebijakan model.Kebijakan

	var err = r.db.Model(&kebijakan).Preload("Program").Take(&kebijakan, id).Error

	return kebijakan, err
}

func (r *kebijakanRepository) FindByProgramId(programId int) ([]model.Kebijakan, error) {
	var kebijakans []model.Kebijakan

	var err = r.db.Where("programId = ? ", programId).Model(&kebijakans).Preload("Program").Find(&kebijakans).Error

	return kebijakans, err
}

func (r *kebijakanRepository) Create(kebijakan model.Kebijakan) (model.Kebijakan, error) {
	var err = r.db.Create(&kebijakan).Error

	return kebijakan, err
}

func (r *kebijakanRepository) Update(kebijakan model.Kebijakan) (model.Kebijakan, error) {
	var err = r.db.Model(&kebijakan).Updates(model.Kebijakan{
		ProgramId:     kebijakan.ProgramId,
		NamaKebijakan: kebijakan.NamaKebijakan,
	}).Error

	return kebijakan, err
}

func (r *kebijakanRepository) Delete(kebijakan model.Kebijakan) (model.Kebijakan, error) {
	var err = r.db.Delete(&kebijakan).Error

	return kebijakan, err
}
