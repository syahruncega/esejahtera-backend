package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type DetailKegiatanService interface {
	FindAll() ([]model.DetailKegiatan, error)
	FindById(id int) (model.DetailKegiatan, error)
	Create(detailKegiatanRequest request.CreateDetailKegiatanRequest) (model.DetailKegiatan, error)
	Update(idi int, detailKegiatanRequest request.UpdateDetailKegiatanRequest) (model.DetailKegiatan, error)
	Delete(id int) (model.DetailKegiatan, error)
}

type detailKegiatanService struct {
	detailKegiatanRepository repository.DetailKegiatanRepository
}

func NewDetailKegiatanService(detailKegiatanRepository repository.DetailKegiatanRepository) *detailKegiatanService {
	return &detailKegiatanService{detailKegiatanRepository}
}

func (s *detailKegiatanService) FindAll() ([]model.DetailKegiatan, error) {
	var detailKegiatans, err = s.detailKegiatanRepository.FindAll()

	return detailKegiatans, err
}

func (s *detailKegiatanService) FindById(id int) (model.DetailKegiatan, error) {
	var detailKegiatan, err = s.detailKegiatanRepository.FindById(id)

	return detailKegiatan, err
}

func (s *detailKegiatanService) Create(detailKegiatanRequest request.CreateDetailKegiatanRequest) (model.DetailKegiatan, error) {
	var detailKegiatan = model.DetailKegiatan{
		KegiatanId:   detailKegiatanRequest.KegiatanId,
		PaguKegiatan: detailKegiatanRequest.PaguKegiatan,
		Tipe:         detailKegiatanRequest.Tipe,
	}

	newDetailKegiatan, err := s.detailKegiatanRepository.Create(detailKegiatan)

	return newDetailKegiatan, err
}

func (s *detailKegiatanService) Update(id int, detailKegiatanRequest request.UpdateDetailKegiatanRequest) (model.DetailKegiatan, error) {
	var detailKegiatan, err = s.detailKegiatanRepository.FindById(id)

	detailKegiatan.KegiatanId = detailKegiatanRequest.KegiatanId
	detailKegiatan.PaguKegiatan = detailKegiatanRequest.PaguKegiatan
	detailKegiatan.Tipe = detailKegiatanRequest.Tipe

	updatedDetailKegiatan, err := s.detailKegiatanRepository.Update(detailKegiatan)

	return updatedDetailKegiatan, err
}

func (s *detailKegiatanService) Delete(id int) (model.DetailKegiatan, error) {
	var detailKegiatan, err = s.detailKegiatanRepository.FindById(id)

	deletedDetailKegiatan, err := s.detailKegiatanRepository.Delete(detailKegiatan)

	return deletedDetailKegiatan, err
}
