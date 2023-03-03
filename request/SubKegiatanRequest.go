package request

type CreateSubKegiatanRequest struct {
	KegiatanId      int    `json:"kegiatanId" binding:"required"`
	Tahun           string `json:"tahun" binding:"required"`
	KodeSubKegiatan string `json:"kodeSubKegiatan" binding:"required"`
	NamaSubKegiatan string `json:"namaSubKegiatan" binding:"required"`
}

type UpdateSubKegiatanRequest struct {
	KegiatanId      int    `json:"kegiatanId"`
	Tahun           string `json:"tahun"`
	KodeSubKegiatan string `json:"kodeSubKegiatan"`
	NamaSubKegiatan string `json:"namaSubKegiatan"`
}
