package request

type CreateDosenRequest struct {
	UserId      int    `json:"userId" binding:"required"`
	NamaLengkap string `json:"namaLengkap"`
	Universitas string `json:"universitas"`
}

type UpdateDosenRequest struct {
	UserId      int    `json:"userId"`
	NamaLengkap string `json:"namaLengkap"`
	Universitas string `json:"universitas"`
}
