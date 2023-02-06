package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type ProgramService interface {
	FindAll() ([]model.Program, error)
	FindById(id int) (model.Program, error)
	Create(programRequest request.CreateProgramRequest) (model.Program, error)
	Update(id int, programRequest request.UpdateProgramRequest) (model.Program, error)
	Delete(id int) (model.Program, error)
}

type programService struct {
	programRepository repository.ProgramRepository
}

func NewProgramService(programRepository repository.ProgramRepository) *programService {
	return &programService{programRepository}
}

func (s *programService) FindAll() ([]model.Program, error) {
	var programs, err = s.programRepository.FindAll()

	return programs, err
}

func (s *programService) FindById(id int) (model.Program, error) {
	var program, err = s.programRepository.FindById(id)

	return program, err
}

func (s *programService) Create(programRequest request.CreateProgramRequest) (model.Program, error) {
	var program = model.Program{
		ProgramId:   programRequest.ProgramId,
		NamaProgram: programRequest.NamaProgram,
	}

	newProgram, err := s.programRepository.Create(program)

	return newProgram, err
}

func (s *programService) Update(id int, programRequest request.UpdateProgramRequest) (model.Program, error) {
	var program, err = s.programRepository.FindById(id)

	program.ProgramId = programRequest.ProgramId
	program.NamaProgram = programRequest.NamaProgram

	updatedProgram, err := s.programRepository.Update(program)

	return updatedProgram, err
}

func (s *programService) Delete(id int) (model.Program, error) {
	var program, err = s.programRepository.FindById(id)

	deletedProgram, err := s.programRepository.Delete(program)

	return deletedProgram, err
}
