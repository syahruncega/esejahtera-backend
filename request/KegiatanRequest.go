package request

type CreateKegiatanRequest struct {
	Tahun        string `json:"tahun" binding:"required"`
	KodeKegiatan string `json:"kodeKegiatan" binding:"required"`
	NamaKegiatan string `json:"namaKegiatan" binding:"required"`
}

type UpdateKegiatanRequest struct {
	Tahun        string `json:"tahun"`
	KodeKegiatan string `json:"kodeKegiatan"`
	NamaKegiatan string `json:"namaKegiatan"`
}
