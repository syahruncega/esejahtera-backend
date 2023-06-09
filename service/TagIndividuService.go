package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type TagIndividuService interface {
	FindAll() ([]model.TagIndividu, error)
	FindById(id int) (model.TagIndividu, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.TagIndividu, error)
	Create(tagIndividuRequest request.CreateTagIndividuRequest) (model.TagIndividu, error)
	Update(id int, tagIndividuRequest request.UpdateTagIndividuRequest) (model.TagIndividu, error)
	Delete(id int) (model.TagIndividu, error)
}

type tagIndividuService struct {
	tagIndividuRepository repository.TagIndividuRepository
}

func NewTagIndividuService(tagIndividuRepository repository.TagIndividuRepository) *tagIndividuService {
	return &tagIndividuService{tagIndividuRepository}
}

func (s *tagIndividuService) FindAll() ([]model.TagIndividu, error) {
	var tagIndividus, err = s.tagIndividuRepository.FindAll()

	return tagIndividus, err
}

func (s *tagIndividuService) FindById(id int) (model.TagIndividu, error) {
	var tagIndividu, err = s.tagIndividuRepository.FindById(id)

	return tagIndividu, err
}

func (s *tagIndividuService) FindBySearch(whereClause map[string]interface{}) ([]model.TagIndividu, error) {
	var tagIndividus, err = s.tagIndividuRepository.FindBySearch(whereClause)

	return tagIndividus, err
}

func (s *tagIndividuService) Create(tagIndividuRequest request.CreateTagIndividuRequest) (model.TagIndividu, error) {
	var tagIndividu = model.TagIndividu{
		FokusBelanjaId: tagIndividuRequest.FokusBelanjaId,
		IndividuId:     tagIndividuRequest.IndividuId,
	}

	var newTagIndividu, err = s.tagIndividuRepository.Create(tagIndividu)

	return newTagIndividu, err
}

func (s *tagIndividuService) Update(id int, tagIndividuRequest request.UpdateTagIndividuRequest) (model.TagIndividu, error) {
	var tagIndividu, err = s.tagIndividuRepository.FindById(id)

	tagIndividu.FokusBelanjaId = tagIndividuRequest.FokusBelanjaId
	tagIndividu.IndividuId = tagIndividuRequest.IndividuId

	updatedTagIndividu, err := s.tagIndividuRepository.Update(tagIndividu)

	return updatedTagIndividu, err
}

func (s *tagIndividuService) Delete(id int) (model.TagIndividu, error) {
	var tagIndividu, err = s.tagIndividuRepository.FindById(id)

	deletedTagIndividu, err := s.tagIndividuRepository.Delete(tagIndividu)

	return deletedTagIndividu, err
}
