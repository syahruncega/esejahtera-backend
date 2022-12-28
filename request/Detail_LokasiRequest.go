package request

type CreateDetail_LokasiRequest struct {
	KabupatenKotaId     string `json:"kabupatenKotaId" binding:"required"`
	KecamatanId         string `json:"kecamatanId" binding:"required"`
	KelurahanId         string `json:"kelurahanId" binding:"required"`
	DetailSubKegiatanId int    `json:"detailSubKegiatanId" binding:"required"`
}

type UpdateDetail_LokasiRequest struct {
	KabupatenKotaId     string `json:"kabupatenKotaId"`
	KecamatanId         string `json:"kecamatanId"`
	KelurahanId         string `json:"kelurahanId"`
	DetailSubKegiatanId int    `json:"detailSubKegiatanId"`
}
