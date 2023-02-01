package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type IndividuService interface {
	FindAll() ([]model.Individu, error)
	FindById(id int) (model.Individu, error)
	FindByIdKeluarga(idKeluarga string) ([]model.Individu, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.Individu, error)
	Update(id int, individuReqeust request.UpdateIndividuRequest) (model.Individu, error)
}

type individuService struct {
	individuRepository repository.IndividuRepository
}

func NewIndividuService(individuRepository repository.IndividuRepository) *individuService {
	return &individuService{individuRepository}
}

func (s *individuService) FindAll() ([]model.Individu, error) {
	var individus, err = s.individuRepository.FindAll()

	return individus, err
}

func (s *individuService) FindById(id int) (model.Individu, error) {
	var individu, err = s.individuRepository.FindById(id)

	return individu, err
}

func (s *individuService) FindByIdKeluarga(idKeluarga string) ([]model.Individu, error) {
	var individus, err = s.individuRepository.FindByIdKeluarga(idKeluarga)

	return individus, err
}

func (s *individuService) FindBySearch(whereClause map[string]interface{}) ([]model.Individu, error) {
	var individus, err = s.individuRepository.FindBySearch(whereClause)

	return individus, err
}

func (s *individuService) Update(id int, individuRequest request.UpdateIndividuRequest) (model.Individu, error) {
	var individu, err = s.individuRepository.FindById(id)

	individu.UserId = individuRequest.UserId
	individu.MahasiswaId = individuRequest.MahasiswaId
	individu.StatusVerifikasi = individuRequest.StatusVerifikasi

	updatedIndividu, err := s.individuRepository.Update(individu)

	return updatedIndividu, err
}
