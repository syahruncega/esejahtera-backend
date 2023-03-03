package request

type CreateFokusBelanjaRequest struct {
	SubKegiatanId    int     `json:"subKegiatanId" binding:"required"`
	NamaFokusBelanja string  `json:"namaFokusBelanja" binding:"required"`
	Indikator        string  `json:"indikator" binding:"required"`
	Target           float32 `json:"target" binding:"required"`
	Satuan           string  `json:"satuan" binding:"required"`
	PaguFokusBelanja int     `json:"paguFokusBelanja" binding:"required"`
	Keterangan       string  `json:"keterangan" binding:"required"`
}

type UpdateFokusBelanjaRequest struct {
	SubKegiatanId    int     `json:"subKegiatanId"`
	NamaFokusBelanja string  `json:"namaFokusBelanja"`
	Indikator        string  `json:"indikator"`
	Target           float32 `json:"target"`
	Satuan           string  `json:"satuan"`
	PaguFokusBelanja int     `json:"paguFokusBelanja"`
	Keterangan       string  `json:"keterangan"`
}
