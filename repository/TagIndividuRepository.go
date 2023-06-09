package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TagIndividuRepository interface {
	FindAll() ([]model.TagIndividu, error)
	FindById(id int) (model.TagIndividu, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.TagIndividu, error)
	Create(tagIndividu model.TagIndividu) (model.TagIndividu, error)
	Update(tagIndividu model.TagIndividu) (model.TagIndividu, error)
	Delete(tagIndividu model.TagIndividu) (model.TagIndividu, error)
}

type tagIndividuRepository struct {
	db *gorm.DB
}

func NewTagIndividuRepository(db *gorm.DB) *tagIndividuRepository {
	return &tagIndividuRepository{db}
}

func (r *tagIndividuRepository) FindAll() ([]model.TagIndividu, error) {
	var tagIndividus []model.TagIndividu

	var err = r.db.Model(&tagIndividus).Preload(clause.Associations).Preload("FokusBelanja").Preload("Individu." + clause.Associations).Find(&tagIndividus).Error

	return tagIndividus, err
}

func (r *tagIndividuRepository) FindById(id int) (model.TagIndividu, error) {
	var tagIndividu model.TagIndividu

	var err = r.db.Model(&tagIndividu).Preload(clause.Associations).Preload("FokusBelanja").Preload("Individu."+clause.Associations).Preload(clause.Associations).Take(&tagIndividu, id).Error

	return tagIndividu, err
}

func (r *tagIndividuRepository) FindBySearch(whereClause map[string]interface{}) ([]model.TagIndividu, error) {
	var tagIndividus []model.TagIndividu

	var err = r.db.Where(whereClause).Model(&tagIndividus).Preload(clause.Associations).Preload("FokusBelanja").Preload("Individu." + clause.Associations).Find(&tagIndividus).Error

	return tagIndividus, err
}

func (r *tagIndividuRepository) Create(tagIndividu model.TagIndividu) (model.TagIndividu, error) {
	var err = r.db.Create(&tagIndividu).Error

	return tagIndividu, err
}

func (r *tagIndividuRepository) Update(tagIndividu model.TagIndividu) (model.TagIndividu, error) {
	var err = r.db.Model(&tagIndividu).Updates(model.TagIndividu{
		FokusBelanjaId: tagIndividu.FokusBelanjaId,
		IndividuId:     tagIndividu.IndividuId,
	}).Error

	return tagIndividu, err
}

func (r *tagIndividuRepository) Delete(tagIndividu model.TagIndividu) (model.TagIndividu, error) {
	var err = r.db.Delete(&tagIndividu).Error

	return tagIndividu, err
}
