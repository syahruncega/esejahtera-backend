package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type RencanaSubKegiatanService interface {
	FindAll() ([]model.RencanaSubKegiatan, error)
	FindById(id int) (model.RencanaSubKegiatan, error)
	FindBySearch(whereClase map[string]interface{}) ([]model.RencanaSubKegiatan, error)
	Create(rencanaSubKegiatanRequest request.CreateRencanaSubKegiatanRequest) (model.RencanaSubKegiatan, error)
	Update(id int, rencanaSubKegiatanRequest request.UpdateRencanaSubKegiatanRequest) (model.RencanaSubKegiatan, error)
	Delete(id int) (model.RencanaSubKegiatan, error)
}

type rencanaSubKegiatanService struct {
	rencanaSubKegiatanRepository repository.RencanaSubKegiatanRepository
}

func NewRencanaSubKegiatanService(rencanaSubKegiatanRepository repository.RencanaSubKegiatanRepository) *rencanaSubKegiatanService {
	return &rencanaSubKegiatanService{rencanaSubKegiatanRepository}
}

func (s *rencanaSubKegiatanService) FindAll() ([]model.RencanaSubKegiatan, error) {
	var rencanaSubKegiatans, err = s.rencanaSubKegiatanRepository.FindAll()

	return rencanaSubKegiatans, err
}

func (s *rencanaSubKegiatanService) FindById(id int) (model.RencanaSubKegiatan, error) {
	var rencanaSubKegiatan, err = s.rencanaSubKegiatanRepository.FindById(id)

	return rencanaSubKegiatan, err
}

func (s *rencanaSubKegiatanService) FindBySearch(whereClause map[string]interface{}) ([]model.RencanaSubKegiatan, error) {
	var rencanaSubKegiatans, err = s.rencanaSubKegiatanRepository.FindBySearch(whereClause)

	return rencanaSubKegiatans, err
}

func (s *rencanaSubKegiatanService) Create(rencanaSubKegiatanRequest request.CreateRencanaSubKegiatanRequest) (model.RencanaSubKegiatan, error) {
	var rencanaSubKegiatan = model.RencanaSubKegiatan{
		RencanaKegiatanId: rencanaSubKegiatanRequest.RencanaKegiatanId,
		SubKegiatanId:     rencanaSubKegiatanRequest.SubKegiatanId,
		PaguSubKegiatan:   rencanaSubKegiatanRequest.PaguSubKegiatan,
		Tipe:              rencanaSubKegiatanRequest.Tipe,
	}

	newRencanaSubKegiatan, err := s.rencanaSubKegiatanRepository.Create(rencanaSubKegiatan)

	return newRencanaSubKegiatan, err
}

func (s *rencanaSubKegiatanService) Update(id int, rencanaSubKegiatanRequest request.UpdateRencanaSubKegiatanRequest) (model.RencanaSubKegiatan, error) {
	var rencanaSubKegiatan, err = s.rencanaSubKegiatanRepository.FindById(id)

	rencanaSubKegiatan.RencanaKegiatanId = rencanaSubKegiatanRequest.RencanaKegiatanId
	rencanaSubKegiatan.SubKegiatanId = rencanaSubKegiatanRequest.SubKegiatanId
	rencanaSubKegiatan.PaguSubKegiatan = rencanaSubKegiatanRequest.PaguSubKegiatan
	rencanaSubKegiatan.Tipe = rencanaSubKegiatanRequest.Tipe

	updatedRencanaSubKegiatan, err := s.rencanaSubKegiatanRepository.Update(rencanaSubKegiatan)

	return updatedRencanaSubKegiatan, err
}

func (s *rencanaSubKegiatanService) Delete(id int) (model.RencanaSubKegiatan, error) {
	var rencanaSubKegiatan, err = s.rencanaSubKegiatanRepository.FindById(id)

	deletedRencanaSubKegiatan, err := s.rencanaSubKegiatanRepository.Delete(rencanaSubKegiatan)

	return deletedRencanaSubKegiatan, err
}
