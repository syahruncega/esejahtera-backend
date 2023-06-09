package request

type CreateTagIndividuRequest struct {
	FokusBelanjaId int `json:"fokusBelanjaId" binding:"required"`
	IndividuId     int `json:"individuId" binding:"required"`
}

type UpdateTagIndividuRequest struct {
	FokusBelanjaId int `json:"fokusBelanjaId"`
	IndividuId     int `json:"individuId"`
}
