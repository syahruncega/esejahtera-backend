package request

type CreateAdminOpdRequest struct {
	UserId        int    `json:"userId" binding:"required"`
	NamaLengkap   string `json:"namaLengkap"`
	InstansiId    int    `json:"instansiId" binding:"required"`
	UrlFotoProfil string `json:"urlFotoProfil"`
}

type UpdateAdminOpdRequest struct {
	UserId        int    `json:"userId"`
	NamaLengkap   string `json:"namaLengkap"`
	InstansiId    int    `json:"instansiId"`
	UrlFotoProfil string `json:"urlFotoProfil"`
}
