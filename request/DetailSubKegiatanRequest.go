package request

type CreateDetailSubKegiatanRequest struct {
	KegiatanId      string `json:"kegiatanId" binding:"required"`
	SubKegiatanId   string `json:"subKegiatanId" binding:"required"`
	PaguSubKegiatan int64  `json:"paguSubKegiatan" binding:"required"`
}

type UpdateDetailSubKegiatanRequest struct {
	KegiatanId      string `json:"kegiatanId"`
	SubKegiatanId   string `json:"subKegiatanId"`
	PaguSubKegiatan int64  `json:"paguSubKegiatan"`
}
