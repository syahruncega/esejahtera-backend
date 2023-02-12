package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type PusbangRepository interface {
	FindAll() ([]model.Pusbang, error)
	FindById(id int) (model.Pusbang, error)
	FindByUserId(userId int) (model.Pusbang, error)
	FindAllRelation() ([]model.Pusbang, error)
	Create(pusbang model.Pusbang) (model.Pusbang, error)
	Update(pusbang model.Pusbang) (model.Pusbang, error)
	Delete(pusbang model.Pusbang) (model.Pusbang, error)
}

type pusbangRepository struct {
	db *gorm.DB
}

func NewPusbangRepository(db *gorm.DB) *pusbangRepository {
	return &pusbangRepository{db}
}

func (r *pusbangRepository) FindAll() ([]model.Pusbang, error) {
	var pusbangs []model.Pusbang

	var err = r.db.Find(&pusbangs).Error

	return pusbangs, err
}

func (r *pusbangRepository) FindById(id int) (model.Pusbang, error) {
	var pusbang model.Pusbang

	var err = r.db.Take(&pusbang, id).Error

	return pusbang, err
}

func (r *pusbangRepository) FindByUserId(userId int) (model.Pusbang, error) {
	var pusbang model.Pusbang

	var err = r.db.Where("userId = ?", userId).Take(&pusbang).Error

	return pusbang, err
}

func (r *pusbangRepository) FindAllRelation() ([]model.Pusbang, error) {
	var pusbangs []model.Pusbang

	var err = r.db.Model(&pusbangs).Preload("User").Find(&pusbangs).Error

	return pusbangs, err
}

func (r *pusbangRepository) Create(pusbang model.Pusbang) (model.Pusbang, error) {
	var err = r.db.Create(&pusbang).Error

	return pusbang, err
}

func (r *pusbangRepository) Update(pusbang model.Pusbang) (model.Pusbang, error) {
	var err = r.db.Model(&pusbang).Updates(model.Pusbang{
		UserId:      pusbang.UserId,
		NamaLengkap: pusbang.NamaLengkap,
		Universitas: pusbang.Universitas,
	}).Error

	return pusbang, err
}

func (r *pusbangRepository) Delete(pusbang model.Pusbang) (model.Pusbang, error) {
	var err = r.db.Delete(&pusbang).Error

	return pusbang, err
}
