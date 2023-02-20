package request

type CreateAnalisRequest struct {
	UserId        int    `json:"userId" binding:"required"`
	NamaLengkap   string `json:"namaLengkap"`
	Universitas   string `json:"universitas"`
	UrlFotoProfil string `json:"urlFotoProfil"`
}

type UpdateAnalisRequest struct {
	UserId        int    `json:"userId"`
	NamaLengkap   string `json:"namaLengkap"`
	Universitas   string `json:"universitas"`
	UrlFotoProfil string `json:"urlFotoProfil"`
}
