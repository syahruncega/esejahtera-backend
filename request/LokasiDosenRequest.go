package request

type CreateLokasiDosenRequest struct {
	DosenId     int    `json:"dosenId" binding:"required"`
	KelurahanId string `json:"kelurahanId" binding:"required"`
}

type UpdateLokasiDosenRequest struct {
	DosenId     int    `json:"dosenId"`
	KelurahanId string `json:"kelurahanId"`
}
