package request

type CreateKeluargaVerifikasiRequest struct {
	Id                     int    `json:"id" binding:"required"`
	IdKeluarga             string `json:"idKeluarga" binding:"required"`
	ProvinsiId             string `json:"provinsiId" binding:"required"`
	KabupatenKotaId        string `json:"kabupatenKotaId" binding:"required"`
	KecamatanId            string `json:"kecamatanId" binding:"required"`
	KelurahanId            string `json:"kelurahanId" binding:"required"`
	DesilKesejahteraan     string `json:"desilKesejahteraan" binding:"required"`
	Alamat                 string `json:"alamat" binding:"required"`
	KepalaKeluarga         string `json:"kepalaKeluarga" binding:"required"`
	Nik                    string `json:"nik" binding:"required"`
	PadanDukcapil          string `json:"padanDukcapil" binding:"required"`
	JenisKelamin           string `json:"jenisKelamin" binding:"required"`
	TanggalLahir           string `json:"tanggalLahir" binding:"required"`
	Pekerjaan              string `json:"pekerjaan" binding:"required"`
	Pendidikan             string `json:"pendidikan" binding:"required"`
	KepemilikanRumah       string `json:"kepemilikanRumah" binding:"required"`
	Simpanan               string `json:"simpanan" binding:"required"`
	JenisAtap              string `json:"jenisAtap" binding:"required"`
	JenisDinding           string `json:"jenisDinding" binding:"required"`
	JenisLantai            string `json:"jenisLantai" binding:"required"`
	SumberPenerangan       string `json:"sumberPenerangan" binding:"required"`
	BahanBakarMemasak      string `json:"bahanBakarMemasak" binding:"required"`
	SumberAirMinum         string `json:"sumberAirMinum" binding:"required"`
	FasilitasBuangAirBesar string `json:"fasilitasBuangAirBesar" binding:"required"`
	PenerimaBPNT           string `json:"penerimaBPNT" binding:"required"`
	PenerimaBPUM           string `json:"penerimaBPUM" binding:"required"`
	PenerimaBST            string `json:"penerimaBST" binding:"required"`
	PenerimaPKH            string `json:"penerimaPKH" binding:"required"`
	PenerimaSembako        string `json:"penerimaSembako" binding:"required"`
	PenerimaLainnya        string `json:"penerimaLainnya"`
	StatusResponden        string `json:"statusResponden" binding:"required"`
	UserId                 int    `json:"userId" binding:"required"`
	MahasiswaId            int    `json:"mahasiswaid" binding:"required"`
}

type UpdateKeluargaVerifikasiRequest struct {
	Id                     int    `json:"id"`
	IdKeluarga             string `json:"idKeluarga"`
	ProvinsiId             string `json:"provinsiId"`
	KabupatenKotaId        string `json:"kabupatenKotaId"`
	KecamatanId            string `json:"kecamatanId"`
	KelurahanId            string `json:"kelurahanId"`
	DesilKesejahteraan     string `json:"desilKesejahteraan"`
	Alamat                 string `json:"alamat"`
	KepalaKeluarga         string `json:"kepalaKeluarga"`
	Nik                    string `json:"nik"`
	PadanDukcapil          string `json:"padanDukcapil"`
	JenisKelamin           string `json:"jenisKelamin"`
	TanggalLahir           string `json:"tanggalLahir"`
	Pekerjaan              string `json:"pekerjaan"`
	Pendidikan             string `json:"pendidikan"`
	KepemilikanRumah       string `json:"kepemilikanRumah"`
	Simpanan               string `json:"simpanan"`
	JenisAtap              string `json:"jenisAtap"`
	JenisDinding           string `json:"jenisDinding"`
	JenisLantai            string `json:"jenisLantai"`
	SumberPenerangan       string `json:"sumberPenerangan"`
	BahanBakarMemasak      string `json:"bahanBakarMemasak"`
	SumberAirMinum         string `json:"sumberAirMinum"`
	FasilitasBuangAirBesar string `json:"fasilitasBuangAirBesar"`
	PenerimaBPNT           string `json:"penerimaBPNT"`
	PenerimaBPUM           string `json:"penerimaBPUM"`
	PenerimaBST            string `json:"penerimaBST"`
	PenerimaPKH            string `json:"penerimaPKH"`
	PenerimaSembako        string `json:"penerimaSembako"`
	PenerimaLainnya        string `json:"penerimaLainnya"`
	StatusResponden        string `json:"statusResponded"`
	UserId                 int    `json:"userId"`
	MahasiswaId            int    `json:"mahasiswaId"`
}
