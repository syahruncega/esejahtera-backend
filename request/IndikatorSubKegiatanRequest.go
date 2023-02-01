package request

type CreateIndikatorSubKegiatanRequest struct {
	KegiatanId                  string `json:"kegiatanId" binding:"required"`
	SubKegiatanId               string `json:"subKegiatanId" binding:"required"`
	IndikatorKinerjaSubKegiatan string `json:"indikatorKinerjaSubKegiatan" binding:"required"`
	PaguSubKegiatan             int    `json:"paguSubKegiatan"`
}

type UpdateIndikatorSubKegiatanRequest struct {
	KegiatanId                  string `json:"kegiatanId"`
	SubKegiatanId               string `json:"subKegiatanId"`
	IndikatorKinerjaSubKegiatan string `json:"indikatorKinerjaSubKegiatan"`
	PaguSubKegiatan             int    `json:"paguSubKegiatan"`
}
