package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type SubKegiatanService interface {
	FindAll() ([]model.Sub_Kegiatan, error)
	FindById(id int) (model.Sub_Kegiatan, error)
	FindAllRelation() ([]model.Sub_Kegiatan, error)
	Create(subKegiatanRequest request.CreateSub_KegiatanRequest) (model.Sub_Kegiatan, error)
	Update(id int, subKegiatanRequest request.UpdateSub_KegiatanRequest) (model.Sub_Kegiatan, error)
	Delete(id int) (model.Sub_Kegiatan, error)
}

type subKegiatanService struct {
	subKegiatanRepository repository.SubKegiatanRepository
}

func NewSubKegiatanService(subKegiatanRepository repository.SubKegiatanRepository) *subKegiatanService {
	return &subKegiatanService{subKegiatanRepository}
}

func (s *subKegiatanService) FindAll() ([]model.Sub_Kegiatan, error) {
	var subKegiatans, err = s.subKegiatanRepository.FindAll()

	return subKegiatans, err
}

func (s *subKegiatanService) FindById(id int) (model.Sub_Kegiatan, error) {
	var subKegiatan, err = s.subKegiatanRepository.FindById(id)

	return subKegiatan, err
}

func (s *subKegiatanService) FindAllRelation() ([]model.Sub_Kegiatan, error) {
	var subKegiatanRelations, err = s.subKegiatanRepository.FindAllRelation()

	return subKegiatanRelations, err
}

func (s *subKegiatanService) Create(subKegiatanRequest request.CreateSub_KegiatanRequest) (model.Sub_Kegiatan, error) {
	var subKegiatan = model.Sub_Kegiatan{
		NamaSubKegiatan:             subKegiatanRequest.NamaSubKegiatan,
		IndikatorKinerjaSubKegiatan: subKegiatanRequest.IndikatorKinerjaSubKegiatan,
		PaguSubKegiatan:             subKegiatanRequest.PaguSubKegiatan,
		KegiatanId:                  subKegiatanRequest.KegiatanId,
	}

	var newSubKegiatan, err = s.subKegiatanRepository.Create(subKegiatan)

	return newSubKegiatan, err
}

func (s *subKegiatanService) Update(id int, subKegiatanRequest request.UpdateSub_KegiatanRequest) (model.Sub_Kegiatan, error) {
	var subKegiatan, err = s.subKegiatanRepository.FindById(id)

	subKegiatan.NamaSubKegiatan = subKegiatanRequest.NamaSubKegiatan
	subKegiatan.IndikatorKinerjaSubKegiatan = subKegiatanRequest.IndikatorKinerjaSubKegiatan
	subKegiatan.PaguSubKegiatan = subKegiatanRequest.PaguSubKegiatan
	subKegiatan.KegiatanId = subKegiatanRequest.KegiatanId

	updatedSubKegiatan, err := s.subKegiatanRepository.Update(subKegiatan)

	return updatedSubKegiatan, err
}

func (s *subKegiatanService) Delete(id int) (model.Sub_Kegiatan, error) {
	var subKegiatan, err = s.subKegiatanRepository.FindById(id)

	deletedSubKegiatan, err := s.subKegiatanRepository.Delete(subKegiatan)

	return deletedSubKegiatan, err
}
