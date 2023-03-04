package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type FokusBelanjaService interface {
	FindAll() ([]model.FokusBelanja, error)
	FindById(id int) (model.FokusBelanja, error)
	FindBySubKegiatanId(subKegiatanId int) ([]model.FokusBelanja, error)
	Create(fokusBelanjaRequest request.CreateFokusBelanjaRequest) (model.FokusBelanja, error)
	Update(id int, fokusBelanjaRequest request.UpdateFokusBelanjaRequest) (model.FokusBelanja, error)
	Delete(id int) (model.FokusBelanja, error)
}

type fokusBelanjaService struct {
	fokusBelanjaRepository repository.FokusBelanjaRepository
}

func NewFokusBelanjaService(fokusBelanjaRepository repository.FokusBelanjaRepository) *fokusBelanjaService {
	return &fokusBelanjaService{fokusBelanjaRepository}
}

func (s *fokusBelanjaService) FindAll() ([]model.FokusBelanja, error) {
	var fokusBelanjas, err = s.fokusBelanjaRepository.FindAll()

	return fokusBelanjas, err
}

func (s *fokusBelanjaService) FindById(id int) (model.FokusBelanja, error) {
	var fokusBelanja, err = s.fokusBelanjaRepository.FindById(id)

	return fokusBelanja, err
}

func (s *fokusBelanjaService) FindBySubKegiatanId(subKegiatanId int) ([]model.FokusBelanja, error) {
	var fokusBelanjas, err = s.fokusBelanjaRepository.FindBySubKegiatanId(subKegiatanId)

	return fokusBelanjas, err
}

func (s *fokusBelanjaService) Create(fokusBelanjaRequest request.CreateFokusBelanjaRequest) (model.FokusBelanja, error) {
	var detailSubKegiatan = model.FokusBelanja{
		RencanaSubKegiatanId: fokusBelanjaRequest.RencanaSubKegiatanId,
		NamaFokusBelanja:     fokusBelanjaRequest.NamaFokusBelanja,
		Indikator:            fokusBelanjaRequest.Indikator,
		Target:               fokusBelanjaRequest.Target,
		Satuan:               fokusBelanjaRequest.Satuan,
		PaguFokusBelanja:     fokusBelanjaRequest.PaguFokusBelanja,
		Keterangan:           fokusBelanjaRequest.Keterangan,
	}

	var newFokusBelanja, err = s.fokusBelanjaRepository.Create(detailSubKegiatan)

	return newFokusBelanja, err
}

func (s *fokusBelanjaService) Update(id int, fokusBelanjaRequest request.UpdateFokusBelanjaRequest) (model.FokusBelanja, error) {
	var fokusBelanja, err = s.fokusBelanjaRepository.FindById(id)

	fokusBelanja.RencanaSubKegiatanId = fokusBelanjaRequest.RencanaSubKegiatanId
	fokusBelanja.NamaFokusBelanja = fokusBelanjaRequest.NamaFokusBelanja
	fokusBelanja.Indikator = fokusBelanjaRequest.Indikator
	fokusBelanja.Target = fokusBelanjaRequest.Target
	fokusBelanja.Satuan = fokusBelanjaRequest.Satuan
	fokusBelanja.PaguFokusBelanja = fokusBelanjaRequest.PaguFokusBelanja
	fokusBelanja.Keterangan = fokusBelanjaRequest.Keterangan

	updatedFokusBelanja, err := s.fokusBelanjaRepository.Update(fokusBelanja)

	return updatedFokusBelanja, err
}

func (s *fokusBelanjaService) Delete(id int) (model.FokusBelanja, error) {
	var fokusBelanja, err = s.fokusBelanjaRepository.FindById(id)

	deletedFokusBelanja, err := s.fokusBelanjaRepository.Delete(fokusBelanja)

	return deletedFokusBelanja, err
}
