package request

type CreateInstansiRequest struct {
	Id           string `json:"id" binding:"required"`
	NamaInstansi string `json:"namaInstansi" binding:"required"`
}

type UpdateInstansiRequest struct {
	Id           string `json:"id"`
	NamaInstansi string `json:"namaInstansi"`
}
