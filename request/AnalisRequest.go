package request

type CreateAnalisRequest struct {
	UserId      int    `json:"userId" binding:"required"`
	NamaLengkap string `json:"namaLengkap"`
	Universitas string `json:"universitas"`
}

type UpdateAnalisRequest struct {
	UserId      int    `json:"userId"`
	NamaLengkap string `json:"namaLengkap"`
	Universitas string `json:"universitas"`
}
