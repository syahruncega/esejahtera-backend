package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type InstansiService interface {
	FindAll() ([]model.Instansi, error)
	FindById(id int) (model.Instansi, error)
	FindAllRelation() ([]model.Instansi, error)
	Create(instansiRequest request.CreateInstansiRequest) (model.Instansi, error)
	Update(id int, instansiRequest request.UpdateInstansiRequest) (model.Instansi, error)
	Delete(id int) (model.Instansi, error)
}

type instansiService struct {
	instansiRepository repository.InstansiRepository
}

func NewInstansiService(instansiRepository repository.InstansiRepository) *instansiService {
	return &instansiService{instansiRepository}
}

func (s *instansiService) FindAll() ([]model.Instansi, error) {
	var instansis, err = s.instansiRepository.FindAll()

	return instansis, err
}

func (s *instansiService) FindById(id int) (model.Instansi, error) {
	var instansi, err = s.instansiRepository.FindById(id)

	return instansi, err
}

func (s *instansiService) FindAllRelation() ([]model.Instansi, error) {
	var instansiRelations, err = s.instansiRepository.FindAllRelation()

	return instansiRelations, err
}

func (s *instansiService) Create(instansiRequest request.CreateInstansiRequest) (model.Instansi, error) {
	var instansi = model.Instansi{
		NamaInstansi:   instansiRequest.NamaInstansi,
		BidangUrusanId: instansiRequest.BidangUrusanId,
	}

	var newInstansi, err = s.instansiRepository.Create(instansi)

	return newInstansi, err
}

func (s *instansiService) Update(id int, instansiRequest request.UpdateInstansiRequest) (model.Instansi, error) {
	var instansi, err = s.instansiRepository.FindById(id)

	instansi.NamaInstansi = instansiRequest.NamaInstansi
	instansi.BidangUrusanId = instansiRequest.BidangUrusanId

	updatedInstansi, err := s.instansiRepository.Update(instansi)

	return updatedInstansi, err
}

func (s *instansiService) Delete(id int) (model.Instansi, error) {
	var instansi, err = s.instansiRepository.FindById(id)

	deletedInstansi, err := s.instansiRepository.Delete(instansi)

	return deletedInstansi, err
}
