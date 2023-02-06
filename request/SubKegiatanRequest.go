package request

type CreateSubKegiatanRequest struct {
	SubKegiatanId   string `json:"subKegiatanId" binding:"required"`
	NamaSubKegiatan string `json:"namaSubKegiatan" binding:"required"`
}

type UpdateSubKegiatanRequest struct {
	SubKegiatanId   string `json:"subKegiatanId"`
	NamaSubKegiatan string `json:"namaSubKegiatan"`
}
