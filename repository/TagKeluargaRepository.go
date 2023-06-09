package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TagKeluargaRepository interface {
	FindAll() ([]model.TagKeluarga, error)
	FindById(id int) (model.TagKeluarga, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.TagKeluarga, error)
	Create(tagKeluarga model.TagKeluarga) (model.TagKeluarga, error)
	Update(tagKeluarga model.TagKeluarga) (model.TagKeluarga, error)
	Delete(tagKeluarga model.TagKeluarga) (model.TagKeluarga, error)
}

type tagKeluargaRepository struct {
	db *gorm.DB
}

func NewTagKeluargaRepository(db *gorm.DB) *tagKeluargaRepository {
	return &tagKeluargaRepository{db}
}

func (r *tagKeluargaRepository) FindAll() ([]model.TagKeluarga, error) {
	var tagKeluargas []model.TagKeluarga

	var err = r.db.Model(&tagKeluargas).Preload(clause.Associations).Preload("FokusBelanja").Preload("Keluarga." + clause.Associations).Find(&tagKeluargas).Error

	return tagKeluargas, err
}

func (r *tagKeluargaRepository) FindById(id int) (model.TagKeluarga, error) {
	var tagKeluarga model.TagKeluarga

	var err = r.db.Model(&tagKeluarga).Preload(clause.Associations).Preload("FokusBelanja").Preload("Keluarga."+clause.Associations).Take(&tagKeluarga, id).Error

	return tagKeluarga, err
}

func (r *tagKeluargaRepository) FindBySearch(whereClause map[string]interface{}) ([]model.TagKeluarga, error) {
	var tagKeluargas []model.TagKeluarga

	var err = r.db.Where(whereClause).Model(&tagKeluargas).Preload(clause.Associations).Preload("FokusBelanja").Preload("Keluarga." + clause.Associations).Find(&tagKeluargas).Error

	return tagKeluargas, err
}

func (r *tagKeluargaRepository) Create(tagKeluarga model.TagKeluarga) (model.TagKeluarga, error) {
	var err = r.db.Create(&tagKeluarga).Error

	return tagKeluarga, err
}

func (r *tagKeluargaRepository) Update(tagKeluarga model.TagKeluarga) (model.TagKeluarga, error) {
	var err = r.db.Model(&tagKeluarga).Updates(model.TagKeluarga{
		FokusBelanjaId: tagKeluarga.FokusBelanjaId,
		KeluargaId:     tagKeluarga.KeluargaId,
	}).Error

	return tagKeluarga, err
}

func (r *tagKeluargaRepository) Delete(tagKeluarga model.TagKeluarga) (model.TagKeluarga, error) {
	var err = r.db.Delete(&tagKeluarga).Error

	return tagKeluarga, err
}
