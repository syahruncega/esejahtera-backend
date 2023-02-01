package request

type CreateFokusBelanjaRequest struct {
	SubKegiatanId    string  `json:"subKegiatanId" binding:"required"`
	FokusBelanja     string  `json:"fokusBelanja" binding:"required"`
	Indikator        string  `json:"indikator" binding:"required"`
	Target           float32 `json:"target"`
	Satuan           string  `json:"satuan" binding:"required"`
	PaguFokusBelanja int     `json:"paguFokusBelanja"`
	Keterangan       string  `json:"keterangan" binding:"required"`
}

type UpdateFokusBelanjaRequest struct {
	SubKegiatanId    string  `json:"subKegiatanId"`
	FokusBelanja     string  `json:"fokusBelanja"`
	Indikator        string  `json:"indikator"`
	Target           float32 `json:"target"`
	Satuan           string  `json:"satuan"`
	PaguFokusBelanja int     `json:"paguFokusBelanja"`
	Keterangan       string  `json:"keterangan"`
}
