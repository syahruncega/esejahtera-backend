package request

type CreateBidangUrusan struct {
	Id               string `json:"id" binding:"required"`
	NamaBidangUrusan string `json:"namaBidangUrusan" binding:"required"`
}

type UpdateBidangUrusan struct {
	Id               string `json:"id"`
	NamaBidangUrusan string `json:"namaBidangUrusan"`
}
