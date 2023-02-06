package request

type CreateInstansiRequest struct {
	InstansiId   string `json:"instansiId" binding:"required"`
	NamaInstansi string `json:"namaInstansi" binding:"required"`
}

type UpdateInstansiRequest struct {
	InstansiId   string `json:"instansiId"`
	NamaInstansi string `json:"namaInstansi"`
}
