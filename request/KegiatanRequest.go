package request

type CreateKegiatanRequest struct {
	KegiatanId   string `json:"kegiatanId" binding:"required"`
	NamaKegiatan string `json:"namaKegiatan" binding:"required"`
}

type UpdateKegiatanRequest struct {
	KegiatanId   string `json:"kegiatanId"`
	NamaKegiatan string `json:"namaKegiatan"`
}
