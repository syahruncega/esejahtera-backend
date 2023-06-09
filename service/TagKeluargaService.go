package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type TagKeluargaService interface {
	FindAll() ([]model.TagKeluarga, error)
	FindById(id int) (model.TagKeluarga, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.TagKeluarga, error)
	Create(tagKeluargaRequest request.CreateTagKeluargaRequest) (model.TagKeluarga, error)
	Update(id int, tagKeluargaRequest request.UpdateTagKeluargaRequest) (model.TagKeluarga, error)
	Delete(id int) (model.TagKeluarga, error)
}

type tagKeluargaService struct {
	tagKeluargaRepository repository.TagKeluargaRepository
}

func NewTagKeluargaService(tagKeluargaRepository repository.TagKeluargaRepository) *tagKeluargaService {
	return &tagKeluargaService{tagKeluargaRepository}
}

func (s *tagKeluargaService) FindAll() ([]model.TagKeluarga, error) {
	var tagKeluargas, err = s.tagKeluargaRepository.FindAll()

	return tagKeluargas, err
}

func (s *tagKeluargaService) FindById(id int) (model.TagKeluarga, error) {
	var tagKeluarga, err = s.tagKeluargaRepository.FindById(id)

	return tagKeluarga, err
}

func (s *tagKeluargaService) FindBySearch(whereClause map[string]interface{}) ([]model.TagKeluarga, error) {
	var tagKeluargas, err = s.tagKeluargaRepository.FindBySearch(whereClause)

	return tagKeluargas, err
}

func (s *tagKeluargaService) Create(tagKeluargaRequest request.CreateTagKeluargaRequest) (model.TagKeluarga, error) {
	var tagKeluarga = model.TagKeluarga{
		FokusBelanjaId: tagKeluargaRequest.FokusBelanjaId,
		KeluargaId:     tagKeluargaRequest.KeluargaId,
	}

	var newTagKeluarga, err = s.tagKeluargaRepository.Create(tagKeluarga)

	return newTagKeluarga, err
}

func (s *tagKeluargaService) Update(id int, tagKeluargaRequest request.UpdateTagKeluargaRequest) (model.TagKeluarga, error) {
	var tagKeluarga, err = s.tagKeluargaRepository.FindById(id)

	tagKeluarga.FokusBelanjaId = tagKeluargaRequest.FokusBelanjaId
	tagKeluarga.KeluargaId = tagKeluargaRequest.KeluargaId

	updatedTagKeluarga, err := s.tagKeluargaRepository.Update(tagKeluarga)

	return updatedTagKeluarga, err
}

func (s *tagKeluargaService) Delete(id int) (model.TagKeluarga, error) {
	var tagKeluarga, err = s.tagKeluargaRepository.FindById(id)

	deletedTagKeluarga, err := s.tagKeluargaRepository.Delete(tagKeluarga)

	return deletedTagKeluarga, err
}
