package service

import (
	"errors"
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type FokusBelanjaService interface {
	FindAll() ([]model.FokusBelanja, error)
	FindById(id int) (model.FokusBelanja, error)
	FindByRencanaSubKegiatanId(rencanaSubKegiatanId int) ([]model.FokusBelanja, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.FokusBelanja, error)
	SumAllPaguFokusBelanja(tahun, tipe string) int64
	SumPaguFokusBelanjaByInstansi(tahun, tipe string, instansis []model.Instansi) []int64
	CountJumlahFokusBelanja(tahun string) (int64, error)
	CountJumlahFokusBelanjaByInstansi(tahun string, tipe string, instansis []model.Instansi) []int64
	Create(fokusBelanjaRequest request.CreateFokusBelanjaRequest) (model.FokusBelanja, error)
	Update(id int, fokusBelanjaRequest request.UpdateFokusBelanjaRequest) (model.FokusBelanja, error)
	Delete(id int) (model.FokusBelanja, error)
}

type fokusBelanjaService struct {
	fokusBelanjaRepository       repository.FokusBelanjaRepository
	rencanaSubKegiatanRepository repository.RencanaSubKegiatanRepository
}

func NewFokusBelanjaService(fokusBelanjaRepository repository.FokusBelanjaRepository, rencanaSubKegiatanRepository repository.RencanaSubKegiatanRepository) *fokusBelanjaService {
	return &fokusBelanjaService{fokusBelanjaRepository, rencanaSubKegiatanRepository}
}

func (s *fokusBelanjaService) FindAll() ([]model.FokusBelanja, error) {
	var fokusBelanjas, err = s.fokusBelanjaRepository.FindAll()

	return fokusBelanjas, err
}

func (s *fokusBelanjaService) FindById(id int) (model.FokusBelanja, error) {
	var fokusBelanja, err = s.fokusBelanjaRepository.FindById(id)

	return fokusBelanja, err
}

func (s *fokusBelanjaService) FindByRencanaSubKegiatanId(rencanaSubKegiatanId int) ([]model.FokusBelanja, error) {
	var fokusBelanjas, err = s.fokusBelanjaRepository.FindByRencanaSubKegiatanId(rencanaSubKegiatanId)

	return fokusBelanjas, err
}

func (s *fokusBelanjaService) FindBySearch(whereClause map[string]interface{}) ([]model.FokusBelanja, error) {
	var fokusBelanjas, err = s.fokusBelanjaRepository.FindBySearch(whereClause)

	return fokusBelanjas, err
}

func (s *fokusBelanjaService) SumAllPaguFokusBelanja(tahun, tipe string) int64 {
	var jumlahPaguFokusBelanja = s.fokusBelanjaRepository.SumAllPaguFokusBelanja(tahun, tipe)

	return jumlahPaguFokusBelanja
}

func (s *fokusBelanjaService) SumPaguFokusBelanjaByInstansi(tahun, tipe string, instansis []model.Instansi) []int64 {
	var jumlahPaguFokusBelanja = s.fokusBelanjaRepository.SumPaguFokusBelanjaByInstansi(tahun, tipe, instansis)

	return jumlahPaguFokusBelanja
}

func (s *fokusBelanjaService) CountJumlahFokusBelanja(tahun string) (int64, error) {
	var jumlahFokusBelanja, err = s.fokusBelanjaRepository.CountJumlahFokusBelanja(tahun)

	return jumlahFokusBelanja, err
}

func (s *fokusBelanjaService) CountJumlahFokusBelanjaByInstansi(tahun string, tipe string, instansis []model.Instansi) []int64 {
	var jumlahFokusBelanja = s.fokusBelanjaRepository.CountJumlahFokusBelanjaByInstansi(tahun, tipe, instansis)

	return jumlahFokusBelanja
}

func (s *fokusBelanjaService) Create(fokusBelanjaRequest request.CreateFokusBelanjaRequest) (model.FokusBelanja, error) {
	var fokusBelanja = model.FokusBelanja{
		RencanaSubKegiatanId: fokusBelanjaRequest.RencanaSubKegiatanId,
		NamaFokusBelanja:     fokusBelanjaRequest.NamaFokusBelanja,
		Indikator:            fokusBelanjaRequest.Indikator,
		Target:               fokusBelanjaRequest.Target,
		Satuan:               fokusBelanjaRequest.Satuan,
		PaguFokusBelanja:     fokusBelanjaRequest.PaguFokusBelanja,
		Keterangan:           fokusBelanjaRequest.Keterangan,
		Tahun:                fokusBelanjaRequest.Tahun,
	}

	var rencanaSubKegiatan, _ = s.rencanaSubKegiatanRepository.FindById(fokusBelanja.RencanaSubKegiatanId)

	var totalPaguFokusBelanja, _ = s.fokusBelanjaRepository.SumPaguFokusBelanja(fokusBelanja.RencanaSubKegiatanId)
	totalPaguFokusBelanja = totalPaguFokusBelanja + int64(fokusBelanja.PaguFokusBelanja)

	if rencanaSubKegiatan.PaguSubKegiatan >= totalPaguFokusBelanja {
		var newFokusBelanja, err = s.fokusBelanjaRepository.Create(fokusBelanja)
		return newFokusBelanja, err
	} else {
		return fokusBelanja, errors.New("pagu fokus belanja melebihi pagu sub kegiatan")
	}

}

func (s *fokusBelanjaService) Update(id int, fokusBelanjaRequest request.UpdateFokusBelanjaRequest) (model.FokusBelanja, error) {
	var fokusBelanja, _ = s.fokusBelanjaRepository.FindById(id)

	var currentPagu = fokusBelanja.PaguFokusBelanja

	fokusBelanja.RencanaSubKegiatanId = fokusBelanjaRequest.RencanaSubKegiatanId
	fokusBelanja.NamaFokusBelanja = fokusBelanjaRequest.NamaFokusBelanja
	fokusBelanja.Indikator = fokusBelanjaRequest.Indikator
	fokusBelanja.Target = fokusBelanjaRequest.Target
	fokusBelanja.Satuan = fokusBelanjaRequest.Satuan
	fokusBelanja.PaguFokusBelanja = fokusBelanjaRequest.PaguFokusBelanja
	fokusBelanja.Keterangan = fokusBelanjaRequest.Keterangan
	fokusBelanja.Tahun = fokusBelanjaRequest.Tahun

	var rencanaSubKegiatan, _ = s.rencanaSubKegiatanRepository.FindById(fokusBelanja.RencanaSubKegiatanId)

	var totalPaguFokusBelanja, _ = s.fokusBelanjaRepository.SumPaguFokusBelanja(fokusBelanja.RencanaSubKegiatanId)
	totalPaguFokusBelanja = totalPaguFokusBelanja - currentPagu + fokusBelanjaRequest.PaguFokusBelanja

	if rencanaSubKegiatan.PaguSubKegiatan >= totalPaguFokusBelanja {
		updatedFokusBelanja, err := s.fokusBelanjaRepository.Update(fokusBelanja)
		return updatedFokusBelanja, err
	} else {
		return fokusBelanja, errors.New("pagu fokus belanja melebihi pagu sub-kegiatannya")
	}
}

func (s *fokusBelanjaService) Delete(id int) (model.FokusBelanja, error) {
	var fokusBelanja, err = s.fokusBelanjaRepository.FindById(id)

	deletedFokusBelanja, err := s.fokusBelanjaRepository.Delete(fokusBelanja)

	return deletedFokusBelanja, err
}
