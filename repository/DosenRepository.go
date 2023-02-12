package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type DosenRepository interface {
	FindAll() ([]model.Dosen, error)
	FindById(id int) (model.Dosen, error)
	FindByUserId(userId int) (model.Dosen, error)
	FindAllRelation() ([]model.Dosen, error)
	Create(dosen model.Dosen) (model.Dosen, error)
	Update(dosen model.Dosen) (model.Dosen, error)
	Delete(dosen model.Dosen) (model.Dosen, error)
}

type dosenRepository struct {
	db *gorm.DB
}

func NewDosenRepository(db *gorm.DB) *dosenRepository {
	return &dosenRepository{db}
}

func (r *dosenRepository) FindAll() ([]model.Dosen, error) {
	var dosens []model.Dosen

	var err = r.db.Find(&dosens).Error

	return dosens, err
}

func (r *dosenRepository) FindById(id int) (model.Dosen, error) {
	var dosen model.Dosen

	var err = r.db.Take(&dosen, id).Error

	return dosen, err
}

func (r *dosenRepository) FindByUserId(userId int) (model.Dosen, error) {
	var dosen model.Dosen

	var err = r.db.Where("userId = ? ", userId).Take(&dosen).Error

	return dosen, err
}

func (r *dosenRepository) FindAllRelation() ([]model.Dosen, error) {
	var dosens []model.Dosen

	var err = r.db.Model(&dosens).Preload("User").Find(&dosens).Error

	return dosens, err
}

func (r *dosenRepository) Create(dosen model.Dosen) (model.Dosen, error) {
	var err = r.db.Create(&dosen).Error

	return dosen, err
}

func (r *dosenRepository) Update(dosen model.Dosen) (model.Dosen, error) {
	var err = r.db.Model(&dosen).Updates(model.Dosen{
		UserId:      dosen.UserId,
		NamaLengkap: dosen.NamaLengkap,
		Universitas: dosen.Universitas,
	}).Error

	return dosen, err
}

func (r *dosenRepository) Delete(dosen model.Dosen) (model.Dosen, error) {
	var err = r.db.Delete(&dosen).Error

	return dosen, err
}
