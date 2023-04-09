package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type KegiatanOnSubKegiatanService interface {
	FindAll() ([]model.KegiatanOnSubKegiatan, error)
	FindById(id int) (model.KegiatanOnSubKegiatan, error)
	FindByKegiatanId(kegiatanId int) ([]model.KegiatanOnSubKegiatan, error)
	FindBySearch(whereClause map[string]interface{}, tahun string) ([]model.KegiatanOnSubKegiatan, error)
	Create(kegiatanOnSubKegiatanRequest request.CreateKegiatanOnSubKegiatanRequest) (model.KegiatanOnSubKegiatan, error)
	Update(id int, kegiatanOnSubKegiatanRequest request.UpdateKegiatanOnSubKegiatanRequest) (model.KegiatanOnSubKegiatan, error)
	Delete(id int) (model.KegiatanOnSubKegiatan, error)
}

type kegiatanOnSubKegiatanService struct {
	kegiatanOnSubKegiatanRepository repository.KegiatanOnSubKegiatanRepository
}

func NewKegiatanOnSubKegiatanService(kegiatanOnSubKegiatanRepository repository.KegiatanOnSubKegiatanRepository) *kegiatanOnSubKegiatanService {
	return &kegiatanOnSubKegiatanService{kegiatanOnSubKegiatanRepository}
}

func (s *kegiatanOnSubKegiatanService) FindAll() ([]model.KegiatanOnSubKegiatan, error) {
	var kegiatanOnSubKegiatans, err = s.kegiatanOnSubKegiatanRepository.FindAll()

	return kegiatanOnSubKegiatans, err
}

func (s *kegiatanOnSubKegiatanService) FindById(id int) (model.KegiatanOnSubKegiatan, error) {
	var kegiatanOnSubKegiatan, err = s.kegiatanOnSubKegiatanRepository.FindById(id)

	return kegiatanOnSubKegiatan, err
}

func (s *kegiatanOnSubKegiatanService) FindByKegiatanId(kegiatanId int) ([]model.KegiatanOnSubKegiatan, error) {
	var kegiatanOnSubKegiatans, err = s.kegiatanOnSubKegiatanRepository.FindByKegiatanId(kegiatanId)

	return kegiatanOnSubKegiatans, err
}

func (s *kegiatanOnSubKegiatanService) FindBySearch(whereClause map[string]interface{}, tahun string) ([]model.KegiatanOnSubKegiatan, error) {
	var kegiatanOnSubKegiatans, err = s.kegiatanOnSubKegiatanRepository.FindBySearch(whereClause, tahun)

	return kegiatanOnSubKegiatans, err
}

func (s *kegiatanOnSubKegiatanService) Create(kegiatanOnSubKegiatanRequest request.CreateKegiatanOnSubKegiatanRequest) (model.KegiatanOnSubKegiatan, error) {
	var kegiatanOnSubKegiatan = model.KegiatanOnSubKegiatan{
		KegiatanId:    kegiatanOnSubKegiatanRequest.KegiatanId,
		SubKegiatanId: kegiatanOnSubKegiatanRequest.SubKegiatanId,
	}

	newKegiatanOnSubKegiatan, err := s.kegiatanOnSubKegiatanRepository.Create(kegiatanOnSubKegiatan)

	return newKegiatanOnSubKegiatan, err
}

func (s *kegiatanOnSubKegiatanService) Update(id int, kegiatanOnSubKegiatanRequest request.UpdateKegiatanOnSubKegiatanRequest) (model.KegiatanOnSubKegiatan, error) {
	var kegiatanOnSubKegiatan, err = s.kegiatanOnSubKegiatanRepository.FindById(id)

	kegiatanOnSubKegiatan.KegiatanId = kegiatanOnSubKegiatanRequest.KegiatanId
	kegiatanOnSubKegiatan.SubKegiatanId = kegiatanOnSubKegiatanRequest.SubKegiatanId

	updatedKegiatanOnSubKegiatan, err := s.kegiatanOnSubKegiatanRepository.FindById(id)

	return updatedKegiatanOnSubKegiatan, err
}

func (s *kegiatanOnSubKegiatanService) Delete(id int) (model.KegiatanOnSubKegiatan, error) {
	var kegiatanOnSubKegiatan, err = s.kegiatanOnSubKegiatanRepository.FindById(id)

	deletedKegiatanOnSubKegiatan, err := s.kegiatanOnSubKegiatanRepository.Delete(kegiatanOnSubKegiatan)

	return deletedKegiatanOnSubKegiatan, err
}
