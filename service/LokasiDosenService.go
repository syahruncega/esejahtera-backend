package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type LokasiDosenService interface {
	FindAll() ([]model.LokasiDosen, error)
	FindById(id int) (model.LokasiDosen, error)
	FindAllRelation() ([]model.LokasiDosen, error)
	FindRelationByDosenId(dosenId int) ([]model.LokasiDosen, error)
	Create(lokasiDosenRequest request.CreateLokasiDosenRequest) (model.LokasiDosen, error)
	Update(id int, lokasiDosenRequest request.UpdateLokasiDosenRequest) (model.LokasiDosen, error)
	Delete(id int) (model.LokasiDosen, error)
}

type lokasiDosenService struct {
	lokasiDosenRepository repository.LokasiDosenRepository
}

func NewLokasiDosenService(lokasiDosenRepository repository.LokasiDosenRepository) *lokasiDosenService {
	return &lokasiDosenService{lokasiDosenRepository}
}

func (s *lokasiDosenService) FindAll() ([]model.LokasiDosen, error) {
	var lokasiDosens, err = s.lokasiDosenRepository.FindAll()

	return lokasiDosens, err
}

func (s *lokasiDosenService) FindById(id int) (model.LokasiDosen, error) {
	var lokasiDosen, err = s.lokasiDosenRepository.FindById(id)

	return lokasiDosen, err
}

func (s *lokasiDosenService) FindAllRelation() ([]model.LokasiDosen, error) {
	var lokasiDosens, err = s.lokasiDosenRepository.FindAllRelation()

	return lokasiDosens, err
}

func (s *lokasiDosenService) FindRelationByDosenId(dosenId int) ([]model.LokasiDosen, error) {
	var lokasiDosens, err = s.lokasiDosenRepository.FindRelationByDosenId(dosenId)

	return lokasiDosens, err
}

func (s *lokasiDosenService) Create(lokasiDosenRequest request.CreateLokasiDosenRequest) (model.LokasiDosen, error) {
	var lokasiDosen = model.LokasiDosen{
		DosenId:         lokasiDosenRequest.DosenId,
		KabupatenKotaId: lokasiDosenRequest.KabupatenKotaId,
		KecamatanId:     lokasiDosenRequest.KecamatanId,
		KelurahanId:     lokasiDosenRequest.KelurahanId,
	}

	newLokasiDosen, err := s.lokasiDosenRepository.Create(lokasiDosen)

	return newLokasiDosen, err
}

func (s *lokasiDosenService) Update(id int, lokasiDosenRequest request.UpdateLokasiDosenRequest) (model.LokasiDosen, error) {
	var lokasiDosen, err = s.lokasiDosenRepository.FindById(id)

	lokasiDosen.DosenId = lokasiDosenRequest.DosenId
	lokasiDosen.KabupatenKotaId = lokasiDosenRequest.KabupatenKotaId
	lokasiDosen.KecamatanId = lokasiDosenRequest.KecamatanId
	lokasiDosen.KelurahanId = lokasiDosenRequest.KelurahanId

	updatedLokasiDosen, err := s.lokasiDosenRepository.Update(lokasiDosen)

	return updatedLokasiDosen, err
}

func (s *lokasiDosenService) Delete(id int) (model.LokasiDosen, error) {
	var lokasiDosen, err = s.lokasiDosenRepository.FindById(id)

	deletedLokasiDosen, err := s.lokasiDosenRepository.Delete(lokasiDosen)

	return deletedLokasiDosen, err
}
