package request

type CreateSubKegiatanRequest struct {
	Id              string `json:"id" binding:"required"`
	NamaSubKegiatan string `json:"namaSubKegiatan" binding:"required"`
}

type UpdateSubKegiatanRequest struct {
	Id              string `json:"id"`
	NamaSubKegiatan string `json:"namaSubKegiatan"`
}
