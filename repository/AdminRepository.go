package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type AdminRepository interface {
	FindAll() ([]model.Admin, error)
	FindById(id int) (model.Admin, error)
	FindAllRelation() ([]model.Admin, error)
	Create(admin model.Admin) (model.Admin, error)
	Update(admin model.Admin) (model.Admin, error)
	Delete(admin model.Admin) (model.Admin, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db}
}

func (r *adminRepository) FindAll() ([]model.Admin, error) {
	var admins []model.Admin

	var err = r.db.Find(&admins).Error

	return admins, err
}

func (r *adminRepository) FindById(id int) (model.Admin, error) {
	var admin model.Admin

	var err = r.db.Take(&admin, id).Error

	return admin, err
}

func (r *adminRepository) FindAllRelation() ([]model.Admin, error) {
	var admins []model.Admin

	var err = r.db.Model(&admins).Preload("User").Find(&admins).Error

	return admins, err
}

func (r *adminRepository) Create(admin model.Admin) (model.Admin, error) {
	var err = r.db.Create(&admin).Error

	return admin, err
}

func (r *adminRepository) Update(admin model.Admin) (model.Admin, error) {
	var err = r.db.Model(&admin).Updates(model.Admin{
		UserId:      admin.UserId,
		NamaLengkap: admin.NamaLengkap,
	}).Error

	return admin, err
}

func (r *adminRepository) Delete(admin model.Admin) (model.Admin, error) {
	var err = r.db.Delete(&admin).Error

	return admin, err
}
