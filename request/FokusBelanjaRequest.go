package request

type CreateFokusBelanjaRequest struct {
	RencanaSubKegiatanId int     `json:"rencanaSubKegiatanId" binding:"required"`
	NamaFokusBelanja     string  `json:"namaFokusBelanja" binding:"required"`
	Indikator            string  `json:"indikator" binding:"required"`
	Target               float32 `json:"target" binding:"required"`
	Satuan               string  `json:"satuan" binding:"required"`
	PaguFokusBelanja     int64   `json:"paguFokusBelanja" binding:"required"`
	Keterangan           string  `json:"keterangan" binding:"required"`
	Tahun                string  `json:"tahun" binding:"required"`
}

type UpdateFokusBelanjaRequest struct {
	RencanaSubKegiatanId int     `json:"rencanaSubKegiatanId"`
	SubKegiatanId        int     `json:"subKegiatanId"`
	NamaFokusBelanja     string  `json:"namaFokusBelanja"`
	Indikator            string  `json:"indikator"`
	Target               float32 `json:"target"`
	Satuan               string  `json:"satuan"`
	PaguFokusBelanja     int64   `json:"paguFokusBelanja"`
	Keterangan           string  `json:"keterangan"`
	Tahun                string  `json:"tahun"`
}
