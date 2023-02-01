package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type IndikatorSubKegiatanService interface {
	FindAll() ([]model.IndikatorSubKegiatan, error)
	FindById(id int) (model.IndikatorSubKegiatan, error)
	Create(indikatorSubKegiatanRequest request.CreateIndikatorSubKegiatanRequest) (model.IndikatorSubKegiatan, error)
	Update(id int, indikatorSubKegiatanRequest request.UpdateIndikatorSubKegiatanRequest) (model.IndikatorSubKegiatan, error)
	Delete(id int) (model.IndikatorSubKegiatan, error)
}

type indikatorSubKegiatanService struct {
	indikatorSubKegiatanRepository repository.IndikatorSubKegiatanRepository
}

func NewIndikatorSubKegiatanService(indikatorSubKegiatanRepository repository.IndikatorSubKegiatanRepository) *indikatorSubKegiatanService {
	return &indikatorSubKegiatanService{indikatorSubKegiatanRepository}
}

func (s *indikatorSubKegiatanService) FindAll() ([]model.IndikatorSubKegiatan, error) {
	var indikatorSubKegiatans, err = s.indikatorSubKegiatanRepository.FindAll()

	return indikatorSubKegiatans, err
}

func (s *indikatorSubKegiatanService) FindById(id int) (model.IndikatorSubKegiatan, error) {
	var indikatorSubKegiatan, err = s.indikatorSubKegiatanRepository.FindById(id)

	return indikatorSubKegiatan, err
}

func (s *indikatorSubKegiatanService) Create(indikatorSubKegiatanRequest request.CreateIndikatorSubKegiatanRequest) (model.IndikatorSubKegiatan, error) {
	var indikatorSubKegiatan = model.IndikatorSubKegiatan{
		KegiatanId:                  indikatorSubKegiatanRequest.KegiatanId,
		SubKegiatanId:               indikatorSubKegiatanRequest.SubKegiatanId,
		IndikatorKinerjaSubKegiatan: indikatorSubKegiatanRequest.IndikatorKinerjaSubKegiatan,
		PaguSubKegiatan:             indikatorSubKegiatanRequest.PaguSubKegiatan,
	}

	var newIndikatorSubKegiatan, err = s.indikatorSubKegiatanRepository.Create(indikatorSubKegiatan)

	return newIndikatorSubKegiatan, err
}

func (s *indikatorSubKegiatanService) Update(id int, indikatorSubKegiatanRequest request.UpdateIndikatorSubKegiatanRequest) (model.IndikatorSubKegiatan, error) {
	var subKegiatan, err = s.indikatorSubKegiatanRepository.FindById(id)

	subKegiatan.KegiatanId = indikatorSubKegiatanRequest.KegiatanId
	subKegiatan.SubKegiatanId = indikatorSubKegiatanRequest.SubKegiatanId
	subKegiatan.IndikatorKinerjaSubKegiatan = indikatorSubKegiatanRequest.IndikatorKinerjaSubKegiatan
	subKegiatan.PaguSubKegiatan = indikatorSubKegiatanRequest.PaguSubKegiatan

	updatedIndikatorSubkegiatan, err := s.indikatorSubKegiatanRepository.Update(subKegiatan)

	return updatedIndikatorSubkegiatan, err
}

func (s *indikatorSubKegiatanService) Delete(id int) (model.IndikatorSubKegiatan, error) {
	var indikatorSubKegiatan, err = s.indikatorSubKegiatanRepository.FindById(id)

	deletedIndikatorSubKegiatan, err := s.indikatorSubKegiatanRepository.Delete(indikatorSubKegiatan)

	return deletedIndikatorSubKegiatan, err
}
