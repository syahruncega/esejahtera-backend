package request

type CreateInstansiRequest struct {
	BidangUrusanId int    `json:"bidangUrusanId" binding:"required"`
	KodeInstansi   string `json:"KodeInstansi" binding:"required"`
	NamaInstansi   string `json:"namaInstansi" binding:"required"`
}

type UpdateInstansiRequest struct {
	BidangUrusanId int    `json:"bidangUrusanId"`
	KodeInstansi   string `json:"kodeInstansi"`
	NamaInstansi   string `json:"namaInstansi"`
}
