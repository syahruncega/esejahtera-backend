package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type DetailLokasiService interface {
	FindAll() ([]model.Detail_Lokasi, error)
	FindById(id int) (model.Detail_Lokasi, error)
	Create(detailLokasiRequest request.CreateDetail_LokasiRequest) (model.Detail_Lokasi, error)
	Update(id int, detailLokasiRequest request.UpdateDetail_LokasiRequest) (model.Detail_Lokasi, error)
	Delete(id int) (model.Detail_Lokasi, error)
}

type detailLokasiService struct {
	detailLokasiRepository repository.DetailLokasiRepository
}

func NewDetailLokasiService(detailLokasiRepository repository.DetailLokasiRepository) *detailLokasiService {
	return &detailLokasiService{detailLokasiRepository}
}

func (s *detailLokasiService) FindAll() ([]model.Detail_Lokasi, error) {
	var detailLokasis, err = s.detailLokasiRepository.FindAll()

	return detailLokasis, err
}

func (s *detailLokasiService) FindById(id int) (model.Detail_Lokasi, error) {
	var detailLokasi, err = s.detailLokasiRepository.FindById(id)

	return detailLokasi, err
}

func (s *detailLokasiService) Create(detailLokasiRequest request.CreateDetail_LokasiRequest) (model.Detail_Lokasi, error) {
	var detailLokasi = model.Detail_Lokasi{
		FokusBelanjaId:  detailLokasiRequest.FokusBelanjaId,
		KabupatenKotaId: detailLokasiRequest.KabupatenKotaId,
		KecamatanId:     detailLokasiRequest.KecamatanId,
		KelurahanId:     detailLokasiRequest.KelurahanId,
	}

	var newDetailLokasi, err = s.detailLokasiRepository.Create(detailLokasi)

	return newDetailLokasi, err
}

func (s *detailLokasiService) Update(id int, detailLokasiRequest request.UpdateDetail_LokasiRequest) (model.Detail_Lokasi, error) {
	var detailLokasi, err = s.detailLokasiRepository.FindById(id)

	detailLokasi.KabupatenKotaId = detailLokasiRequest.KabupatenKotaId
	detailLokasi.KecamatanId = detailLokasiRequest.KecamatanId
	detailLokasi.KelurahanId = detailLokasiRequest.KelurahanId
	detailLokasi.FokusBelanjaId = detailLokasiRequest.FokusBelanjaId

	updatedDetailLokasi, err := s.detailLokasiRepository.Update(detailLokasi)

	return updatedDetailLokasi, err
}

func (s *detailLokasiService) Delete(id int) (model.Detail_Lokasi, error) {
	var detailLokasi, err = s.detailLokasiRepository.FindById(id)

	deletedDetailLokasi, err := s.detailLokasiRepository.Delete(detailLokasi)

	return deletedDetailLokasi, err
}
