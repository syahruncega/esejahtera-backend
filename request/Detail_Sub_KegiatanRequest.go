package request

type CreateDetail_Sub_KegiatanRequest struct {
	FokusBelanja     string  `json:"fokusBelanja" binding:"required"`
	Indikator        string  `json:"indikator" binding:"required"`
	Target           float32 `json:"target" binding:"required"`
	Satuan           string  `json:"satuan" binding:"required"`
	PaguFokusBelanja int     `json:"paguFokusBelanja" binding:"required"`
	Keterangan       string  `json:"keterangan" binding:"required"`
	SubKegiatanId    int     `json:"subKegiatanId" binding:"required"`
}

type UpdateDetail_Sub_KegiatanRequest struct {
	FokusBelanja     string  `json:"fokusBelanja"`
	Indikator        string  `json:"indikator"`
	Target           float32 `json:"target"`
	Satuan           string  `json:"satuan"`
	PaguFokusBelanja int     `json:"paguFokusBelanja"`
	Keterangan       string  `json:"keterangan"`
	SubKegiatanId    int     `json:"subKegiatanId"`
}
