package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type ProgramService interface {
	FindAll() ([]model.Program, error)
	FindById(id int) (model.Program, error)
	FindAllRelation() ([]model.Program, error)
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

func (s *programService) FindAllRelation() ([]model.Program, error) {
	var programRelations, err = s.programRepository.FindAllRelation()

	return programRelations, err
}

func (s *programService) Create(programRequest request.CreateProgramRequest) (model.Program, error) {
	var program = model.Program{
		BidangUrusanId:          programRequest.BidangUrusanId,
		InstansiId:              programRequest.InstansiId,
		Sasaran:                 programRequest.Sasaran,
		IndikatorSasaran:        programRequest.IndikatorSasaran,
		Kebijakan:               programRequest.Kebijakan,
		NamaProgram:             programRequest.NamaProgram,
		IndikatorKinerjaProgram: programRequest.IndikatorKinerjaProgram,
		PaguProgram:             programRequest.PaguProgram,
	}

	var newProgram, err = s.programRepository.Create(program)

	return newProgram, err
}

func (s *programService) Update(id int, programRequest request.UpdateProgramRequest) (model.Program, error) {
	var program, err = s.programRepository.FindById(id)

	program.BidangUrusanId = programRequest.BidangUrusanId
	program.InstansiId = programRequest.InstansiId
	program.Sasaran = programRequest.Sasaran
	program.IndikatorSasaran = programRequest.IndikatorSasaran
	program.Kebijakan = programRequest.Kebijakan
	program.NamaProgram = programRequest.NamaProgram
	program.IndikatorKinerjaProgram = programRequest.IndikatorKinerjaProgram
	program.PaguProgram = programRequest.PaguProgram

	updatedProgram, err := s.programRepository.Update(program)

	return updatedProgram, err
}

func (s *programService) Delete(id int) (model.Program, error) {
	var program, err = s.programRepository.FindById(id)

	deletedProgram, err := s.programRepository.Delete(program)

	return deletedProgram, err
}
