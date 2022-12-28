package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type DetailSubKegiatanService interface {
	FindAll() ([]model.Detail_Sub_Kegiatan, error)
	FindById(id int) (model.Detail_Sub_Kegiatan, error)
	FindAllRelation() ([]model.Detail_Sub_Kegiatan, error)
	Create(detailSubKegiatanRequest request.CreateDetail_Sub_KegiatanRequest) (model.Detail_Sub_Kegiatan, error)
	Update(id int, detailSubKegiatanRequest request.UpdateDetail_Sub_KegiatanRequest) (model.Detail_Sub_Kegiatan, error)
	Delete(id int) (model.Detail_Sub_Kegiatan, error)
}

type detailSubKegiatanService struct {
	detailSubKegiatanRepository repository.DetailSubKegiatanRepository
}

func NewDetailSubKegiatanService(detailSubKegiatanRepository repository.DetailSubKegiatanRepository) *detailSubKegiatanService {
	return &detailSubKegiatanService{detailSubKegiatanRepository}
}

func (s *detailSubKegiatanService) FindAll() ([]model.Detail_Sub_Kegiatan, error) {
	var detailSubKegiatans, err = s.detailSubKegiatanRepository.FindAll()

	return detailSubKegiatans, err
}

func (s *detailSubKegiatanService) FindById(id int) (model.Detail_Sub_Kegiatan, error) {
	var detailSubKegiatan, err = s.detailSubKegiatanRepository.FindById(id)

	return detailSubKegiatan, err
}

func (s *detailSubKegiatanService) FindAllRelation() ([]model.Detail_Sub_Kegiatan, error) {
	var detailSubKegiatanRelation, err = s.detailSubKegiatanRepository.FindAllRelation()

	return detailSubKegiatanRelation, err
}

func (s *detailSubKegiatanService) Create(detailSubKegiatanRequest request.CreateDetail_Sub_KegiatanRequest) (model.Detail_Sub_Kegiatan, error) {
	var detailSubKegiatan = model.Detail_Sub_Kegiatan{
		FokusBelanja:     detailSubKegiatanRequest.FokusBelanja,
		Indikator:        detailSubKegiatanRequest.Indikator,
		Target:           detailSubKegiatanRequest.Target,
		Satuan:           detailSubKegiatanRequest.Satuan,
		PaguFokusBelanja: detailSubKegiatanRequest.PaguFokusBelanja,
		Keterangan:       detailSubKegiatanRequest.Keterangan,
		SubKegiatanId:    detailSubKegiatanRequest.SubKegiatanId,
	}

	var newDetailSubKegiatan, err = s.detailSubKegiatanRepository.Create(detailSubKegiatan)

	return newDetailSubKegiatan, err
}

func (s *detailSubKegiatanService) Update(id int, detailSubKegiatanRequest request.UpdateDetail_Sub_KegiatanRequest) (model.Detail_Sub_Kegiatan, error) {
	var detailSubKegiatan, err = s.detailSubKegiatanRepository.FindById(id)

	detailSubKegiatan.FokusBelanja = detailSubKegiatanRequest.FokusBelanja
	detailSubKegiatan.Indikator = detailSubKegiatanRequest.Indikator
	detailSubKegiatan.Target = detailSubKegiatanRequest.Target
	detailSubKegiatan.Satuan = detailSubKegiatanRequest.Satuan
	detailSubKegiatan.PaguFokusBelanja = detailSubKegiatanRequest.PaguFokusBelanja
	detailSubKegiatan.Keterangan = detailSubKegiatanRequest.Keterangan
	detailSubKegiatan.SubKegiatanId = detailSubKegiatanRequest.SubKegiatanId

	updatedDetailSubKegiatan, err := s.detailSubKegiatanRepository.Update(detailSubKegiatan)

	return updatedDetailSubKegiatan, err
}

func (s *detailSubKegiatanService) Delete(id int) (model.Detail_Sub_Kegiatan, error) {
	var detailSubKegiatan, err = s.detailSubKegiatanRepository.FindById(id)

	deletedDetailSubKegiatan, err := s.detailSubKegiatanRepository.Delete(detailSubKegiatan)

	return deletedDetailSubKegiatan, err
}
