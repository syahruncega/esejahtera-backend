package request

type CreateDetailSubKegiatanRequest struct {
	SubKegiatanId   int    `json:"subKegiatanId" binding:"required"`
	PaguSubKegiatan int64  `json:"paguSubKegiatan" binding:"required"`
	Tipe            string `json:"tipe" binding:"required"`
}

type UpdateDetailSubKegiatanRequest struct {
	SubKegiatanId   int    `json:"subKegiatanId"`
	PaguSubKegiatan int64  `json:"paguSubKegiatan"`
	Tipe            string `json:"tipe"`
}
