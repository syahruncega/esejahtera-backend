package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type RencanaKegiatanService interface {
	FindAll() ([]model.RencanaKegiatan, error)
	FindById(id int) (model.RencanaKegiatan, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.RencanaKegiatan, error)
	Create(rencanaKegiatanRequest request.CreateRencanaKegiatanRequest) (model.RencanaKegiatan, error)
	Update(idi int, rencanaKegiatanRequest request.UpdateRencanaKegiatanRequest) (model.RencanaKegiatan, error)
	Delete(id int) (model.RencanaKegiatan, error)
}

type rencanaKegiatanService struct {
	rencanaKegiatanRepository repository.RencanaKegiatanRepository
}

func NewRencanaKegiatanService(rencanaKegiatanRepository repository.RencanaKegiatanRepository) *rencanaKegiatanService {
	return &rencanaKegiatanService{rencanaKegiatanRepository}
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

func (s *rencanaKegiatanService) Create(rencanaKegiatanRequest request.CreateRencanaKegiatanRequest) (model.RencanaKegiatan, error) {
	var rencanaKegiatan = model.RencanaKegiatan{
		RencanaProgramId: rencanaKegiatanRequest.RencanaProgramId,
		KegiatanId:       rencanaKegiatanRequest.KegiatanId,
		PaguKegiatan:     rencanaKegiatanRequest.PaguKegiatan,
		Tipe:             rencanaKegiatanRequest.Tipe,
	}

	newRencanaKegiatan, err := s.rencanaKegiatanRepository.Create(rencanaKegiatan)

	return newRencanaKegiatan, err
}

func (s *rencanaKegiatanService) Update(id int, rencanaKegiatanRequest request.UpdateRencanaKegiatanRequest) (model.RencanaKegiatan, error) {
	var rencanaKegiatan, err = s.rencanaKegiatanRepository.FindById(id)

	rencanaKegiatan.RencanaProgramId = rencanaKegiatanRequest.RencanaProgramId
	rencanaKegiatan.KegiatanId = rencanaKegiatanRequest.KegiatanId
	rencanaKegiatan.PaguKegiatan = rencanaKegiatanRequest.PaguKegiatan
	rencanaKegiatan.Tipe = rencanaKegiatanRequest.Tipe

	updatedRencanaKegiatan, err := s.rencanaKegiatanRepository.Update(rencanaKegiatan)

	return updatedRencanaKegiatan, err
}

func (s *rencanaKegiatanService) Delete(id int) (model.RencanaKegiatan, error) {
	var rencanaKegiatan, err = s.rencanaKegiatanRepository.FindById(id)

	deletedRencanaKegiatan, err := s.rencanaKegiatanRepository.Delete(rencanaKegiatan)

	return deletedRencanaKegiatan, err
}