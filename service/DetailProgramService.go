package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type DetailProgramService interface {
	FindAll() ([]model.DetailProgram, error)
	FindById(id int) (model.DetailProgram, error)
	FindByInstansi(instansiId int) ([]model.DetailProgram, error)
	Create(detailProgramRequest request.CreateDetailProgramRequest) (model.DetailProgram, error)
	Update(id int, detailProgramRequest request.UpdateDetailProgramRequest) (model.DetailProgram, error)
	Delete(id int) (model.DetailProgram, error)
}

type detailProgramService struct {
	detailProgramRepository repository.DetailProgramRepository
}

func NewDetailProgramService(detailProgramRepository repository.DetailProgramRepository) *detailProgramService {
	return &detailProgramService{detailProgramRepository}
}

func (s *detailProgramService) FindAll() ([]model.DetailProgram, error) {
	var detailPrograms, err = s.detailProgramRepository.FindAll()

	return detailPrograms, err
}

func (s *detailProgramService) FindById(id int) (model.DetailProgram, error) {
	var detailProgram, err = s.detailProgramRepository.FindById(id)

	return detailProgram, err
}

func (s *detailProgramService) FindByInstansi(instansiId int) ([]model.DetailProgram, error) {
	var detailPrograms, err = s.detailProgramRepository.FindByInstansi(instansiId)

	return detailPrograms, err
}

func (s *detailProgramService) Create(detailProgramRequest request.CreateDetailProgramRequest) (model.DetailProgram, error) {
	var detailProgram = model.DetailProgram{
		InstansiId:  detailProgramRequest.InstansiId,
		ProgramId:   detailProgramRequest.ProgramId,
		PaguProgram: detailProgramRequest.PaguProgram,
	}

	newDetailProgram, err := s.detailProgramRepository.Create(detailProgram)

	return newDetailProgram, err
}

func (s *detailProgramService) Update(id int, detailProgramRequest request.UpdateDetailProgramRequest) (model.DetailProgram, error) {
	var detailProgram, err = s.detailProgramRepository.FindById(id)

	detailProgram.InstansiId = detailProgramRequest.InstansiId
	detailProgram.ProgramId = detailProgramRequest.ProgramId
	detailProgram.PaguProgram = detailProgramRequest.PaguProgram

	updatedDetailProgram, err := s.detailProgramRepository.Update(detailProgram)

	return updatedDetailProgram, err
}

func (s *detailProgramService) Delete(id int) (model.DetailProgram, error) {
	var detailProgram, err = s.detailProgramRepository.FindById(id)

	deletedDetailProgram, err := s.detailProgramRepository.Delete(detailProgram)

	return deletedDetailProgram, err
}
