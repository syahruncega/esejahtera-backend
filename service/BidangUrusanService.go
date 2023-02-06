package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type BidangUrusanService interface {
	FindAll() ([]model.BidangUrusan, error)
	FindById(id int) (model.BidangUrusan, error)
	Create(bidangUrusanRequest request.CreateBidangUrusan) (model.BidangUrusan, error)
	Update(id int, bidangUrusanRequest request.UpdateBidangUrusan) (model.BidangUrusan, error)
	Delete(id int) (model.BidangUrusan, error)
}

type bidangUrusanService struct {
	bidangUrusanRepository repository.BidangUrusanRepository
}

func NewBidangUrusanService(bidangUrusanRepository repository.BidangUrusanRepository) *bidangUrusanService {
	return &bidangUrusanService{bidangUrusanRepository}
}

func (s *bidangUrusanService) FindAll() ([]model.BidangUrusan, error) {
	var bidangUrusans, err = s.bidangUrusanRepository.FindAll()

	return bidangUrusans, err
}

func (s *bidangUrusanService) FindById(id int) (model.BidangUrusan, error) {
	var bidangUrusan, err = s.bidangUrusanRepository.FindById(id)

	return bidangUrusan, err
}

func (s *bidangUrusanService) Create(bidangUrusanRequest request.CreateBidangUrusan) (model.BidangUrusan, error) {
	var bidangUrusan = model.BidangUrusan{
		BidangUrusanId:   bidangUrusanRequest.BidangUrusanId,
		NamaBidangUrusan: bidangUrusanRequest.NamaBidangUrusan,
	}

	var newBidangUrusan, err = s.bidangUrusanRepository.Create(bidangUrusan)

	return newBidangUrusan, err
}

func (s *bidangUrusanService) Update(id int, bidangUrusanRequest request.UpdateBidangUrusan) (model.BidangUrusan, error) {
	var bidangUrusan, err = s.bidangUrusanRepository.FindById(id)

	bidangUrusan.BidangUrusanId = bidangUrusanRequest.BidangUrusanId
	bidangUrusan.NamaBidangUrusan = bidangUrusanRequest.NamaBidangUrusan

	updatedBidangUrusan, err := s.bidangUrusanRepository.Update(bidangUrusan)

	return updatedBidangUrusan, err
}

func (s *bidangUrusanService) Delete(id int) (model.BidangUrusan, error) {
	var bidangUrusan, err = s.bidangUrusanRepository.FindById(id)

	deletedBidangUrusan, err := s.bidangUrusanRepository.Delete(bidangUrusan)

	return deletedBidangUrusan, err
}
