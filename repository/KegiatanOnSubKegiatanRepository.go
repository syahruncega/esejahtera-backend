package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type KegiatanOnSubKegiatanRepository interface {
	FindAll() ([]model.KegiatanOnSubKegiatan, error)
	FindById(id int) (model.KegiatanOnSubKegiatan, error)
	FindByKegiatanId(kegiatanId int) ([]model.KegiatanOnSubKegiatan, error)
	FindBySearch(whereClause map[string]interface{}, tahun string) ([]model.KegiatanOnSubKegiatan, error)
	CountJumlahSubKegiatanAllInstansi(tahun string, instansis []model.Instansi) []int64
	CountJumlahSubKegiatanByInstansiId(tahun string, instansi model.Instansi) int64
	Create(kegiatanOnSubKegiatan model.KegiatanOnSubKegiatan) (model.KegiatanOnSubKegiatan, error)
	Update(kegiatanOnSubKegiatan model.KegiatanOnSubKegiatan) (model.KegiatanOnSubKegiatan, error)
	Delete(kegiatanOnSubKegiatan model.KegiatanOnSubKegiatan) (model.KegiatanOnSubKegiatan, error)
}

type kegiatanOnSubKegiatanRepository struct {
	db *gorm.DB
}

func NewKegiatanOnSubKegiatanRepository(db *gorm.DB) *kegiatanOnSubKegiatanRepository {
	return &kegiatanOnSubKegiatanRepository{db}
}

func (r *kegiatanOnSubKegiatanRepository) FindAll() ([]model.KegiatanOnSubKegiatan, error) {
	var kegiatanOnSubKegiatan []model.KegiatanOnSubKegiatan

	var err = r.db.Model(&kegiatanOnSubKegiatan).Preload("Kegiatan").Preload("SubKegiatan").Find(&kegiatanOnSubKegiatan).Error

	return kegiatanOnSubKegiatan, err
}

func (r *kegiatanOnSubKegiatanRepository) FindById(id int) (model.KegiatanOnSubKegiatan, error) {
	var kegiatanOnSubKegiatan model.KegiatanOnSubKegiatan

	var err = r.db.Model(&kegiatanOnSubKegiatan).Preload("Kegiatan").Preload("SubKegiatan").Take(&kegiatanOnSubKegiatan, id).Error

	return kegiatanOnSubKegiatan, err
}

func (r *kegiatanOnSubKegiatanRepository) FindByKegiatanId(kegiatanId int) ([]model.KegiatanOnSubKegiatan, error) {
	var kegiatanOnSubKegiatans []model.KegiatanOnSubKegiatan

	var err = r.db.Where("kegiatanId = ?", kegiatanId).Model(&kegiatanOnSubKegiatans).Preload("Kegiatan").Preload("SubKegiatan").Find(&kegiatanOnSubKegiatans).Error

	return kegiatanOnSubKegiatans, err
}

func (r *kegiatanOnSubKegiatanRepository) FindBySearch(whereClause map[string]interface{}, tahun string) ([]model.KegiatanOnSubKegiatan, error) {
	var kegiatanOnSubKegiatans []model.KegiatanOnSubKegiatan
	var err error

	if tahun != "" {
		err = r.db.Where(whereClause).Model(&kegiatanOnSubKegiatans).Preload("Kegiatan").Preload("SubKegiatan", "tahun = ?", tahun).Find(&kegiatanOnSubKegiatans).Error
	} else {
		err = r.db.Where(whereClause).Model(&kegiatanOnSubKegiatans).Preload("Kegiatan").Preload("SubKegiatan").Find(&kegiatanOnSubKegiatans).Error
	}

	if len(kegiatanOnSubKegiatans) == 0 {
		return kegiatanOnSubKegiatans, gorm.ErrRecordNotFound
	}

	return kegiatanOnSubKegiatans, err
}

func (r *kegiatanOnSubKegiatanRepository) CountJumlahSubKegiatanAllInstansi(tahun string, instansis []model.Instansi) []int64 {
	var count int64
	var hasil []int64

	for i := 0; i < len(instansis); i++ {
		var _ = r.db.Select("count(*)").Table("instansis as i").Joins("inner join instansi_on_programs as iop on i.id = iop.instansiId").Joins("inner join programs as p on p.id = iop.programId").Joins("inner join program_on_kegiatans as pok on pok.programId = p.id").Joins("inner join kegiatans as k on pok.kegiatanId = k.id").Joins("inner join kegiatan_on_sub_kegiatans as kosk on kosk.kegiatanId = k.id").Where("p.tahun = ? and i.id = ?", tahun, instansis[i].Id).Scan(&count)

		hasil = append(hasil, count)
	}

	return hasil
}

func (r *kegiatanOnSubKegiatanRepository) CountJumlahSubKegiatanByInstansiId(tahun string, instansi model.Instansi) int64 {
	var count int64

	var _ = r.db.Select("count(*)").Table("instansis as i").Joins("inner join instansi_on_programs as iop on i.id = iop.instansiId").Joins("inner join programs as p on p.id = iop.programId").Joins("inner join program_on_kegiatans as pok on pok.programId = p.id").Joins("inner join kegiatans as k on pok.kegiatanId = k.id").Joins("inner join kegiatan_on_sub_kegiatans as kosk on kosk.kegiatanId = k.id").Where("p.tahun = ? and i.id = ?", tahun, instansi.Id).Scan(&count)

	return count
}

func (r *kegiatanOnSubKegiatanRepository) Create(kegiatanOnSubKegiatan model.KegiatanOnSubKegiatan) (model.KegiatanOnSubKegiatan, error) {
	var err = r.db.Create(&kegiatanOnSubKegiatan).Error

	return kegiatanOnSubKegiatan, err
}

func (r *kegiatanOnSubKegiatanRepository) Update(kegiatanOnSubKegiatan model.KegiatanOnSubKegiatan) (model.KegiatanOnSubKegiatan, error) {
	var err = r.db.Model(&kegiatanOnSubKegiatan).Updates(model.KegiatanOnSubKegiatan{
		KegiatanId:    kegiatanOnSubKegiatan.KegiatanId,
		SubKegiatanId: kegiatanOnSubKegiatan.SubKegiatanId,
	}).Error

	return kegiatanOnSubKegiatan, err
}

func (r *kegiatanOnSubKegiatanRepository) Delete(kegiatanOnSubKegiatan model.KegiatanOnSubKegiatan) (model.KegiatanOnSubKegiatan, error) {
	var err = r.db.Delete(&kegiatanOnSubKegiatan).Error

	return kegiatanOnSubKegiatan, err
}
