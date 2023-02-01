package request

type CreateDetailSubKegiatanRequest struct {
	KegiatanId    string `json:"kegiatanId" binding:"required"`
	SubKegiatanId string `json:"subKegiatanId" binding:"required"`
}

type UpdateDetailSubKegiatanRequest struct {
	KegiatanId    string `json:"kegiatanId"`
	SubKegiatanId string `json:"subKegiatanId"`
}
