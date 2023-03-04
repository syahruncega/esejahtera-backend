package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type RencanaProgramService interface {
	FindAll() ([]model.RencanaProgram, error)
	FindById(id int) (model.RencanaProgram, error)
	Create(rencanaProgramRequest request.CreateRencanaProgramRequest) (model.RencanaProgram, error)
	Update(id int, rencanaProgramRequest request.UpdateRencanaProgramRequest) (model.RencanaProgram, error)
	Delete(id int) (model.RencanaProgram, error)
}

type rencanaProgramService struct {
	rencanaProgramRepository repository.RencanaProgramRepository
}

func NewRencanaProgramService(rencanaProgramRepository repository.RencanaProgramRepository) *rencanaProgramService {
	return &rencanaProgramService{rencanaProgramRepository}
}

func (s *rencanaProgramService) FindAll() ([]model.RencanaProgram, error) {
	var rencanaPrograms, err = s.rencanaProgramRepository.FindAll()

	return rencanaPrograms, err
}

func (s *rencanaProgramService) FindById(id int) (model.RencanaProgram, error) {
	var rencanaProgram, err = s.rencanaProgramRepository.FindById(id)

	return rencanaProgram, err
}

func (s *rencanaProgramService) Create(rencanaProgramRequest request.CreateRencanaProgramRequest) (model.RencanaProgram, error) {
	var rencanaProgram = model.RencanaProgram{
		InstansiId:  rencanaProgramRequest.InstansiId,
		ProgramId:   rencanaProgramRequest.ProgramId,
		PaguProgram: rencanaProgramRequest.PaguProgram,
		Tipe:        rencanaProgramRequest.Tipe,
	}

	newRencanaProgram, err := s.rencanaProgramRepository.Create(rencanaProgram)

	return newRencanaProgram, err
}

func (s *rencanaProgramService) Update(id int, rencanaProgramRequest request.UpdateRencanaProgramRequest) (model.RencanaProgram, error) {
	var rencanaProgram, err = s.rencanaProgramRepository.FindById(id)

	rencanaProgram.InstansiId = rencanaProgramRequest.InstansiId
	rencanaProgram.ProgramId = rencanaProgramRequest.ProgramId
	rencanaProgram.PaguProgram = rencanaProgramRequest.PaguProgram
	rencanaProgram.Tipe = rencanaProgramRequest.Tipe

	updatedRencanaProgram, err := s.rencanaProgramRepository.Update(rencanaProgram)

	return updatedRencanaProgram, err
}

func (s *rencanaProgramService) Delete(id int) (model.RencanaProgram, error) {
	var rencanaProgram, err = s.rencanaProgramRepository.FindById(id)

	deletedRencanaProgram, err := s.rencanaProgramRepository.Delete(rencanaProgram)

	return deletedRencanaProgram, err
}
