package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type PusbangService interface {
	FindAll() ([]model.Pusbang, error)
	FindById(id int) (model.Pusbang, error)
	FindByUserId(userId int) (model.Pusbang, error)
	FindAllRelation() ([]model.Pusbang, error)
	Create(pusbangRequest request.CreatePusbangRequest) (model.Pusbang, error)
	Update(id int, pusbangRequest request.UpdatePusbangRequest) (model.Pusbang, error)
	Delete(id int) (model.Pusbang, error)
}

type pusbangService struct {
	pusbangRepository repository.PusbangRepository
}

func NewPusbangService(pusbangRepository repository.PusbangRepository) *pusbangService {
	return &pusbangService{pusbangRepository}
}

func (s *pusbangService) FindAll() ([]model.Pusbang, error) {
	var pusbangs, err = s.pusbangRepository.FindAll()

	return pusbangs, err
}

func (s *pusbangService) FindById(id int) (model.Pusbang, error) {
	var pusbang, err = s.pusbangRepository.FindById(id)

	return pusbang, err
}

func (s *pusbangService) FindByUserId(userId int) (model.Pusbang, error) {
	var pusbang, err = s.pusbangRepository.FindByUserId(userId)

	return pusbang, err
}

func (s *pusbangService) FindAllRelation() ([]model.Pusbang, error) {
	var pusbangs, err = s.pusbangRepository.FindAllRelation()

	return pusbangs, err
}

func (s *pusbangService) Create(pusbangRequest request.CreatePusbangRequest) (model.Pusbang, error) {
	var pusbang = model.Pusbang{
		UserId:      pusbangRequest.UserId,
		NamaLengkap: pusbangRequest.NamaLengkap,
		Universitas: pusbangRequest.Universitas,
	}

	newDosen, err := s.pusbangRepository.Create(pusbang)

	return newDosen, err
}

func (s *pusbangService) Update(id int, pusbangRequest request.UpdatePusbangRequest) (model.Pusbang, error) {
	var pusbang, err = s.pusbangRepository.FindById(id)

	pusbang.UserId = pusbangRequest.UserId
	pusbang.NamaLengkap = pusbangRequest.NamaLengkap
	pusbang.Universitas = pusbangRequest.Universitas

	updatedPusbang, err := s.pusbangRepository.Update(pusbang)

	return updatedPusbang, err
}

func (s *pusbangService) Delete(id int) (model.Pusbang, error) {
	var pusbang, err = s.pusbangRepository.FindById(id)

	deletedPusbang, err := s.pusbangRepository.Delete(pusbang)

	return deletedPusbang, err
}
