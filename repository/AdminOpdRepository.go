package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type AdminOpdRepository interface {
	FindAll() ([]model.AdminOpd, error)
	FindById(id int) (model.AdminOpd, error)
	FindByUserId(userId int) (model.AdminOpd, error)
	Create(adminOpd model.AdminOpd) (model.AdminOpd, error)
	Update(adminOpd model.AdminOpd) (model.AdminOpd, error)
	Delete(adminOpd model.AdminOpd) (model.AdminOpd, error)
}

type adminOpdRepository struct {
	db *gorm.DB
}

func NewAdminOpdRepository(db *gorm.DB) *adminOpdRepository {
	return &adminOpdRepository{db}
}

func (r *adminOpdRepository) FindAll() ([]model.AdminOpd, error) {
	var adminOpds []model.AdminOpd

	var err = r.db.Model(&adminOpds).Preload("User").Preload("Instansi").Find(&adminOpds).Error

	return adminOpds, err
}

func (r *adminOpdRepository) FindById(id int) (model.AdminOpd, error) {
	var adminOpd model.AdminOpd

	var err = r.db.Model(&adminOpd).Preload("User").Preload("Instansi").Take(&adminOpd, id).Error

	return adminOpd, err
}

func (r *adminOpdRepository) FindByUserId(userId int) (model.AdminOpd, error) {
	var adminOpd model.AdminOpd

	var err = r.db.Where("userId = ?", userId).Model(&adminOpd).Preload("User").Preload("Instansi").Find(&adminOpd).Error

	return adminOpd, err
}

func (r *adminOpdRepository) Create(adminOpd model.AdminOpd) (model.AdminOpd, error) {
	var err = r.db.Create(&adminOpd).Error

	return adminOpd, err
}

func (r *adminOpdRepository) Update(adminOpd model.AdminOpd) (model.AdminOpd, error) {
	var err = r.db.Model(&adminOpd).Updates(model.AdminOpd{
		UserId:        adminOpd.UserId,
		NamaLengkap:   adminOpd.NamaLengkap,
		InstansiId:    adminOpd.InstansiId,
		UrlFotoProfil: adminOpd.UrlFotoProfil,
	}).Error

	return adminOpd, err
}

func (r *adminOpdRepository) Delete(adminOpd model.AdminOpd) (model.AdminOpd, error) {
	var err = r.db.Delete(&adminOpd).Error

	return adminOpd, err
}
