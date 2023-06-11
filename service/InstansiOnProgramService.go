package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type InstansiOnProgramService interface {
	FindAll() ([]model.InstansiOnProgram, error)
	FindById(id int) (model.InstansiOnProgram, error)
	FindByInstansiId(instansiId int) ([]model.InstansiOnProgram, error)
	FindBySearch(whereClause map[string]interface{}, tahun string) ([]model.InstansiOnProgram, error)
	CountJumlahProgramAllInstansi(tahun string, instansis []model.Instansi) []int64
	Create(instansiOnProgramRequest request.CreateInstansiOnProgramRequest) (model.InstansiOnProgram, error)
	Update(id int, instansiOnProgramRequest request.UpdateInstansiOnProgramRequest) (model.InstansiOnProgram, error)
	Delete(id int) (model.InstansiOnProgram, error)
}

type instansiOnProgramService struct {
	instansiOnProgramRepository repository.InstansiOnProgramRepository
}

func NewInstansiOnProgramService(instansiOnProgramRepository repository.InstansiOnProgramRepository) *instansiOnProgramService {
	return &instansiOnProgramService{instansiOnProgramRepository}
}

func (s *instansiOnProgramService) FindAll() ([]model.InstansiOnProgram, error) {
	var instansiOnPrograms, err = s.instansiOnProgramRepository.FindAll()

	return instansiOnPrograms, err
}

func (s *instansiOnProgramService) FindById(id int) (model.InstansiOnProgram, error) {
	var instansiOnProgram, err = s.instansiOnProgramRepository.FindById(id)

	return instansiOnProgram, err
}

func (s *instansiOnProgramService) FindByInstansiId(instansiId int) ([]model.InstansiOnProgram, error) {
	var instansiOnPrograms, err = s.instansiOnProgramRepository.FindByInstansiId(instansiId)

	return instansiOnPrograms, err
}

func (s *instansiOnProgramService) FindBySearch(whereClause map[string]interface{}, tahun string) ([]model.InstansiOnProgram, error) {
	var instansiOnPrograms, err = s.instansiOnProgramRepository.FindBySearch(whereClause, tahun)

	return instansiOnPrograms, err
}

func (s *instansiOnProgramService) CountJumlahProgramAllInstansi(tahun string, instansis []model.Instansi) []int64 {
	var hasil = s.instansiOnProgramRepository.CountJumlahProgramAllInstansi(tahun, instansis)

	return hasil
}

func (s *instansiOnProgramService) Create(instansiOnProgramRequest request.CreateInstansiOnProgramRequest) (model.InstansiOnProgram, error) {
	var instansiOnProgram = model.InstansiOnProgram{
		InstansiId: instansiOnProgramRequest.InstansiId,
		ProgramId:  instansiOnProgramRequest.ProgramId,
	}

	newInstansiOnProgram, err := s.instansiOnProgramRepository.Create(instansiOnProgram)

	return newInstansiOnProgram, err
}

func (s *instansiOnProgramService) Update(id int, instansiOnProgramRequest request.UpdateInstansiOnProgramRequest) (model.InstansiOnProgram, error) {
	var instansiOnProgram, err = s.instansiOnProgramRepository.FindById(id)

	instansiOnProgram.InstansiId = instansiOnProgramRequest.InstansiId
	instansiOnProgram.ProgramId = instansiOnProgramRequest.ProgramId

	updatedInstansiOnProgram, err := s.instansiOnProgramRepository.Update(instansiOnProgram)

	return updatedInstansiOnProgram, err
}

func (s *instansiOnProgramService) Delete(id int) (model.InstansiOnProgram, error) {
	var instansiOnProgram, err = s.instansiOnProgramRepository.FindById(id)

	deletedInstansiOnProgram, err := s.instansiOnProgramRepository.Delete(instansiOnProgram)

	return deletedInstansiOnProgram, err
}
