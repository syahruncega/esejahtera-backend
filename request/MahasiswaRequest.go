package request

type CreateMahasiswaRequest struct {
	UserId       int    `json:"userId" binding:"required"`
	NamaLengkap  string `json:"namaLengkap"`
	Universitas  string `json:"universitas"`
	JenisKelamin string `json:"jenisKelamin"`
	TanggalLahir string `json:"tanggalLahir"`
	KelurahanId  string `json:"kelurahanID" binding:"required"`
}

type UpdateMahasiswaRequest struct {
	UserId       int    `json:"userId"`
	NamaLengkap  string `json:"namaLengkap"`
	Universitas  string `json:"universitas"`
	JenisKelamin string `json:"jenisKelamin"`
	TanggalLahir string `json:"tanggalLahir"`
	KelurahanId  string `json:"kelurahanId"`
}
