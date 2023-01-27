package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type IndividuService interface {
	FindAll(kabupatenKotaId string) ([]model.Individu, error)
	FindById(kabupatenKotaid string, id int) (model.Individu, error)
	FindByIdKeluarga(kabupatenKotaId string, idKeluarga string) ([]model.Individu, error)
	Update(kabupatenKotaId string, id int, individuReqeust request.UpdateIndividuRequest) (model.Individu, error)
}

type individuService struct {
	individuRepository repository.IndividuRepository
}

func NewIndividuService(individuRepository repository.IndividuRepository) *individuService {
	return &individuService{individuRepository}
}

func (s *individuService) FindAll(kabupatenKotaId string) ([]model.Individu, error) {
	var individus, err = s.individuRepository.FindAll(kabupatenKotaId)

	return individus, err
}

func (s *individuService) FindById(kabupatenKotaId string, id int) (model.Individu, error) {
	var individu, err = s.individuRepository.FindById(kabupatenKotaId, id)

	return individu, err
}

func (s *individuService) FindByIdKeluarga(kabupatenKotaId string, idKeluarga string) ([]model.Individu, error) {
	var individus, err = s.individuRepository.FindByIdKeluarga(kabupatenKotaId, idKeluarga)

	return individus, err
}

func (s *individuService) Update(kabupatenKotaId string, id int, individuRequest request.UpdateIndividuRequest) (model.Individu, error) {
	var individu, err = s.individuRepository.FindById(kabupatenKotaId, id)

	individu.UserId = individuRequest.UserId
	individu.MahasiswaId = individuRequest.MahasiswaId
	individu.StatusVerifikasi = individuRequest.StatusVerifikasi

	updatedIndividu, err := s.individuRepository.Update(individu)

	return updatedIndividu, err
}
