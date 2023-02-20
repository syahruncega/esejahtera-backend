package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type AnalisService interface {
	FindAll() ([]model.Analis, error)
	FindById(id int) (model.Analis, error)
	FindByUserId(userId int) (model.Analis, error)
	FindAllRelation() ([]model.Analis, error)
	Create(analisRequest request.CreateAnalisRequest) (model.Analis, error)
	Update(id int, analisRequst request.UpdateAnalisRequest) (model.Analis, error)
	Delete(id int) (model.Analis, error)
}

type analisService struct {
	analisRepository repository.AnalisRepository
}

func NewAnalisService(analisRepository repository.AnalisRepository) *analisService {
	return &analisService{analisRepository}
}

func (s *analisService) FindAll() ([]model.Analis, error) {
	var analiss, err = s.analisRepository.FindAll()

	return analiss, err
}

func (s *analisService) FindById(id int) (model.Analis, error) {
	var analis, err = s.analisRepository.FindById(id)

	return analis, err
}

func (s *analisService) FindByUserId(userId int) (model.Analis, error) {
	var analis, err = s.analisRepository.FindByUserId(userId)

	return analis, err
}

func (s *analisService) FindAllRelation() ([]model.Analis, error) {
	var analiss, err = s.analisRepository.FindAllRelation()

	return analiss, err
}

func (s *analisService) Create(analisRequest request.CreateAnalisRequest) (model.Analis, error) {
	var analis = model.Analis{
		UserId:        analisRequest.UserId,
		NamaLengkap:   analisRequest.NamaLengkap,
		Universitas:   analisRequest.Universitas,
		UrlFotoProfil: analisRequest.UrlFotoProfil,
	}

	newAnalis, err := s.analisRepository.Create(analis)

	return newAnalis, err
}

func (s *analisService) Update(id int, analisRequest request.UpdateAnalisRequest) (model.Analis, error) {
	var analis, err = s.analisRepository.FindById(id)

	analis.UserId = analisRequest.UserId
	analis.NamaLengkap = analisRequest.NamaLengkap
	analis.Universitas = analisRequest.Universitas
	analis.UrlFotoProfil = analisRequest.UrlFotoProfil

	updatedAnalis, err := s.analisRepository.Update(analis)

	return updatedAnalis, err
}

func (s *analisService) Delete(id int) (model.Analis, error) {
	var analis, err = s.analisRepository.FindById(id)

	deletedInstansi, err := s.analisRepository.Delete(analis)

	return deletedInstansi, err
}
