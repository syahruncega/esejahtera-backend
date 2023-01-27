package request

type CreateMahasiswaRequest struct {
	UserId          int    `json:"userId" binding:"required"`
	NamaLengkap     string `json:"namaLengkap"`
	Universitas     string `json:"universitas"`
	JenisKelamin    string `json:"jenisKelamin"`
	TanggalLahir    string `json:"tanggalLahir"`
	KabupatenKotaId string `json:"kabupatenKotaId" binding:"required"`
	KecamatanId     string `json:"kecamatanid" binding:"required"`
	KelurahanId     string `json:"kelurahanId" binding:"required"`
}

type UpdateMahasiswaRequest struct {
	UserId          int    `json:"userId"`
	NamaLengkap     string `json:"namaLengkap"`
	Universitas     string `json:"universitas"`
	JenisKelamin    string `json:"jenisKelamin"`
	TanggalLahir    string `json:"tanggalLahir"`
	KabupatenKotaId string `json:"kabupatenKotaId"`
	KecamatanId     string `json:"kecamatanid"`
	KelurahanId     string `json:"kelurahanId"`
}
