package request

type CreateLokasiDosenRequest struct {
	DosenId         int    `json:"dosenId" binding:"required"`
	KabupatenKotaId string `json:"kabupatenKotaId" binding:"required"`
	KecamatanId     string `json:"kecamatanId" binding:"required"`
	KelurahanId     string `json:"kelurahanId" binding:"required"`
}

type UpdateLokasiDosenRequest struct {
	DosenId         int    `json:"dosenId"`
	KabupatenKotaId string `json:"kabupatenKotaId"`
	KecamatanId     string `json:"kecamatanId"`
	KelurahanId     string `json:"kelurahanId"`
}
