package request

type CreateInstansiRequest struct {
	NamaInstansi   string `json:"namaInstansi" binding:"required"`
	BidangUrusanId int    `json:"bidangUrusanId" binding:"required"`
}

type UpdateInstansiRequest struct {
	NamaInstansi   string `json:"namaInstansi"`
	BidangUrusanId int    `json:"bidangUrusanId"`
}
