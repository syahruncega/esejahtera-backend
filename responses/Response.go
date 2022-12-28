package responses

import "time"

type InstansiResponse struct {
	Id           int       `json:"id"`
	NamaInstansi string    `json:"namaInstansi"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type ProgramResponse struct {
	Id                      int       `json:"id"`
	NamaProgram             string    `json:"namaProgram"`
	IndikatorKinerjaProgram string    `json:"indikatorKinerjaProgram"`
	PaguProgram             int       `json:"paguProgram"`
	InstansiId              int       `json:"instansiId"`
	CreatedAt               time.Time `json:"createdAt"`
	UpdatedAt               time.Time `json:"updatedAt"`
}

type ProgramWithInstansiResponse struct {
	Id                      int              `json:"id"`
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
