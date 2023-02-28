package request

type CreateDetailInstansiRequest struct {
	InstansiId     int `json:"instansiId" binding:"required"`
	BidangUrusanId int `json:"bidangUrusanId" binding:"required"`
}

type UpdateDetailInstansiRequest struct {
	InstansiId     int `json:"instansiId"`
	BidangUrusanId int `json:"bidangUrusanId"`
}
