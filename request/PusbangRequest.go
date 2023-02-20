package request

type CreatePusbangRequest struct {
	UserId        int    `json:"userId" binding:"required"`
	NamaLengkap   string `json:"namaLengkap"`
	Universitas   string `json:"universitas"`
	UrlFotoProfil string `json:"urlFotoProfil"`
}

type UpdatePusbangRequest struct {
	UserId        int    `json:"userId"`
	NamaLengkap   string `json:"namaLengkap"`
	Universitas   string `json:"universitas"`
	UrlFotoProfil string `json:"urlFotoProfil"`
}
