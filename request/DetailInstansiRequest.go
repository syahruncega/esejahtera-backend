package request

type CreateDetailInstansiRequest struct {
	InstansiId     string `json:"instansiId" binding:"required"`
	BidangUrusanId string `json:"bidangUrusanId" binding:"required"`
}

type UpdateDetailInstansiRequest struct {
	InstansiId     string `json:"instansiId"`
	BidangUrusanId string `json:"bidangUrusanId"`
}
