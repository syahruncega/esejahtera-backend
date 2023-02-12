package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type AnalisRepository interface {
	FindAll() ([]model.Analis, error)
	FindById(id int) (model.Analis, error)
	FindByUserId(userId int) (model.Analis, error)
	FindAllRelation() ([]model.Analis, error)
	Create(analis model.Analis) (model.Analis, error)
	Update(analis model.Analis) (model.Analis, error)
	Delete(analis model.Analis) (model.Analis, error)
}

type analisRepository struct {
	db *gorm.DB
}

func NewAnalisRepository(db *gorm.DB) *analisRepository {
	return &analisRepository{db}
}

func (r *analisRepository) FindAll() ([]model.Analis, error) {
	var analiss []model.Analis

	var err = r.db.Find(&analiss).Error

	return analiss, err
}

func (r *analisRepository) FindById(id int) (model.Analis, error) {
	var analis model.Analis

	var err = r.db.Take(&analis, id).Error

	return analis, err
}

func (r *analisRepository) FindByUserId(userId int) (model.Analis, error) {
	var analis model.Analis

	var err = r.db.Where("userId = ?", userId).Take(&analis).Error

	return analis, err
}

func (r *analisRepository) FindAllRelation() ([]model.Analis, error) {
	var analiss []model.Analis

	var err = r.db.Model(&analiss).Preload("User").Find(&analiss).Error

	return analiss, err
}

func (r *analisRepository) Create(analis model.Analis) (model.Analis, error) {
	var err = r.db.Create(&analis).Error

	return analis, err
}

func (r *analisRepository) Update(analis model.Analis) (model.Analis, error) {
	var err = r.db.Model(&analis).Updates(model.Analis{
		UserId:      analis.UserId,
		NamaLengkap: analis.NamaLengkap,
		Universitas: analis.Universitas,
	}).Error

	return analis, err
}

func (r *analisRepository) Delete(analis model.Analis) (model.Analis, error) {
	var err = r.db.Delete(&analis).Error

	return analis, err
}
