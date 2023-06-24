package service

import (
	"errors"
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type RencanaKegiatanService interface {
	FindAll() ([]model.RencanaKegiatan, error)
	FindById(id int) (model.RencanaKegiatan, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.RencanaKegiatan, error)
	SumAllPaguRencanaKegiatan(tahun, tipe string) int64
	SumPaguRencanaKegiatanByInstansi(tahun, tipe string, instansis []model.Instansi) []int64
	CountJumlahRencanaKegiatan(tahun string, tipe string) (int64, error)
	CountJumlahRencanaKegiatanAllInstansi(tahun string, tipe string, instansis []model.Instansi) []int64
	Create(rencanaKegiatanRequest request.CreateRencanaKegiatanRequest) (model.RencanaKegiatan, error)
	Update(id int, rencanaKegiatanRequest request.UpdateRencanaKegiatanRequest) (model.RencanaKegiatan, error)
	Delete(id int) (model.RencanaKegiatan, error)
}

type rencanaKegiatanService struct {
	rencanaKegiatanRepository    repository.RencanaKegiatanRepository
	rencanaProgramRepository     repository.RencanaProgramRepository
	rencanaSubKegiatanRepository repository.RencanaSubKegiatanRepository
}

func NewRencanaKegiatanService(rencanaKegiatanRepository repository.RencanaKegiatanRepository, rencanaProgramRepository repository.RencanaProgramRepository, rencanaSubKegiatanRepository repository.RencanaSubKegiatanRepository) *rencanaKegiatanService {
	return &rencanaKegiatanService{rencanaKegiatanRepository, rencanaProgramRepository, rencanaSubKegiatanRepository}
}

func (s *rencanaKegiatanService) FindAll() ([]model.RencanaKegiatan, error) {
	var rencanaKegiatans, err = s.rencanaKegiatanRepository.FindAll()

	return rencanaKegiatans, err
}

func (s *rencanaKegiatanService) FindById(id int) (model.RencanaKegiatan, error) {
	var rencanaKegiatan, err = s.rencanaKegiatanRepository.FindById(id)

	return rencanaKegiatan, err
}

func (s *rencanaKegiatanService) FindBySearch(whereClause map[string]interface{}) ([]model.RencanaKegiatan, error) {
	var rencanaKegiatans, err = s.rencanaKegiatanRepository.FindBySearch(whereClause)

	return rencanaKegiatans, err
}

func (s *rencanaKegiatanService) SumAllPaguRencanaKegiatan(tahun, tipe string) int64 {
	var totalPaguRencanaKegiatan = s.rencanaKegiatanRepository.SumAllPaguRencanaKegiatan(tahun, tipe)

	return totalPaguRencanaKegiatan
}

func (s *rencanaKegiatanService) SumPaguRencanaKegiatanByInstansi(tahun, tipe string, instansis []model.Instansi) []int64 {
	var totalPaguRencanaKegiatan = s.rencanaKegiatanRepository.SumPaguRencanaKegiatanByInstansi(tahun, tipe, instansis)

	return totalPaguRencanaKegiatan
}

func (s *rencanaKegiatanService) CountJumlahRencanaKegiatan(tahun string, tipe string) (int64, error) {
	var jumlahRencanaKegiatan, err = s.rencanaKegiatanRepository.CountJumlahRencanaKegiatan(tahun, tipe)

	return jumlahRencanaKegiatan, err
}

func (s *rencanaKegiatanService) CountJumlahRencanaKegiatanAllInstansi(tahun string, tipe string, instansis []model.Instansi) []int64 {
	var jumlahRencanaKegiatan = s.rencanaKegiatanRepository.CountJumlahRencanaKegiatanAllInstansi(tahun, tipe, instansis)

	return jumlahRencanaKegiatan
}

func (s *rencanaKegiatanService) Create(rencanaKegiatanRequest request.CreateRencanaKegiatanRequest) (model.RencanaKegiatan, error) {
	var rencanaKegiatan = model.RencanaKegiatan{
		RencanaProgramId: rencanaKegiatanRequest.RencanaProgramId,
		KegiatanId:       rencanaKegiatanRequest.KegiatanId,
		PaguKegiatan:     rencanaKegiatanRequest.PaguKegiatan,
		Tipe:             rencanaKegiatanRequest.Tipe,
		Tahun:            rencanaKegiatanRequest.Tahun,
	}

	var rencanaProgram, _ = s.rencanaProgramRepository.FindById(rencanaKegiatan.RencanaProgramId)

	var totalPaguRencanaKegiatan, _ = s.rencanaKegiatanRepository.SumPaguRencanaKegiatan(rencanaKegiatan.RencanaProgramId)
	totalPaguRencanaKegiatan = totalPaguRencanaKegiatan + rencanaKegiatan.PaguKegiatan

	if rencanaProgram.PaguProgram >= totalPaguRencanaKegiatan {
		newRencanaKegiatan, err := s.rencanaKegiatanRepository.Create(rencanaKegiatan)
		return newRencanaKegiatan, err
	} else {
		return rencanaKegiatan, errors.New("pagu kegiatan melebihi pagu program")
	}
}

func (s *rencanaKegiatanService) Update(id int, rencanaKegiatanRequest request.UpdateRencanaKegiatanRequest) (model.RencanaKegiatan, error) {
	var rencanaKegiatan, _ = s.rencanaKegiatanRepository.FindById(id)

	var currentPagu = rencanaKegiatan.PaguKegiatan

	rencanaKegiatan.RencanaProgramId = rencanaKegiatanRequest.RencanaProgramId
	rencanaKegiatan.KegiatanId = rencanaKegiatanRequest.KegiatanId
	rencanaKegiatan.PaguKegiatan = rencanaKegiatanRequest.PaguKegiatan
	rencanaKegiatan.Tipe = rencanaKegiatanRequest.Tipe
	rencanaKegiatan.Tahun = rencanaKegiatanRequest.Tahun

	var rencanaProgram, _ = s.rencanaProgramRepository.FindById(rencanaKegiatan.RencanaProgramId)

	var totalPaguRencanaSubKegiatan, _ = s.rencanaSubKegiatanRepository.SumPaguRencanaSubKegiatan(id)

	var totalPaguRencanaKegiatan, _ = s.rencanaKegiatanRepository.SumPaguRencanaKegiatan(rencanaKegiatan.RencanaProgramId)
	totalPaguRencanaKegiatan = totalPaguRencanaKegiatan - currentPagu + rencanaKegiatanRequest.PaguKegiatan

	if rencanaKegiatanRequest.PaguKegiatan >= totalPaguRencanaSubKegiatan {
		if rencanaProgram.PaguProgram >= totalPaguRencanaKegiatan {
			updatedRencanaKegiatan, err := s.rencanaKegiatanRepository.Update(rencanaKegiatan)
			return updatedRencanaKegiatan, err
		} else {
			return rencanaKegiatan, errors.New("pagu kegiatan melebihi pagu programnya")
		}
	} else {
		return rencanaKegiatan, errors.New("pagu parent lebih kecil dari pagu child")
	}

}

func (s *rencanaKegiatanService) Delete(id int) (model.RencanaKegiatan, error) {
	var rencanaKegiatan, err = s.rencanaKegiatanRepository.FindById(id)

	deletedRencanaKegiatan, err := s.rencanaKegiatanRepository.Delete(rencanaKegiatan)

	return deletedRencanaKegiatan, err
}
