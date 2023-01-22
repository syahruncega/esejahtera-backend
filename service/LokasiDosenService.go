package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type LokasiDosenService interface {
	FindAll() ([]model.Lokasi_Dosen, error)
	FindById(id int) (model.Lokasi_Dosen, error)
	FindAllRelation() ([]model.Lokasi_Dosen, error)
	FindRelationByDosenId(dosenId int) ([]model.Lokasi_Dosen, error)
	Create(lokasiDosenRequest request.CreateLokasiDosenRequest) (model.Lokasi_Dosen, error)
	Update(id int, lokasiDosenRequest request.UpdateLokasiDosenRequest) (model.Lokasi_Dosen, error)
	Delete(id int) (model.Lokasi_Dosen, error)
}

type lokasiDosenService struct {
	lokasiDosenRepository repository.LokasiDosenRepository
}

func NewLokasiDosenService(lokasiDosenRepository repository.LokasiDosenRepository) *lokasiDosenService {
	return &lokasiDosenService{lokasiDosenRepository}
}

func (s *lokasiDosenService) FindAll() ([]model.Lokasi_Dosen, error) {
	var lokasiDosens, err = s.lokasiDosenRepository.FindAll()

	return lokasiDosens, err
}

func (s *lokasiDosenService) FindById(id int) (model.Lokasi_Dosen, error) {
	var lokasiDosen, err = s.lokasiDosenRepository.FindById(id)

	return lokasiDosen, err
}

func (s *lokasiDosenService) FindAllRelation() ([]model.Lokasi_Dosen, error) {
	var lokasiDosens, err = s.lokasiDosenRepository.FindAllRelation()

	return lokasiDosens, err
}

func (s *lokasiDosenService) FindRelationByDosenId(dosenId int) ([]model.Lokasi_Dosen, error) {
	var lokasiDosens, err = s.lokasiDosenRepository.FindRelationByDosenId(dosenId)

	return lokasiDosens, err
}

func (s *lokasiDosenService) Create(lokasiDosenRequest request.CreateLokasiDosenRequest) (model.Lokasi_Dosen, error) {
	var lokasiDosen = model.Lokasi_Dosen{
		DosenId:     lokasiDosenRequest.DosenId,
		KelurahanId: lokasiDosenRequest.KelurahanId,
	}

	newLokasiDosen, err := s.lokasiDosenRepository.Create(lokasiDosen)

	return newLokasiDosen, err
}

func (s *lokasiDosenService) Update(id int, lokasiDosenRequest request.UpdateLokasiDosenRequest) (model.Lokasi_Dosen, error) {
	var lokasiDosen, err = s.lokasiDosenRepository.FindById(id)

	lokasiDosen.DosenId = lokasiDosenRequest.DosenId
	lokasiDosen.KelurahanId = lokasiDosenRequest.KelurahanId

	updatedLokasiDosen, err := s.lokasiDosenRepository.Update(lokasiDosen)

	return updatedLokasiDosen, err
}

func (s *lokasiDosenService) Delete(id int) (model.Lokasi_Dosen, error) {
	var lokasiDosen, err = s.lokasiDosenRepository.FindById(id)

	deletedLokasiDosen, err := s.lokasiDosenRepository.Delete(lokasiDosen)

	return deletedLokasiDosen, err
}
