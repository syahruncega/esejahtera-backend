package request

type CreateDetail_LokasiRequest struct {
	FokusBelanjaId  int    `json:"fokusBelanjaId" binding:"required"`
	KabupatenKotaId string `json:"kabupatenKotaId" binding:"required"`
	KecamatanId     string `json:"kecamatanId" binding:"required"`
	KelurahanId     string `json:"kelurahanId" binding:"required"`
	TipePenerima    string `json:"tipePenerima" binding:"required"`
}

type UpdateDetail_LokasiRequest struct {
	FokusBelanjaId  int    `json:"fokusBelanjaId"`
	KabupatenKotaId string `json:"kabupatenKotaId"`
	KecamatanId     string `json:"kecamatanId"`
	KelurahanId     string `json:"kelurahanId"`
	TipePenerima    string `json:"tipePenerima"`
}
