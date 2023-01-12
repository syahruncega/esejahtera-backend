package request

type CreateBidangUrusan struct {
	NamaBidangUrusan string `json:"namaBidangUrusan" binding:"required"`
}

type UpdateBidangUrusan struct {
	NamaBidangUrusan string `json:"namaBidangUrusan"`
}
