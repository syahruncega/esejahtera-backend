package request

type CreateDetailSubKegiatanRequest struct {
	KegiatanId      int   `json:"kegiatanId" binding:"required"`
	SubKegiatanId   int   `json:"subKegiatanId" binding:"required"`
	PaguSubKegiatan int64 `json:"paguSubKegiatan" binding:"required"`
}

type UpdateDetailSubKegiatanRequest struct {
	KegiatanId      int   `json:"kegiatanId"`
	SubKegiatanId   int   `json:"subKegiatanId"`
	PaguSubKegiatan int64 `json:"paguSubKegiatan"`
}
