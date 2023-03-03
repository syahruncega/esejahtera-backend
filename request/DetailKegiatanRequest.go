package request

type CreateDetailKegiatanRequest struct {
	KegiatanId   int    `json:"kegiatanId" binding:"required"`
	PaguKegiatan int64  `json:"paguKegiatan" binding:"required"`
	Tipe         string `json:"tipe" binding:"required"`
}

type UpdateDetailKegiatanRequest struct {
	KegiatanId   int    `json:"kegiatanId"`
	PaguKegiatan int64  `json:"paguKegiatan"`
	Tipe         string `json:"tipe" binding:"required"`
}
