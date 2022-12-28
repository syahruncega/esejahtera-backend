package request

type CreateInstansiRequest struct {
	NamaInstansi string `json:"namaInstansi" binding:"required"`
}

type UpdateInstansiRequest struct {
	NamaInstansi string `json:"namaInstansi"`
}
