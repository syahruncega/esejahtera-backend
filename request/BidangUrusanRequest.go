package request

type CreateBidangUrusan struct {
	BidangUrusanId   string `json:"bidangUrusanId" binding:"required"`
	NamaBidangUrusan string `json:"namaBidangUrusan" binding:"required"`
}

type UpdateBidangUrusan struct {
	BidangUrusanId   string `json:"bidangUrusanId"`
	NamaBidangUrusan string `json:"namaBidangUrusan"`
}
