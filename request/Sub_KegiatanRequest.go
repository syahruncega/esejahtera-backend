package request

type CreateSub_KegiatanRequest struct {
	NamaSubKegiatan             string `json:"namaSubKegiatan" binding:"required"`
	IndikatorKinerjaSubKegiatan string `json:"indikatorKinerjaSubKegiatan" binding:"required"`
	PaguSubKegiatan             int    `json:"paguSubKegiatan"`
	KegiatanId                  int    `json:"kegiatanId" binding:"required"`
}

type UpdateSub_KegiatanRequest struct {
	NamaSubKegiatan             string `json:"namaSubKegiatan"`
	IndikatorKinerjaSubKegiatan string `json:"indikatorKinerjaSubKegiatan"`
	PaguSubKegiatan             int    `json:"paguSubKegiatan"`
	KegiatanId                  int    `json:"kegiatanId"`
}
