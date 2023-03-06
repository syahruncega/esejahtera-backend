package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type BidangUrusanOnInstansiService interface {
	FindAll() ([]model.BidangUrusanOnInstansi, error)
	FindById(id int) (model.BidangUrusanOnInstansi, error)
	FindByInstansiId(instansiId int) ([]model.BidangUrusanOnInstansi, error)
	Create(bidangUrusanOnInstansiRequest request.CreateBidangUrusanOnInstansiRequest) (model.BidangUrusanOnInstansi, error)
	Update(id int, bidangUrusanOnInstansiRequest request.UpdateBidangUrusanOnInstansiRequest) (model.BidangUrusanOnInstansi, error)
	Delete(id int) (model.BidangUrusanOnInstansi, error)
}

type bidangUrusanOnInstansiService struct {
	bidangUrusanOnInstansiRepository repository.BidangUrusanOnInstansiRepository
}

func NewBidangUrusanOnInstansiService(bidangUrusanOnInstansiRepository repository.BidangUrusanOnInstansiRepository) *bidangUrusanOnInstansiService {
	return &bidangUrusanOnInstansiService{bidangUrusanOnInstansiRepository}
}

func (s *bidangUrusanOnInstansiService) FindAll() ([]model.BidangUrusanOnInstansi, error) {
	var bidangUrusanOnInstansis, err = s.bidangUrusanOnInstansiRepository.FindAll()

	return bidangUrusanOnInstansis, err
}

func (s *bidangUrusanOnInstansiService) FindById(id int) (model.BidangUrusanOnInstansi, error) {
	var bidangUrusanOnInstansi, err = s.bidangUrusanOnInstansiRepository.FindById(id)

	return bidangUrusanOnInstansi, err
}

func (s *bidangUrusanOnInstansiService) FindByInstansiId(instansiId int) ([]model.BidangUrusanOnInstansi, error) {
	var bidangUrusanOnInstansis, err = s.bidangUrusanOnInstansiRepository.FindByInstansiId(instansiId)

	return bidangUrusanOnInstansis, err
}

func (s *bidangUrusanOnInstansiService) Create(bidangUrusanOnInstansiRequest request.CreateBidangUrusanOnInstansiRequest) (model.BidangUrusanOnInstansi, error) {
	var bidangUrusanOnInstansi = model.BidangUrusanOnInstansi{
		InstansiId:     bidangUrusanOnInstansiRequest.InstansiId,
		BidangUrusanId: bidangUrusanOnInstansiRequest.BidangUrusanId,
	}

	newBidangUrusanOnInstansi, err := s.bidangUrusanOnInstansiRepository.Create(bidangUrusanOnInstansi)

	return newBidangUrusanOnInstansi, err
}

func (s *bidangUrusanOnInstansiService) Update(id int, bidangUrusanOnInstansiRequest request.UpdateBidangUrusanOnInstansiRequest) (model.BidangUrusanOnInstansi, error) {
	var bidangUrusanOnInstansi, err = s.bidangUrusanOnInstansiRepository.FindById(id)

	bidangUrusanOnInstansi.InstansiId = bidangUrusanOnInstansiRequest.InstansiId
	bidangUrusanOnInstansi.BidangUrusanId = bidangUrusanOnInstansiRequest.BidangUrusanId

	updatedBidangUrusanOnInstansi, err := s.bidangUrusanOnInstansiRepository.Update(bidangUrusanOnInstansi)

	return updatedBidangUrusanOnInstansi, err
}

func (s *bidangUrusanOnInstansiService) Delete(id int) (model.BidangUrusanOnInstansi, error) {
	var bidangUrusanOnInstansi, err = s.bidangUrusanOnInstansiRepository.FindById(id)

	deletedBidangUrusanOnInstansi, err := s.bidangUrusanOnInstansiRepository.Delete(bidangUrusanOnInstansi)

	return deletedBidangUrusanOnInstansi, err
}
