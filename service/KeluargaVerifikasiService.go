package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type KeluargaVerifikasiService interface {
	FindAll(kabupatenKotaId string) ([]model.KeluargaVerifikasi, error)
	FindById(kabupatenKotaId string, id int) (model.KeluargaVerifikasi, error)
	Create(keluargaVerifikasiRequest request.CreateKeluargaVerifikasiRequest) (model.KeluargaVerifikasi, error)
	Update(kabupatenKotaId string, id int, keluargaVerifikasiRequest request.UpdateKeluargaVerifikasiRequest) (model.KeluargaVerifikasi, error)
	Delete(kabupatenKotaId string, id int) (model.KeluargaVerifikasi, error)
}

type keluargaVerifikasiService struct {
	keluargaVerifikasiRepository repository.KeluargaVerifikasiRepository
}

func NewKeluargaVerifikasiService(keluargaVerifikasiRepository repository.KeluargaVerifikasiRepository) *keluargaVerifikasiService {
	return &keluargaVerifikasiService{keluargaVerifikasiRepository}
}

func (s *keluargaVerifikasiService) FindAll(kabupatenKotaId string) ([]model.KeluargaVerifikasi, error) {
	var keluargaVerifikasis, err = s.keluargaVerifikasiRepository.FindAll(kabupatenKotaId)

	return keluargaVerifikasis, err
}

func (s *keluargaVerifikasiService) FindById(kabupatenKotaId string, id int) (model.KeluargaVerifikasi, error) {
	var keluargaVerifikasi, err = s.keluargaVerifikasiRepository.FindById(kabupatenKotaId, id)

	return keluargaVerifikasi, err
}

func (s *keluargaVerifikasiService) Create(keluargaVerifikasiRequest request.CreateKeluargaVerifikasiRequest) (model.KeluargaVerifikasi, error) {
	var keluargaVerifikasi = model.KeluargaVerifikasi{
		Id:                     keluargaVerifikasiRequest.Id,
		IdKeluarga:             keluargaVerifikasiRequest.IdKeluarga,
		ProvinsiId:             keluargaVerifikasiRequest.ProvinsiId,
		KabupatenKotaId:        keluargaVerifikasiRequest.KabupatenKotaId,
		KecamatanId:            keluargaVerifikasiRequest.KecamatanId,
		KelurahanId:            keluargaVerifikasiRequest.KelurahanId,
		DesilKesejahteraan:     keluargaVerifikasiRequest.DesilKesejahteraan,
		Alamat:                 keluargaVerifikasiRequest.Alamat,
		KepalaKeluarga:         keluargaVerifikasiRequest.KepalaKeluarga,
		Nik:                    keluargaVerifikasiRequest.Nik,
		PadanDukcapil:          keluargaVerifikasiRequest.PadanDukcapil,
		JenisKelamin:           keluargaVerifikasiRequest.JenisKelamin,
		TanggalLahir:           keluargaVerifikasiRequest.TanggalLahir,
		Pekerjaan:              keluargaVerifikasiRequest.Pekerjaan,
		Pendidikan:             keluargaVerifikasiRequest.Pendidikan,
		KepemilikanRumah:       keluargaVerifikasiRequest.KepemilikanRumah,
		Simpanan:               keluargaVerifikasiRequest.Simpanan,
		JenisAtap:              keluargaVerifikasiRequest.JenisAtap,
		JenisDinding:           keluargaVerifikasiRequest.JenisDinding,
		JenisLantai:            keluargaVerifikasiRequest.JenisLantai,
		SumberPenerangan:       keluargaVerifikasiRequest.SumberPenerangan,
		BahanBakarMemasak:      keluargaVerifikasiRequest.BahanBakarMemasak,
		SumberAirMinum:         keluargaVerifikasiRequest.SumberAirMinum,
		FasilitasBuangAirBesar: keluargaVerifikasiRequest.FasilitasBuangAirBesar,
		PenerimaBPNT:           keluargaVerifikasiRequest.PenerimaBPNT,
		PenerimaBPUM:           keluargaVerifikasiRequest.PenerimaBPUM,
		PenerimaBST:            keluargaVerifikasiRequest.PenerimaBST,
		PenerimaPKH:            keluargaVerifikasiRequest.PenerimaPKH,
		PenerimaSembako:        keluargaVerifikasiRequest.PenerimaSembako,
		PenerimaLainnya:        keluargaVerifikasiRequest.PenerimaLainnya,
		StatusResponden:        keluargaVerifikasiRequest.StatusResponden,
		UserId:                 keluargaVerifikasiRequest.UserId,
		MahasiswaId:            keluargaVerifikasiRequest.MahasiswaId,
	}

	newKeluargaVerifikasi, err := s.keluargaVerifikasiRepository.Create(keluargaVerifikasi)

	return newKeluargaVerifikasi, err
}

func (s *keluargaVerifikasiService) Update(kabupatenKotaId string, id int, keluargaVerifikasiRequest request.UpdateKeluargaVerifikasiRequest) (model.KeluargaVerifikasi, error) {
	var keluargaVerifikasi, err = s.keluargaVerifikasiRepository.FindById(kabupatenKotaId, id)

	keluargaVerifikasi.IdKeluarga = keluargaVerifikasiRequest.IdKeluarga
	keluargaVerifikasi.ProvinsiId = keluargaVerifikasiRequest.ProvinsiId
	keluargaVerifikasi.KabupatenKotaId = keluargaVerifikasiRequest.KabupatenKotaId
	keluargaVerifikasi.KecamatanId = keluargaVerifikasiRequest.KecamatanId
	keluargaVerifikasi.KelurahanId = keluargaVerifikasiRequest.KelurahanId
	keluargaVerifikasi.DesilKesejahteraan = keluargaVerifikasiRequest.DesilKesejahteraan
	keluargaVerifikasi.Alamat = keluargaVerifikasiRequest.Alamat
	keluargaVerifikasi.KepalaKeluarga = keluargaVerifikasiRequest.KepalaKeluarga
	keluargaVerifikasi.Nik = keluargaVerifikasiRequest.Nik
	keluargaVerifikasi.PadanDukcapil = keluargaVerifikasiRequest.PadanDukcapil
	keluargaVerifikasi.JenisKelamin = keluargaVerifikasiRequest.JenisKelamin
	keluargaVerifikasi.TanggalLahir = keluargaVerifikasiRequest.TanggalLahir
	keluargaVerifikasi.Pekerjaan = keluargaVerifikasiRequest.Pekerjaan
	keluargaVerifikasi.Pendidikan = keluargaVerifikasiRequest.Pendidikan
	keluargaVerifikasi.KepemilikanRumah = keluargaVerifikasiRequest.KepemilikanRumah
	keluargaVerifikasi.Simpanan = keluargaVerifikasiRequest.Simpanan
	keluargaVerifikasi.JenisAtap = keluargaVerifikasiRequest.JenisAtap
	keluargaVerifikasi.JenisDinding = keluargaVerifikasiRequest.JenisDinding
	keluargaVerifikasi.JenisLantai = keluargaVerifikasiRequest.JenisLantai
	keluargaVerifikasi.SumberPenerangan = keluargaVerifikasiRequest.SumberPenerangan
	keluargaVerifikasi.BahanBakarMemasak = keluargaVerifikasiRequest.BahanBakarMemasak
	keluargaVerifikasi.SumberAirMinum = keluargaVerifikasiRequest.SumberAirMinum
	keluargaVerifikasi.FasilitasBuangAirBesar = keluargaVerifikasiRequest.FasilitasBuangAirBesar
	keluargaVerifikasi.PenerimaBPNT = keluargaVerifikasiRequest.PenerimaBPNT
	keluargaVerifikasi.PenerimaBPUM = keluargaVerifikasiRequest.PenerimaBPUM
	keluargaVerifikasi.PenerimaBST = keluargaVerifikasiRequest.PenerimaBST
	keluargaVerifikasi.PenerimaPKH = keluargaVerifikasiRequest.PenerimaPKH
	keluargaVerifikasi.PenerimaSembako = keluargaVerifikasiRequest.PenerimaSembako
	keluargaVerifikasi.PenerimaLainnya = keluargaVerifikasiRequest.PenerimaLainnya
	keluargaVerifikasi.StatusResponden = keluargaVerifikasiRequest.StatusResponden
	keluargaVerifikasi.UserId = keluargaVerifikasiRequest.UserId
	keluargaVerifikasi.MahasiswaId = keluargaVerifikasiRequest.MahasiswaId

	updatedKeluargaVerifikasi, err := s.keluargaVerifikasiRepository.Update(keluargaVerifikasi)

	return updatedKeluargaVerifikasi, err
}

func (s *keluargaVerifikasiService) Delete(kabupatenKotaId string, id int) (model.KeluargaVerifikasi, error) {
	var keluargaVerifikasi, err = s.keluargaVerifikasiRepository.FindById(kabupatenKotaId, id)

	deletedKeluargaVerifikasi, err := s.keluargaVerifikasiRepository.Delete(keluargaVerifikasi)

	return deletedKeluargaVerifikasi, err
}
