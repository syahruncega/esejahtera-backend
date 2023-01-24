package helper

import (
	"kemiskinan/model"
	"kemiskinan/responses"
)

func ConvertToBidangUrusanResponse(b model.BidangUrusan) responses.BidangUrusanResponse {
	return responses.BidangUrusanResponse{
		Id:               b.Id,
		NamaBidangUrusan: b.NamaBidangUrusan,
		CreatedAt:        b.CreatedAt,
		UpdatedAt:        b.UpdatedAt,
	}
}

func ConvertToInstansiResponse(i model.Instansi) responses.InstansiResponse {
	return responses.InstansiResponse{
		Id:             i.Id,
		BidangUrusanId: i.BidangUrusanId,
		NamaInstansi:   i.NamaInstansi,
		CreatedAt:      i.CreatedAt,
		UpdatedAt:      i.UpdatedAt,
	}
}

func ConvertToInstansiWithBidangUrusanResponse(i model.Instansi) responses.InstansiWithBidangUrusanResponse {
	return responses.InstansiWithBidangUrusanResponse{
		Id:             i.Id,
		NamaInstansi:   i.NamaInstansi,
		CreatedAt:      i.CreatedAt,
		UpdatedAt:      i.UpdatedAt,
		BidangUrusanId: i.BidangUrusanId,
		BidangUrusan: responses.BidangUrusanResponse{
			Id:               i.BidangUrusan.Id,
			NamaBidangUrusan: i.BidangUrusan.NamaBidangUrusan,
			CreatedAt:        i.BidangUrusan.CreatedAt,
			UpdatedAt:        i.BidangUrusan.UpdatedAt,
		},
	}
}

func ConvertToProgramResponse(p model.Program) responses.ProgramResponse {
	return responses.ProgramResponse{
		Id:                      p.Id,
		Sasaran:                 p.Sasaran,
		IndikatorSasaran:        p.IndikatorSasaran,
		Kebijakan:               p.Kebijakan,
		NamaProgram:             p.NamaProgram,
		IndikatorKinerjaProgram: p.IndikatorKinerjaProgram,
		PaguProgram:             p.PaguProgram,
		InstansiId:              p.InstansiId,
		CreatedAt:               p.CreatedAt,
		UpdatedAt:               p.UpdatedAt,
	}
}

func ConvertToProgramWithInstansiResponse(p model.Program) responses.ProgramWithInstansiResponse {
	return responses.ProgramWithInstansiResponse{
		Id:                      p.Id,
		Sasaran:                 p.Sasaran,
		IndikatorSasaran:        p.IndikatorSasaran,
		Kebijakan:               p.Kebijakan,
		NamaProgram:             p.NamaProgram,
		IndikatorKinerjaProgram: p.IndikatorKinerjaProgram,
		PaguProgram:             p.PaguProgram,
		InstansiId:              p.InstansiId,
		CreatedAt:               p.CreatedAt,
		UpdatedAt:               p.UpdatedAt,
		Instansi: responses.InstansiResponse{
			Id:           p.Instansi.Id,
			NamaInstansi: p.Instansi.NamaInstansi,
			CreatedAt:    p.Instansi.CreatedAt,
			UpdatedAt:    p.Instansi.UpdatedAt,
		},
	}
}

func ConvertToKegiatanResponse(k model.Kegiatan) responses.KegiatanResponse {
	return responses.KegiatanResponse{
		Id:                       k.Id,
		NamaKegiatan:             k.NamaKegiatan,
		IndikatorKinerjaKegiatan: k.IndikatorKinerjaKegiatan,
		PaguKegiatan:             k.PaguKegiatan,
		ProgramId:                k.ProgramId,
		CreatedAt:                k.CreatedAt,
		UpdatedAt:                k.UpdatedAt,
	}
}

func ConvertToKegiatanWithProgramResponse(k model.Kegiatan) responses.KegiatanWithProgramResponse {
	return responses.KegiatanWithProgramResponse{
		Id:                       k.Id,
		NamaKegiatan:             k.NamaKegiatan,
		IndikatorKinerjaKegiatan: k.IndikatorKinerjaKegiatan,
		PaguKegiatan:             k.PaguKegiatan,
		ProgramId:                k.ProgramId,
		CreatedAt:                k.CreatedAt,
		UpdatedAt:                k.UpdatedAt,
		Program: responses.ProgramResponse{
			Id:                      k.Program.Id,
			NamaProgram:             k.Program.NamaProgram,
			IndikatorKinerjaProgram: k.Program.IndikatorKinerjaProgram,
			PaguProgram:             k.Program.PaguProgram,
			InstansiId:              k.Program.InstansiId,
			CreatedAt:               k.Program.CreatedAt,
			UpdatedAt:               k.Program.UpdatedAt,
		},
	}
}

func ConvertToSubKegiatanResponse(s model.Sub_Kegiatan) responses.Sub_KegiatanResponse {
	return responses.Sub_KegiatanResponse{
		Id:                          s.Id,
		NamaSubKegiatan:             s.NamaSubKegiatan,
		IndikatorKinerjaSubKegiatan: s.IndikatorKinerjaSubKegiatan,
		PaguSubKegiatan:             s.PaguSubKegiatan,
		KegiatanId:                  s.KegiatanId,
		CreatedAt:                   s.CreatedAt,
		UpdatedAt:                   s.UpdatedAt,
	}
}

func ConvertToSubKegiatanWithKegiatanResponse(s model.Sub_Kegiatan) responses.Sub_KegiatanWithKegiatanResponse {
	return responses.Sub_KegiatanWithKegiatanResponse{
		Id:                          s.Id,
		NamaSubKegiatan:             s.NamaSubKegiatan,
		IndikatorKinerjaSubKegiatan: s.IndikatorKinerjaSubKegiatan,
		PaguSubKegiatan:             s.PaguSubKegiatan,
		CreatedAt:                   s.CreatedAt,
		UpdatedAt:                   s.UpdatedAt,
		KegiatanId:                  s.KegiatanId,
		Kegiatan: responses.KegiatanResponse{
			Id:                       s.Kegiatan.Id,
			NamaKegiatan:             s.Kegiatan.NamaKegiatan,
			IndikatorKinerjaKegiatan: s.Kegiatan.IndikatorKinerjaKegiatan,
			PaguKegiatan:             s.Kegiatan.PaguKegiatan,
			CreatedAt:                s.Kegiatan.CreatedAt,
			UpdatedAt:                s.Kegiatan.UpdatedAt,
			ProgramId:                s.Kegiatan.ProgramId,
		},
	}
}

func ConvertToDetailSubKegiatanResponse(d model.Detail_Sub_Kegiatan) responses.Detail_Sub_KegiatanResponse {
	return responses.Detail_Sub_KegiatanResponse{
		Id:               d.Id,
		FokusBelanja:     d.FokusBelanja,
		Indikator:        d.Indikator,
		Target:           d.Target,
		Satuan:           d.Satuan,
		PaguFokusBelanja: d.PaguFokusBelanja,
		Keterangan:       d.Keterangan,
		SubKegiatanId:    d.SubKegiatanId,
		CreatedAt:        d.CreatedAt,
		UpdatedAt:        d.UpdatedAt,
	}
}

func ConvertToDetailSubKegiatanWithSubKegiatanResponse(d model.Detail_Sub_Kegiatan) responses.Detail_Sub_KegiatanWithSub_KegiatanResponse {
	return responses.Detail_Sub_KegiatanWithSub_KegiatanResponse{
		Id:               d.Id,
		FokusBelanja:     d.FokusBelanja,
		Indikator:        d.Indikator,
		Target:           d.Target,
		Satuan:           d.Satuan,
		PaguFokusBelanja: d.PaguFokusBelanja,
		Keterangan:       d.Keterangan,
		CreatedAt:        d.CreatedAt,
		UpdatedAt:        d.UpdatedAt,
		SubKegiatanId:    d.SubKegiatanId,
		SubKegiatan: responses.Sub_KegiatanResponse{
			Id:                          d.SubKegiatan.Id,
			NamaSubKegiatan:             d.SubKegiatan.NamaSubKegiatan,
			IndikatorKinerjaSubKegiatan: d.SubKegiatan.IndikatorKinerjaSubKegiatan,
			PaguSubKegiatan:             d.SubKegiatan.PaguSubKegiatan,
			CreatedAt:                   d.SubKegiatan.CreatedAt,
			UpdatedAt:                   d.SubKegiatan.UpdatedAt,
			KegiatanId:                  d.SubKegiatan.KegiatanId,
		},
	}
}

func ConvertToDetailLokasiResponse(d model.Detail_Lokasi) responses.Detail_LokasiResponse {
	return responses.Detail_LokasiResponse{
		Id:                  d.Id,
		KabupatenKotaId:     d.KabupatenKotaId,
		KecamatanId:         d.KecamatanId,
		KelurahanId:         d.KelurahanId,
		DetailSubKegiatanId: d.DetailSubKegiatanId,
		CreatedAt:           d.CreatedAt,
		UpdatedAt:           d.UpdatedAt,
	}
}

func ConvertToDetailLokasiWithRelationResponse(d model.Detail_Lokasi) responses.Detail_LokasiWithRelationResponse {
	return responses.Detail_LokasiWithRelationResponse{
		Id:                  d.Id,
		KabupatenKotaId:     d.KabupatenKotaId,
		KecamatanId:         d.KecamatanId,
		KelurahanId:         d.KelurahanId,
		CreatedAt:           d.CreatedAt,
		UpdatedAt:           d.UpdatedAt,
		DetailSubKegiatanId: d.DetailSubKegiatanId,
		DetailSubKegiatan: responses.Detail_Sub_KegiatanResponse{
			Id:               d.DetailSubKegiatan.Id,
			FokusBelanja:     d.DetailSubKegiatan.FokusBelanja,
			Indikator:        d.DetailSubKegiatan.Indikator,
			Target:           d.DetailSubKegiatan.Target,
			Satuan:           d.DetailSubKegiatan.Satuan,
			PaguFokusBelanja: d.DetailSubKegiatan.PaguFokusBelanja,
			Keterangan:       d.DetailSubKegiatan.Keterangan,
			CreatedAt:        d.DetailSubKegiatan.CreatedAt,
			UpdatedAt:        d.DetailSubKegiatan.UpdatedAt,
			SubKegiatanId:    d.DetailSubKegiatan.SubKegiatanId,
		},
		KabupatenKota: responses.Kabupaten_KotaResponse{
			Id:         d.KabupatenKota.Id,
			ProvinsiId: d.KabupatenKota.ProvinsiId,
			Nama:       d.KabupatenKota.Nama,
		},
		Kecamatan: responses.KecamatanResponse{
			Id:              d.Kecamatan.Id,
			KabupatenKotaId: d.Kecamatan.KabupatenKotaId,
			Nama:            d.Kecamatan.Nama,
		},
		Kelurahan: responses.KelurahanResponse{
			Id:          d.Kelurahan.Id,
			KecamatanId: d.Kelurahan.KecamatanId,
			Nama:        d.Kelurahan.Nama,
		},
	}
}

func ConvertToProvinsiResponse(p model.Provinsi) responses.ProvinsiResponse {
	return responses.ProvinsiResponse{
		Id:   p.Id,
		Nama: p.Nama,
	}
}

func ConvertToKabupatenKotaResponse(k model.Kabupaten_Kota) responses.Kabupaten_KotaResponse {
	return responses.Kabupaten_KotaResponse{
		Id:         k.Id,
		ProvinsiId: k.ProvinsiId,
		Nama:       k.Nama,
	}
}

func ConvertToKecamatanResponse(k model.Kecamatan) responses.KecamatanResponse {
	return responses.KecamatanResponse{
		Id:              k.Id,
		KabupatenKotaId: k.KabupatenKotaId,
		Nama:            k.Nama,
	}
}

func ConvertToKelurahanResponse(k model.Kelurahan) responses.KelurahanResponse {
	return responses.KelurahanResponse{
		Id:          k.Id,
		KecamatanId: k.KecamatanId,
		Nama:        k.Nama,
	}
}

func ConvertToKeluargaResponse(k model.Keluarga) responses.KeluargaResponse {
	return responses.KeluargaResponse{
		Id:         k.Id,
		IdKeluarga: k.IdKeluarga,
		ProvinsiId: k.ProvinsiId,
		Provinsi: responses.ProvinsiResponse{
			Id:   k.Provinsi.Id,
			Nama: k.Provinsi.Nama,
		},

		KabupatenKotaId: k.KabupatenKotaId,
		KabupatenKota: responses.Kabupaten_KotaResponse{
			Id:         k.KabupatenKota.Id,
			ProvinsiId: k.KabupatenKota.ProvinsiId,
			Nama:       k.KabupatenKota.Nama,
		},

		KecamatanId: k.KecamatanId,
		Kecamatan: responses.KecamatanResponse{
			Id:              k.Kecamatan.Id,
			KabupatenKotaId: k.Kecamatan.KabupatenKotaId,
			Nama:            k.Kecamatan.Nama,
		},

		KelurahanId: k.KelurahanId,
		Kelurahan: responses.KelurahanResponse{
			Id:          k.Kelurahan.Id,
			KecamatanId: k.Kelurahan.KecamatanId,
			Nama:        k.Kelurahan.Nama,
		},

		DesilKesejahteraan:     k.DesilKesejahteraan,
		Alamat:                 k.Alamat,
		KepalaKeluarga:         k.KepalaKeluarga,
		Nik:                    k.Nik,
		PadanDukcapil:          k.PadanDukcapil,
		JenisKelamin:           k.JenisKelamin,
		TanggalLahir:           k.TanggalLahir,
		Pekerjaan:              k.Pekerjaan,
		Pendidikan:             k.Pendidikan,
		KepemilikanRumah:       k.KepemilikanRumah,
		Simpanan:               k.Simpanan,
		JenisAtap:              k.JenisAtap,
		JenisDinding:           k.JenisDinding,
		JenisLantai:            k.JenisLantai,
		SumberPenerangan:       k.SumberPenerangan,
		BahanBakarMemasak:      k.BahanBakarMemasak,
		SumberAirMinum:         k.SumberAirMinum,
		FasilitasBuangAirBesar: k.FasilitasBuangAirBesar,
		PenerimaBPNT:           k.PenerimaBPNT,
		PenerimaBPUM:           k.PenerimaBPUM,
		PenerimaBST:            k.PenerimaBST,
		PenerimaPKH:            k.PenerimaPKH,
		PenerimaSembako:        k.PenerimaSembako,

		UserId: k.UserId,
		User: responses.UserResponse{
			Id:        k.User.Id,
			Username:  k.User.Username,
			Password:  k.User.Password,
			Email:     k.User.Email,
			NoHp:      k.User.NoHp,
			Role:      k.User.Role,
			CreatedAt: k.User.CreatedAt,
			UpdatedAt: k.User.UpdatedAt,
		},

		StatusVerifikasi: k.StatusVerifikasi,
	}
}

func ConvertToMonevResponse(m model.Monev) responses.MonevResponse {
	return responses.MonevResponse{
		Id:              m.Id,
		NamaPenerima:    m.NamaPenerima,
		KabupatenKotaId: m.KabupatenKotaId,
		KabupatenKota: responses.Kabupaten_KotaResponse{
			Id:         m.KabupatenKota.Id,
			ProvinsiId: m.KabupatenKota.ProvinsiId,
			Nama:       m.KabupatenKota.Nama,
		},
		KecamatanId: m.KecamatanId,
		Kecamatan: responses.KecamatanResponse{
			Id:              m.Kecamatan.Id,
			KabupatenKotaId: m.Kecamatan.KabupatenKotaId,
			Nama:            m.Kecamatan.Nama,
		},
		KelurahanId: m.KelurahanId,
		Kelurahan: responses.KelurahanResponse{
			Id:          m.Kelurahan.Id,
			KecamatanId: m.Kelurahan.KecamatanId,
			Nama:        m.Kelurahan.Nama,
		},
		JenisBantuan:  m.JenisBantuan,
		VolumeBantuan: m.VolumeBantuan,
		SatuanVolume:  m.SatuanVolume,
	}
}

func ConvertToJumlahResponse(jumlah int64) responses.JumlahResponse {
	return responses.JumlahResponse{
		Jumlah: jumlah,
	}
}

func ConvertToUserResponse(u model.User) responses.UserResponse {
	return responses.UserResponse{
		Id:        u.Id,
		Username:  u.Username,
		Password:  u.Password,
		Email:     u.Email,
		NoHp:      u.NoHp,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ConvertToAdminResponse(a model.Admin) responses.AdminResponse {
	return responses.AdminResponse{
		Id:          a.Id,
		UserId:      a.UserId,
		NamaLengkap: a.NamaLengkap,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
	}
}

func ConvertToAdminWithRelationResponse(a model.Admin) responses.AdminWithRelationResponse {
	return responses.AdminWithRelationResponse{
		Id:     a.Id,
		UserId: a.UserId,
		User: responses.UserResponse{
			Id:        a.User.Id,
			Username:  a.User.Username,
			Password:  a.User.Password,
			Email:     a.User.Email,
			NoHp:      a.User.NoHp,
			Role:      a.User.Role,
			CreatedAt: a.User.CreatedAt,
			UpdatedAt: a.User.UpdatedAt,
		},
		NamaLengkap: a.NamaLengkap,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
	}
}

func ConvertToAnalisResponse(a model.Analis) responses.AnalisResponse {
	return responses.AnalisResponse{
		Id:          a.Id,
		UserId:      a.UserId,
		NamaLengkap: a.NamaLengkap,
		Universitas: a.Universitas,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
	}
}

func ConvertToAnalisWithRelationResponse(a model.Analis) responses.AnalisWithRelationResponse {
	return responses.AnalisWithRelationResponse{
		Id:     a.Id,
		UserId: a.UserId,
		User: responses.UserResponse{
			Id:        a.User.Id,
			Username:  a.User.Username,
			Password:  a.User.Password,
			Email:     a.User.Email,
			NoHp:      a.User.NoHp,
			Role:      a.User.Role,
			CreatedAt: a.User.CreatedAt,
			UpdatedAt: a.User.UpdatedAt,
		},
		NamaLengkap: a.NamaLengkap,
		Universitas: a.Universitas,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
	}
}

func ConvertToPusbangResponse(p model.Pusbang) responses.PusbangResponse {
	return responses.PusbangResponse{
		Id:          p.Id,
		UserId:      p.UserId,
		NamaLengkap: p.NamaLengkap,
		Universitas: p.Universitas,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func ConvertToPusbangWithRelationResponse(p model.Pusbang) responses.PusbangWithRelationResponse {
	return responses.PusbangWithRelationResponse{
		Id:     p.Id,
		UserId: p.UserId,
		User: responses.UserResponse{
			Id:        p.User.Id,
			Username:  p.User.Username,
			Password:  p.User.Password,
			Email:     p.User.Email,
			NoHp:      p.User.NoHp,
			Role:      p.User.Role,
			CreatedAt: p.User.CreatedAt,
			UpdatedAt: p.User.UpdatedAt,
		},
		NamaLengkap: p.NamaLengkap,
		Universitas: p.Universitas,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func ConvertToDosenResponse(d model.Dosen) responses.DosenResponse {
	return responses.DosenResponse{
		Id:          d.Id,
		UserId:      d.UserId,
		NamaLengkap: d.NamaLengkap,
		Universitas: d.Universitas,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}

func ConvertToDosenWithRelationResponse(d model.Dosen) responses.DosenWithRelationResponse {
	return responses.DosenWithRelationResponse{
		Id:     d.Id,
		UserId: d.UserId,
		User: responses.UserResponse{
			Id:        d.User.Id,
			Username:  d.User.Username,
			Password:  d.User.Password,
			Email:     d.User.Email,
			NoHp:      d.User.NoHp,
			Role:      d.User.Role,
			CreatedAt: d.User.CreatedAt,
			UpdatedAt: d.User.UpdatedAt,
		},
		NamaLengkap: d.NamaLengkap,
		Universitas: d.Universitas,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}

func ConvertToMahasiswaResponse(m model.Mahasiswa) responses.MahasiswaResponse {
	return responses.MahasiswaResponse{
		Id:           m.Id,
		UserId:       m.UserId,
		NamaLengkap:  m.NamaLengkap,
		Universitas:  m.Universitas,
		JenisKelamin: m.JenisKelamin,
		TanggalLahir: m.TanggalLahir,
		KelurahanId:  m.KelurahanId,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

func ConvertToMahasiswaWithRelationResponse(m model.Mahasiswa) responses.MahasiswaWithRelationResponse {
	return responses.MahasiswaWithRelationResponse{
		Id:     m.Id,
		UserId: m.UserId,
		User: responses.UserResponse{
			Id:        m.User.Id,
			Username:  m.User.Username,
			Password:  m.User.Password,
			Email:     m.User.Email,
			NoHp:      m.User.NoHp,
			Role:      m.User.Role,
			CreatedAt: m.User.CreatedAt,
			UpdatedAt: m.User.UpdatedAt,
		},
		NamaLengkap:  m.NamaLengkap,
		Universitas:  m.Universitas,
		JenisKelamin: m.JenisKelamin,
		TanggalLahir: m.TanggalLahir,
		KelurahanId:  m.KelurahanId,
		Kelurahan: responses.KelurahanResponse{
			Id:          m.Kelurahan.Id,
			KecamatanId: m.Kelurahan.KecamatanId,
			Nama:        m.Kelurahan.Nama,
		},
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ConvertToLokasiDosenResponse(l model.Lokasi_Dosen) responses.LokasiDosenResponse {
	return responses.LokasiDosenResponse{
		Id:          l.Id,
		DosenId:     l.DosenId,
		KelurahanId: l.KelurahanId,
		CreatedAt:   l.CreatedAt,
		UpdatedAt:   l.UpdatedAt,
	}
}

func ConvertToLokasiDosenWithRelationResponse(l model.Lokasi_Dosen) responses.LokasiDosenWithRelationResponse {
	return responses.LokasiDosenWithRelationResponse{
		Id:      l.Id,
		DosenId: l.DosenId,
		Dosen: responses.DosenResponse{
			Id:          l.Dosen.Id,
			UserId:      l.Dosen.UserId,
			NamaLengkap: l.Dosen.NamaLengkap,
			Universitas: l.Dosen.Universitas,
			CreatedAt:   l.Dosen.CreatedAt,
			UpdatedAt:   l.Dosen.UpdatedAt,
		},
		KelurahanId: l.KelurahanId,
		Kelurahan: responses.KelurahanResponse{
			Id:          l.Kelurahan.Id,
			KecamatanId: l.Kelurahan.KecamatanId,
			Nama:        l.Kelurahan.Nama,
		},
		CreatedAt: l.CreatedAt,
		UpdatedAt: l.UpdatedAt,
	}
}

func ConvertToKeluargaVerifikasiResponse(k model.KeluargaVerifikasi) responses.KeluargaVerifikasiResponse {
	return responses.KeluargaVerifikasiResponse{
		Id:         k.Id,
		IdKeluarga: k.IdKeluarga,
		ProvinsiId: k.ProvinsiId,
		Provinsi: responses.ProvinsiResponse{
			Id:   k.Provinsi.Id,
			Nama: k.Provinsi.Nama,
		},

		KabupatenKotaId: k.KabupatenKotaId,
		KabupatenKota: responses.Kabupaten_KotaResponse{
			Id:         k.KabupatenKota.Id,
			ProvinsiId: k.KabupatenKota.ProvinsiId,
			Nama:       k.KabupatenKota.Nama,
		},

		KecamatanId: k.KecamatanId,
		Kecamatan: responses.KecamatanResponse{
			Id:              k.Kecamatan.Id,
			KabupatenKotaId: k.Kecamatan.KabupatenKotaId,
			Nama:            k.Kecamatan.Nama,
		},

		KelurahanId:            k.KelurahanId,
		DesilKesejahteraan:     k.DesilKesejahteraan,
		Alamat:                 k.Alamat,
		KepalaKeluarga:         k.KepalaKeluarga,
		Nik:                    k.Nik,
		PadanDukcapil:          k.PadanDukcapil,
		JenisKelamin:           k.JenisKelamin,
		TanggalLahir:           k.TanggalLahir,
		Pekerjaan:              k.Pekerjaan,
		Pendidikan:             k.Pendidikan,
		KepemilikanRumah:       k.KepemilikanRumah,
		Simpanan:               k.Simpanan,
		JenisAtap:              k.JenisAtap,
		JenisDinding:           k.JenisDinding,
		JenisLantai:            k.JenisLantai,
		SumberPenerangan:       k.SumberPenerangan,
		BahanBakarMemasak:      k.BahanBakarMemasak,
		SumberAirMinum:         k.SumberAirMinum,
		FasilitasBuangAirBesar: k.FasilitasBuangAirBesar,
		PenerimaBPNT:           k.PenerimaBPNT,
		PenerimaBPUM:           k.PenerimaBPUM,
		PenerimaBST:            k.PenerimaBST,
		PenerimaPKH:            k.PenerimaPKH,
		PenerimaSembako:        k.PenerimaSembako,
		PenerimaLainnya:        k.PenerimaLainnya,
		StatusResponden:        k.StatusResponden,

		UserId: k.UserId,
		User: responses.UserResponse{
			Id:        k.User.Id,
			Username:  k.User.Username,
			Password:  k.User.Password,
			Email:     k.User.Email,
			NoHp:      k.User.NoHp,
			Role:      k.User.Role,
			CreatedAt: k.User.CreatedAt,
			UpdatedAt: k.User.UpdatedAt,
		},

		MahasiswaId: k.MahasiswaId,
		Mahasiswa: responses.MahasiswaResponse{
			Id:           k.Mahasiswa.Id,
			UserId:       k.UserId,
			NamaLengkap:  k.Mahasiswa.NamaLengkap,
			Universitas:  k.Mahasiswa.Universitas,
			JenisKelamin: k.JenisKelamin,
			TanggalLahir: k.TanggalLahir,
			KelurahanId:  k.KelurahanId,
		},

		Kelurahan: responses.KelurahanResponse{
			Id:          k.Kelurahan.Id,
			KecamatanId: k.Kelurahan.KecamatanId,
			Nama:        k.Kelurahan.Nama,
		},

		CreatedAt: k.CreatedAt,
		UpdatedAt: k.UpdatedAt,
	}
}
