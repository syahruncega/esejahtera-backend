package responses

import "time"

type BidangUrusanResponse struct {
	Id               int       `json:"id"`
	NamaBidangUrusan string    `json:"namaBidangUrusan"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type InstansiResponse struct {
	Id             int       `json:"id"`
	NamaInstansi   string    `json:"namaInstansi"`
	BidangUrusanId int       `json:"bidangUrusanId"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type InstansiWithBidangUrusanResponse struct {
	Id             int                  `json:"id"`
	NamaInstansi   string               `json:"namaInstansi"`
	BidangUrusanId int                  `json:"bidangUrusanId"`
	CreatedAt      time.Time            `json:"createdAt"`
	UpdatedAt      time.Time            `json:"updatedAt"`
	BidangUrusan   BidangUrusanResponse `json:"bidangUrusan"`
}

type ProgramResponse struct {
	Id                      int       `json:"id"`
	Sasaran                 string    `json:"sasaran"`
	IndikatorSasaran        string    `json:"indikatorSasaran"`
	Kebijakan               string    `json:"kebijakan"`
	NamaProgram             string    `json:"namaProgram"`
	IndikatorKinerjaProgram string    `json:"indikatorKinerjaProgram"`
	PaguProgram             int       `json:"paguProgram"`
	InstansiId              int       `json:"instansiId"`
	CreatedAt               time.Time `json:"createdAt"`
	UpdatedAt               time.Time `json:"updatedAt"`
}

type ProgramWithInstansiResponse struct {
	Id                      int              `json:"id"`
	Sasaran                 string           `json:"sasaran"`
	IndikatorSasaran        string           `json:"indikatorSasaran"`
	Kebijakan               string           `json:"kebijakan"`
	NamaProgram             string           `json:"namaProgram"`
	IndikatorKinerjaProgram string           `json:"indikatorKinerjaProgram"`
	PaguProgram             int              `json:"paguProgram"`
	InstansiId              int              `json:"instansiId"`
	CreatedAt               time.Time        `json:"createdAt"`
	UpdatedAt               time.Time        `json:"updatedAt"`
	Instansi                InstansiResponse `json:"instansi"`
}

type KegiatanResponse struct {
	Id                       int       `json:"id"`
	NamaKegiatan             string    `json:"namaKegiatan"`
	IndikatorKinerjaKegiatan string    `json:"indikatorKinerjaKegiatan"`
	PaguKegiatan             int       `json:"paguKegiatan"`
	ProgramId                int       `json:"programId"`
	CreatedAt                time.Time `json:"createdAt"`
	UpdatedAt                time.Time `json:"updatedAt"`
}

type KegiatanWithProgramResponse struct {
	Id                       int             `json:"id"`
	NamaKegiatan             string          `json:"namaKegiatan"`
	IndikatorKinerjaKegiatan string          `json:"indikatorKinerjaKegiatan"`
	PaguKegiatan             int             `json:"paguKegiatan"`
	ProgramId                int             `json:"programId"`
	CreatedAt                time.Time       `json:"createdAt"`
	UpdatedAt                time.Time       `json:"updatedAt"`
	Program                  ProgramResponse `json:"program"`
}

type Sub_KegiatanResponse struct {
	Id                          int       `json:"id"`
	NamaSubKegiatan             string    `json:"namaSubKegiatan"`
	IndikatorKinerjaSubKegiatan string    `json:"indikatorKinerjaSubKegiatan"`
	PaguSubKegiatan             int       `json:"paguSubKegiatan"`
	KegiatanId                  int       `json:"kegiatanId"`
	CreatedAt                   time.Time `json:"createdAt"`
	UpdatedAt                   time.Time `json:"updatedAt"`
}

type Sub_KegiatanWithKegiatanResponse struct {
	Id                          int              `json:"id"`
	NamaSubKegiatan             string           `json:"namaSubKegiatan"`
	IndikatorKinerjaSubKegiatan string           `json:"indikatorKinerjaSubKegiatan"`
	PaguSubKegiatan             int              `json:"paguSubKegiatan"`
	KegiatanId                  int              `json:"kegiatanId"`
	CreatedAt                   time.Time        `json:"createdAt"`
	UpdatedAt                   time.Time        `json:"updatedAt"`
	Kegiatan                    KegiatanResponse `json:"kegiatan"`
}

type Detail_Sub_KegiatanResponse struct {
	Id               int       `json:"id"`
	FokusBelanja     string    `json:"fokusBelanja"`
	Indikator        string    `json:"indikator"`
	Target           float32   `json:"target"`
	Satuan           string    `json:"satuan"`
	PaguFokusBelanja int       `json:"paguFokusBelanja"`
	Keterangan       string    `json:"keterangan"`
	SubKegiatanId    int       `json:"subKegiatanId"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type Detail_Sub_KegiatanWithSub_KegiatanResponse struct {
	Id               int                  `json:"id"`
	FokusBelanja     string               `json:"fokusBelanja"`
	Indikator        string               `json:"indikator"`
	Target           float32              `json:"target"`
	Satuan           string               `json:"satuan"`
	PaguFokusBelanja int                  `json:"paguFokusBelanja"`
	Keterangan       string               `json:"keterangan"`
	SubKegiatanId    int                  `json:"subKegiatanId"`
	CreatedAt        time.Time            `json:"createdAt"`
	UpdatedAt        time.Time            `json:"updatedAt"`
	SubKegiatan      Sub_KegiatanResponse `json:"subKegiatan"`
}

type Detail_LokasiResponse struct {
	Id                  int       `json:"id"`
	KabupatenKotaId     string    `json:"kabupatenKotaId"`
	KecamatanId         string    `json:"kecamatanId"`
	KelurahanId         string    `json:"kelurahanId"`
	DetailSubKegiatanId int       `json:"detailSubKegiatanId"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
}

type Detail_LokasiWithRelationResponse struct {
	Id                  int                         `json:"id"`
	KabupatenKotaId     string                      `json:"kabupatenKotaId"`
	KecamatanId         string                      `json:"kecamatanId"`
	KelurahanId         string                      `json:"kelurahanId"`
	DetailSubKegiatanId int                         `json:"detailSubKegiatanId"`
	CreatedAt           time.Time                   `json:"createdAt"`
	UpdatedAt           time.Time                   `json:"updatedAt"`
	DetailSubKegiatan   Detail_Sub_KegiatanResponse `json:"detailSubKegiatan"`
	KabupatenKota       Kabupaten_KotaResponse      `json:"kabupatenKota"`
	Kecamatan           KecamatanResponse           `json:"kecamatan"`
	Kelurahan           KelurahanResponse           `json:"kelurahan"`
}

type ProvinsiResponse struct {
	Id   string `json:"id"`
	Nama string `json:"nama"`
}

type Kabupaten_KotaResponse struct {
	Id         string `json:"id"`
	ProvinsiId string `json:"provinsiId"`
	Nama       string `json:"nama"`
}

type KecamatanResponse struct {
	Id              string `json:"id"`
	KabupatenKotaId string `json:"kabupatenKotaId"`
	Nama            string `json:"nama"`
}

type KelurahanResponse struct {
	Id          string `json:"id"`
	KecamatanId string `json:"kecamatanId"`
	Nama        string `json:"nama"`
}

type KeluargaResponse struct {
	Id                     int    `json:"id"`
	IdKeluarga             string `json:"idKeluarga"`
	ProvinsiId             string `json:"provinsiId"`
	Provinsi               ProvinsiResponse
	KabupatenKotaId        string `json:"kabupatenKotaId"`
	KabupatenKota          Kabupaten_KotaResponse
	KecamatanId            string `json:"kecamatanId"`
	Kecamatan              KecamatanResponse
	KelurahanId            string `json:"kelurahanId"`
	Kelurahan              KelurahanResponse
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
}
