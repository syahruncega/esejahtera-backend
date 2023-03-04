package request

type CreateBidangUrusanOnInstansiRequest struct {
	InstansiId     int `json:"instansiId" binding:"required"`
	BidangUrusanId int `json:"bidangUrusanId" binding:"required"`
}

type UpdateBidangUrusanOnInstansiRequest struct {
	InstansiId     int `json:"instansiId"`
	BidangUrusanId int `json:"bidangUrusanId"`
}
