package request

type CreateBidangUrusan struct {
	KodeBidangUrusan string `json:"kodeBidangUrusan" binding:"required"`
	NamaBidangUrusan string `json:"namaBidangUrusan" binding:"required"`
}

type UpdateBidangUrusan struct {
	KodeBidangUrusan string `json:"kodeBidangUrusan"`
	NamaBidangUrusan string `json:"namaBidangUrusan"`
}
