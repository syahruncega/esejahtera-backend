package request

type CreatePusbangRequest struct {
	UserId      int    `json:"userId" binding:"required"`
	NamaLengkap string `json:"namaLengkap"`
	Universitas string `json:"universitas"`
}

type UpdatePusbangRequest struct {
	UserId      int    `json:"userId"`
	NamaLengkap string `json:"namaLengkap"`
	Universitas string `json:"universitas"`
}
