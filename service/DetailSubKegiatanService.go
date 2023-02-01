package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type DetailSubKegiatanService interface {
	FindAll() ([]model.DetailSubKegiatan, error)
	FindById(id int) (model.DetailSubKegiatan, error)
	FindByKegiatanId(kegiatanId string) ([]model.DetailSubKegiatan, error)
	Create(detailSubKegiatanRequest request.CreateDetailSubKegiatanRequest) (model.DetailSubKegiatan, error)
	Update(id int, detailSubKegiatanRequest request.UpdateDetailSubKegiatanRequest) (model.DetailSubKegiatan, error)
	Delete(id int) (model.DetailSubKegiatan, error)
}

type detailSubKegiatanService struct {
	detailSubKegiatanRepository repository.DetailSubKegiatanRepository
}

func NewDetailSubKegiatanService(detailSubKegiatanRepository repository.DetailSubKegiatanRepository) *detailSubKegiatanService {
	return &detailSubKegiatanService{detailSubKegiatanRepository}
}

func (s *detailSubKegiatanService) FindAll() ([]model.DetailSubKegiatan, error) {
	var detailSubKegiatans, err = s.detailSubKegiatanRepository.FindAll()

	return detailSubKegiatans, err
}

func (s *detailSubKegiatanService) FindById(id int) (model.DetailSubKegiatan, error) {
	var detailSubKegiatan, err = s.detailSubKegiatanRepository.FindById(id)

	return detailSubKegiatan, err
}

func (s *detailSubKegiatanService) FindByKegiatanId(kegiatanId string) ([]model.DetailSubKegiatan, error) {
	var detailSubKegiatans, err = s.detailSubKegiatanRepository.FindByKegiatanId(kegiatanId)

	return detailSubKegiatans, err
}

func (s *detailSubKegiatanService) Create(detailSubKegiatanRequest request.CreateDetailSubKegiatanRequest) (model.DetailSubKegiatan, error) {
	var detailSubKegiatan = model.DetailSubKegiatan{
		KegiatanId:    detailSubKegiatanRequest.KegiatanId,
		SubKegiatanId: detailSubKegiatanRequest.SubKegiatanId,
	}

	newDetailSubKegiatan, err := s.detailSubKegiatanRepository.Create(detailSubKegiatan)

	return newDetailSubKegiatan, err
}

func (s *detailSubKegiatanService) Update(id int, detailSubKegiatanRequest request.UpdateDetailSubKegiatanRequest) (model.DetailSubKegiatan, error) {
	var detailSubKegiatan, err = s.detailSubKegiatanRepository.FindById(id)

	detailSubKegiatan.KegiatanId = detailSubKegiatanRequest.KegiatanId
	detailSubKegiatan.SubKegiatanId = detailSubKegiatanRequest.SubKegiatanId

	updatedDetailSubKegiatan, err := s.detailSubKegiatanRepository.Update(detailSubKegiatan)

	return updatedDetailSubKegiatan, err
}

func (s *detailSubKegiatanService) Delete(id int) (model.DetailSubKegiatan, error) {
	var detailSubKegiatan, err = s.detailSubKegiatanRepository.FindById(id)

	deletedDetailSubKegiatan, err := s.detailSubKegiatanRepository.Delete(detailSubKegiatan)

	return deletedDetailSubKegiatan, err
}
