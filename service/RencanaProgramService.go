package service

import (
	"errors"
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type RencanaProgramService interface {
	FindAll() ([]model.RencanaProgram, error)
	FindById(id int) (model.RencanaProgram, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.RencanaProgram, error)
	CountJumlahRencanaProgram(tahun string, tipe string) (int64, error)
	CountJumlahRencanaProgramByIntansi(tahun string, tipe string, instansis []model.Instansi) []int64
	Create(rencanaProgramRequest request.CreateRencanaProgramRequest) (model.RencanaProgram, error)
	Update(id int, rencanaProgramRequest request.UpdateRencanaProgramRequest) (model.RencanaProgram, error)
	Delete(id int) (model.RencanaProgram, error)
}

type rencanaProgramService struct {
	rencanaProgramRepository  repository.RencanaProgramRepository
	rencanaKegiatanRepository repository.RencanaKegiatanRepository
}

func NewRencanaProgramService(rencanaProgramRepository repository.RencanaProgramRepository, rencanaKegiatanRepository repository.RencanaKegiatanRepository) *rencanaProgramService {
	return &rencanaProgramService{rencanaProgramRepository, rencanaKegiatanRepository}
}

func (s *rencanaProgramService) FindAll() ([]model.RencanaProgram, error) {
	var rencanaPrograms, err = s.rencanaProgramRepository.FindAll()

	return rencanaPrograms, err
}

func (s *rencanaProgramService) FindById(id int) (model.RencanaProgram, error) {
	var rencanaProgram, err = s.rencanaProgramRepository.FindById(id)

	return rencanaProgram, err
}

func (s *rencanaProgramService) FindBySearch(whereClause map[string]interface{}) ([]model.RencanaProgram, error) {
	var rencanaPrograms, err = s.rencanaProgramRepository.FindBySearch(whereClause)

	return rencanaPrograms, err
}

func (s *rencanaProgramService) CountJumlahRencanaProgram(tahun string, tipe string) (int64, error) {
	var jumlahRencanaProgram, err = s.rencanaProgramRepository.CountJumlahRencanaProgram(tahun, tipe)

	return jumlahRencanaProgram, err
}

func (s *rencanaProgramService) CountJumlahRencanaProgramByIntansi(tahun string, tipe string, instansis []model.Instansi) []int64 {
	var jumlahRencanaProgramInstansi = s.rencanaProgramRepository.CountJumlahRencanaProgramByIntansi(tahun, tipe, instansis)

	return jumlahRencanaProgramInstansi
}

func (s *rencanaProgramService) Create(rencanaProgramRequest request.CreateRencanaProgramRequest) (model.RencanaProgram, error) {
	var rencanaProgram = model.RencanaProgram{
		InstansiId:  rencanaProgramRequest.InstansiId,
		ProgramId:   rencanaProgramRequest.ProgramId,
		PaguProgram: rencanaProgramRequest.PaguProgram,
		Tipe:        rencanaProgramRequest.Tipe,
		Tahun:       rencanaProgramRequest.Tahun,
	}

	newRencanaProgram, err := s.rencanaProgramRepository.Create(rencanaProgram)

	return newRencanaProgram, err
}

func (s *rencanaProgramService) Update(id int, rencanaProgramRequest request.UpdateRencanaProgramRequest) (model.RencanaProgram, error) {
	var rencanaProgram, err = s.rencanaProgramRepository.FindById(id)
	if err != nil {
		return rencanaProgram, err
	}

	totalPaguRencanaKegiatan, err := s.rencanaKegiatanRepository.SumPaguRencanaKegiatan(id)
	if err != nil {
		return rencanaProgram, err
	}

	if rencanaProgramRequest.PaguProgram >= totalPaguRencanaKegiatan {

		rencanaProgram.InstansiId = rencanaProgramRequest.InstansiId
		rencanaProgram.ProgramId = rencanaProgramRequest.ProgramId
		rencanaProgram.PaguProgram = rencanaProgramRequest.PaguProgram
		rencanaProgram.Tipe = rencanaProgramRequest.Tipe
		rencanaProgram.Tahun = rencanaProgramRequest.Tahun

		updatedRencanaProgram, err := s.rencanaProgramRepository.Update(rencanaProgram)

		return updatedRencanaProgram, err
	} else {
		return rencanaProgram, errors.New("pagu parent lebih kecil dari pagu child")
	}

}

func (s *rencanaProgramService) Delete(id int) (model.RencanaProgram, error) {
	var rencanaProgram, err = s.rencanaProgramRepository.FindById(id)

	deletedRencanaProgram, err := s.rencanaProgramRepository.Delete(rencanaProgram)

	return deletedRencanaProgram, err
}
