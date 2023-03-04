package request

type CreateKegiatanOnSubKegiatanRequest struct {
	KegiatanId    int `json:"kegiatanId" binding:"required"`
	SubKegiatanId int `json:"subKegiatanId" binding:"required"`
}

type UpdateKegiatanOnSubKegiatanRequest struct {
	KegiatanId    int `json:"kegiatanId"`
	SubKegiatanId int `json:"subKegiatanId"`
}
