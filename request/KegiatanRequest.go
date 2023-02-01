package request

type CreateKegiatanRequest struct {
	Id           string `json:"id" binding:"required"`
	NamaKegiatan string `json:"namaKegiatan" binding:"required"`
}

type UpdateKegiatanRequest struct {
	Id           string `json:"id"`
	NamaKegiatan string `json:"namaKegiatan"`
}
