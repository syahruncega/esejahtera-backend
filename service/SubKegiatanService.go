package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type SubKegiatanService interface {
	FindAll() ([]model.SubKegiatan, error)
	FindById(id string) (model.SubKegiatan, error)
	Create(subKegiatanRequest request.CreateSubKegiatanRequest) (model.SubKegiatan, error)
	Update(id string, subKegiatanRequest request.UpdateSubKegiatanRequest) (model.SubKegiatan, error)
	Delete(id string) (model.SubKegiatan, error)
}

type subKegiatanService struct {
	subKegiatanRepository repository.SubKegiatanRepository
}

func NewSubKegiatanService(subKegiatanRepository repository.SubKegiatanRepository) *subKegiatanService {
	return &subKegiatanService{subKegiatanRepository}
}

func (s *subKegiatanService) FindAll() ([]model.SubKegiatan, error) {
	var subKegiatans, err = s.subKegiatanRepository.FindAll()

	return subKegiatans, err
}

func (s *subKegiatanService) FindById(id string) (model.SubKegiatan, error) {
	var subKegiatan, err = s.subKegiatanRepository.FindById(id)

	return subKegiatan, err
}

func (s *subKegiatanService) Create(subKegiatanRequest request.CreateSubKegiatanRequest) (model.SubKegiatan, error) {
	var subKegiatan = model.SubKegiatan{
		Id:              subKegiatanRequest.Id,
		NamaSubKegiatan: subKegiatanRequest.NamaSubKegiatan,
	}

	newSubKegiatan, err := s.subKegiatanRepository.Create(subKegiatan)

	return newSubKegiatan, err
}

func (s *subKegiatanService) Update(id string, subKegiatanRequest request.UpdateSubKegiatanRequest) (model.SubKegiatan, error) {
	var subKegiatan, err = s.subKegiatanRepository.FindById(id)

	subKegiatan.Id = subKegiatanRequest.Id
	subKegiatan.NamaSubKegiatan = subKegiatanRequest.NamaSubKegiatan

	updatedSubKegiatan, err := s.subKegiatanRepository.Update(subKegiatan)

	return updatedSubKegiatan, err
}

func (s *subKegiatanService) Delete(id string) (model.SubKegiatan, error) {
	var subKegiatan, err = s.subKegiatanRepository.FindById(id)

	deletedSubKegiatan, err := s.subKegiatanRepository.Delete(subKegiatan)

	return deletedSubKegiatan, err
}
