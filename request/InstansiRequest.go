package request

type CreateInstansiRequest struct {
	KodeInstansi string `json:"KodeInstansi" binding:"required"`
	NamaInstansi string `json:"namaInstansi" binding:"required"`
}

type UpdateInstansiRequest struct {
	KodeInstansi string `json:"kodeInstansi"`
	NamaInstansi string `json:"namaInstansi"`
}
