package request

type CreateAdminRequest struct {
	UserId      int    `json:"userId" binding:"required"`
	NamaLengkap string `json:"namaLengkap"`
}

type UpdateAdminRequest struct {
	UserId      int    `json:"userId"`
	NamaLengkap string `json:"namaLengkap"`
}
