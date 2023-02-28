package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type DetailInstansiService interface {
	FindAll() ([]model.DetailInstansi, error)
	FindById(id int) (model.DetailInstansi, error)
	FindByBidangUrusan(bidangUrusanId int) ([]model.DetailInstansi, error)
	FindByInstansiId(instansiId int) ([]model.DetailInstansi, error)
	Create(detailInstansiRequest request.CreateDetailInstansiRequest) (model.DetailInstansi, error)
	Update(id int, detailInstansiRequest request.UpdateDetailInstansiRequest) (model.DetailInstansi, error)
	Delete(id int) (model.DetailInstansi, error)
}

type detailInstansiService struct {
	detailInstansiRepository repository.DetailInstansiRepository
}

func NewDetailInstansiService(detailInstansiRepository repository.DetailInstansiRepository) *detailInstansiService {
	return &detailInstansiService{detailInstansiRepository}
}

func (s *detailInstansiService) FindAll() ([]model.DetailInstansi, error) {
	var detailInstansis, err = s.detailInstansiRepository.FindAll()

	return detailInstansis, err
}

func (s *detailInstansiService) FindById(id int) (model.DetailInstansi, error) {
	var detailInstansi, err = s.detailInstansiRepository.FindById(id)

	return detailInstansi, err
}

func (s *detailInstansiService) FindByBidangUrusan(bidangUrusanId int) ([]model.DetailInstansi, error) {
	var detailInstansis, err = s.detailInstansiRepository.FindByBidangUrusan(bidangUrusanId)

	return detailInstansis, err
}

func (s *detailInstansiService) FindByInstansiId(instansiId int) ([]model.DetailInstansi, error) {
	var detailInstansis, err = s.detailInstansiRepository.FindByInstansiId(instansiId)

	return detailInstansis, err
}

func (s *detailInstansiService) Create(detailInstansiRequest request.CreateDetailInstansiRequest) (model.DetailInstansi, error) {
	var detailInstansi = model.DetailInstansi{
		InstansiId:     detailInstansiRequest.InstansiId,
		BidangUrusanId: detailInstansiRequest.BidangUrusanId,
	}

	newDetailInstansi, err := s.detailInstansiRepository.Create(detailInstansi)

	return newDetailInstansi, err
}

func (s *detailInstansiService) Update(id int, detailInstansiRequest request.UpdateDetailInstansiRequest) (model.DetailInstansi, error) {
	var detailInstansi, err = s.detailInstansiRepository.FindById(id)

	detailInstansi.InstansiId = detailInstansiRequest.InstansiId
	detailInstansi.BidangUrusanId = detailInstansiRequest.BidangUrusanId

	updatedDetailInstansi, err := s.detailInstansiRepository.Update(detailInstansi)

	return updatedDetailInstansi, err
}

func (s *detailInstansiService) Delete(id int) (model.DetailInstansi, error) {
	var detailInstansi, err = s.detailInstansiRepository.FindById(id)

	deletedDetailInstansi, err := s.detailInstansiRepository.Delete(detailInstansi)

	return deletedDetailInstansi, err
}
