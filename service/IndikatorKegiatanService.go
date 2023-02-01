package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type IndikatorKegiatanService interface {
	FindAll() ([]model.IndikatorKegiatan, error)
	FindById(id int) (model.IndikatorKegiatan, error)
	Create(indikatorKegiatanRequest request.CreateIndikatorKegiatanRequest) (model.IndikatorKegiatan, error)
	Update(id int, indikatorKegiatanRequest request.UpdateIndikatorKegiatanRequest) (model.IndikatorKegiatan, error)
	Delete(id int) (model.IndikatorKegiatan, error)
}

type indikatorKegiatanService struct {
	indikatorKegiatanRepository repository.IndikatorKegiatanRepository
}

func NewIndikatorKegiatanService(indikatorKegiatanRepository repository.IndikatorKegiatanRepository) *indikatorKegiatanService {
	return &indikatorKegiatanService{indikatorKegiatanRepository}
}

func (s *indikatorKegiatanService) FindAll() ([]model.IndikatorKegiatan, error) {
	var indikatorKegiatans, err = s.indikatorKegiatanRepository.FindAll()

	return indikatorKegiatans, err
}

func (s *indikatorKegiatanService) FindById(id int) (model.IndikatorKegiatan, error) {
	var indikatorKegiatan, err = s.indikatorKegiatanRepository.FindById(id)

	return indikatorKegiatan, err
}

func (s *indikatorKegiatanService) Create(indikatorKegiatanRequest request.CreateIndikatorKegiatanRequest) (model.IndikatorKegiatan, error) {
	var indikatorKegiatan = model.IndikatorKegiatan{
		ProgramId:                indikatorKegiatanRequest.ProgramId,
		KegiatanId:               indikatorKegiatanRequest.KegiatanId,
		IndikatorKinerjaKegiatan: indikatorKegiatanRequest.IndikatorKinerjaKegiatan,
		PaguKegiatan:             indikatorKegiatanRequest.PaguKegiatan,
	}

	var newIndikatorKegiatan, err = s.indikatorKegiatanRepository.Create(indikatorKegiatan)

	return newIndikatorKegiatan, err
}

func (s *indikatorKegiatanService) Update(id int, indikatorKegiatanRequest request.UpdateIndikatorKegiatanRequest) (model.IndikatorKegiatan, error) {
	var kegiatan, err = s.indikatorKegiatanRepository.FindById(id)

	kegiatan.ProgramId = indikatorKegiatanRequest.ProgramId
	kegiatan.KegiatanId = indikatorKegiatanRequest.KegiatanId
	kegiatan.IndikatorKinerjaKegiatan = indikatorKegiatanRequest.IndikatorKinerjaKegiatan
	kegiatan.PaguKegiatan = indikatorKegiatanRequest.PaguKegiatan

	updatedIndikatorKegiatan, err := s.indikatorKegiatanRepository.Update(kegiatan)

	return updatedIndikatorKegiatan, err
}

func (s *indikatorKegiatanService) Delete(id int) (model.IndikatorKegiatan, error) {
	var indikatorKegiatan, err = s.indikatorKegiatanRepository.FindById(id)

	deletedIndikatorKegiatan, err := s.indikatorKegiatanRepository.Delete(indikatorKegiatan)

	return deletedIndikatorKegiatan, err
}
