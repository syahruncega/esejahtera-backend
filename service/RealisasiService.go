package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
	"time"
)

type RealisasiService interface {
	FindAll() ([]model.Realisasi, error)
	FindById(id int) (model.Realisasi, error)
	FindByFokusBelanjaId(fokusBelanjaId int) ([]model.Realisasi, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.Realisasi, error)
	Create(realisasiRequest request.CreateRealisasiRequest) (model.Realisasi, error)
	Update(id int, realisasiRequest request.UpdateRealisasiRequest) (model.Realisasi, error)
	Delete(id int) (model.Realisasi, error)
}

type realisasiService struct {
	realisasiRepository repository.RealisasiRepository
}

func NewRealisasiService(realisasiRepository repository.RealisasiRepository) *realisasiService {
	return &realisasiService{realisasiRepository}
}

func convertIsoStringToMonth(isoString string) (string, error) {
	var isoDatetime = isoString

	var parsedTime, err = time.Parse(time.RFC3339, isoDatetime)

	var indonesianMonths = map[time.Month]string{
		time.January:   "Januari",
		time.February:  "Februari",
		time.March:     "Maret",
		time.April:     "April",
		time.May:       "Mei",
		time.June:      "June",
		time.July:      "Juli",
		time.August:    "August",
		time.September: "September",
		time.October:   "Oktober",
		time.November:  "November",
		time.December:  "December",
	}

	var bulan = indonesianMonths[parsedTime.Month()]

	return bulan, err
}

func (s *realisasiService) FindAll() ([]model.Realisasi, error) {
	var realisasis, err = s.realisasiRepository.FindAll()

	return realisasis, err
}

func (s *realisasiService) FindById(id int) (model.Realisasi, error) {
	var realisasi, err = s.realisasiRepository.FindById(id)

	return realisasi, err
}

func (s *realisasiService) FindBySearch(whereClause map[string]interface{}) ([]model.Realisasi, error) {
	var realisasis, err = s.realisasiRepository.FindBySearch(whereClause)

	return realisasis, err
}

func (s *realisasiService) FindByFokusBelanjaId(fokusBelanjaId int) ([]model.Realisasi, error) {
	var realisasis, err = s.realisasiRepository.FindByFokusBelanjaId(fokusBelanjaId)

	return realisasis, err
}

func (s *realisasiService) Create(realisasiRequest request.CreateRealisasiRequest) (model.Realisasi, error) {

	var bulan, err = convertIsoStringToMonth(realisasiRequest.Tanggal)

	var realisasi = model.Realisasi{
		FokusBelanjaId:  realisasiRequest.FokusBelanjaId,
		Tanggal:         realisasiRequest.Tanggal,
		NomorSp2d:       realisasiRequest.NomorSp2d,
		RealisasiPagu:   realisasiRequest.RealisasiPagu,
		RealisasiTarget: realisasiRequest.RealisasiTarget,
		Bulan:           bulan,
		Keterangan:      realisasiRequest.Keterangan,
	}

	newRealisasi, err := s.realisasiRepository.Create(realisasi)

	return newRealisasi, err
}

func (s *realisasiService) Update(id int, realisasiRequest request.UpdateRealisasiRequest) (model.Realisasi, error) {
	var bulan, err = convertIsoStringToMonth(realisasiRequest.Tanggal)

	realisasi, _ := s.realisasiRepository.FindById(id)

	realisasi.FokusBelanjaId = realisasiRequest.FokusBelanjaId
	realisasi.Tanggal = realisasiRequest.Tanggal
	realisasi.NomorSp2d = realisasiRequest.NomorSp2d
	realisasi.RealisasiPagu = realisasiRequest.RealisasiPagu
	realisasi.RealisasiTarget = realisasiRequest.RealisasiTarget
	realisasi.Bulan = bulan
	realisasi.Keterangan = realisasiRequest.Keterangan

	updatedRealisasi, err := s.realisasiRepository.Update(realisasi)

	return updatedRealisasi, err
}

func (s *realisasiService) Delete(id int) (model.Realisasi, error) {
	var realisasi, err = s.realisasiRepository.FindById(id)

	deletedRealisasi, err := s.realisasiRepository.Delete(realisasi)

	return deletedRealisasi, err
}
