package request

type CreateSubKegiatanRequest struct {
	Tahun           string `json:"tahun" binding:"required"`
	KodeSubKegiatan string `json:"kodeSubKegiatan" binding:"required"`
	NamaSubKegiatan string `json:"namaSubKegiatan" binding:"required"`
}

type UpdateSubKegiatanRequest struct {
	Tahun           string `json:"tahun"`
	KodeSubKegiatan string `json:"kodeSubKegiatan"`
	NamaSubKegiatan string `json:"namaSubKegiatan"`
}
