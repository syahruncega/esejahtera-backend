package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type ProgramOnKegiatanRepository interface {
	FindAll() ([]model.ProgramOnKegiatan, error)
	FindById(id int) (model.ProgramOnKegiatan, error)
	FindByProgramId(programId int) ([]model.ProgramOnKegiatan, error)
	FindBySearch(whereClause map[string]interface{}, tahun string) ([]model.ProgramOnKegiatan, error)
	CountJumlahKegiatanAllIntansi(tahun string, instansis []model.Instansi) []int64
	CountJumlahKegiatanByInstansiId(tahun string, instansi model.Instansi) int64
	Create(programOnKegiatan model.ProgramOnKegiatan) (model.ProgramOnKegiatan, error)
	Update(programOnKegiatan model.ProgramOnKegiatan) (model.ProgramOnKegiatan, error)
	Delete(programOnKegiatan model.ProgramOnKegiatan) (model.ProgramOnKegiatan, error)
}

type programOnKegiatanRepository struct {
	db *gorm.DB
}

func NewProgramOnKegiatanRepository(db *gorm.DB) *programOnKegiatanRepository {
	return &programOnKegiatanRepository{db}
}

func (r *programOnKegiatanRepository) FindAll() ([]model.ProgramOnKegiatan, error) {
	var programOnKegiatans []model.ProgramOnKegiatan

	var err = r.db.Model(&programOnKegiatans).Preload("Program").Preload("Kegiatan").Find(&programOnKegiatans).Error

	return programOnKegiatans, err
}

func (r *programOnKegiatanRepository) FindById(id int) (model.ProgramOnKegiatan, error) {
	var programOnKegiatan model.ProgramOnKegiatan

	var err = r.db.Model(&programOnKegiatan).Preload("Program").Preload("Kegiatan").Take(&programOnKegiatan, id).Error

	return programOnKegiatan, err
}

func (r *programOnKegiatanRepository) FindByProgramId(programId int) ([]model.ProgramOnKegiatan, error) {
	var programOnKegiatans []model.ProgramOnKegiatan

	var err = r.db.Where("programId = ?", programId).Model(&programOnKegiatans).Preload("Program").Preload("Kegiatan").Find(&programOnKegiatans).Error

	return programOnKegiatans, err
}

func (r *programOnKegiatanRepository) FindBySearch(whereClause map[string]interface{}, tahun string) ([]model.ProgramOnKegiatan, error) {
	var programOnKegiatans []model.ProgramOnKegiatan
	var err error

	if tahun != "" {
		err = r.db.Where(whereClause).Model(&programOnKegiatans).Preload("Program").Preload("Kegiatan", "tahun = ?", tahun).Find(&programOnKegiatans).Error
	} else {
		err = r.db.Where(whereClause).Model(&programOnKegiatans).Preload("Program").Preload("Kegiatan").Find(&programOnKegiatans).Error
	}

	if len(programOnKegiatans) == 0 {
		return programOnKegiatans, gorm.ErrRecordNotFound
	}

	return programOnKegiatans, err

}

func (r *programOnKegiatanRepository) CountJumlahKegiatanAllIntansi(tahun string, instansis []model.Instansi) []int64 {
	var count int64
	var hasil []int64

	for i := 0; i < len(instansis); i++ {
		var _ = r.db.Select("count(*)").Table("instansis as i").Joins("inner join instansi_on_programs as iop on i.id = iop.instansiId").Joins("inner join programs as p on p.id = iop.programId").Joins("inner join program_on_kegiatans as pok on pok.programId = p.id").Joins("inner join kegiatans as k on pok.kegiatanId = k.id").Where("p.tahun = ? and i.id = ?", tahun, instansis[i].Id).Scan(&count)

		hasil = append(hasil, count)
	}

	return hasil
}

func (r *programOnKegiatanRepository) CountJumlahKegiatanByInstansiId(tahun string, instansi model.Instansi) int64 {
	var count int64

	var _ = r.db.Select("count(*)").Table("instansis as i").Joins("inner join instansi_on_programs as iop on i.id = iop.instansiId").Joins("inner join programs as p on p.id = iop.programId").Joins("inner join program_on_kegiatans as pok on pok.programId = p.id").Joins("inner join kegiatans as k on pok.kegiatanId = k.id").Where("p.tahun = ? and i.id = ?", tahun, instansi.Id).Scan(&count)

	return count
}

func (r *programOnKegiatanRepository) Create(programOnKegiatan model.ProgramOnKegiatan) (model.ProgramOnKegiatan, error) {
	var err = r.db.Create(&programOnKegiatan).Error

	return programOnKegiatan, err
}

func (r *programOnKegiatanRepository) Update(programOnKegiatan model.ProgramOnKegiatan) (model.ProgramOnKegiatan, error) {
	var err = r.db.Model(&programOnKegiatan).Updates(model.ProgramOnKegiatan{
		ProgramId:  programOnKegiatan.ProgramId,
		KegiatanId: programOnKegiatan.KegiatanId,
	}).Error

	return programOnKegiatan, err
}

func (r *programOnKegiatanRepository) Delete(programOnKegiatan model.ProgramOnKegiatan) (model.ProgramOnKegiatan, error) {
	var err = r.db.Delete(&programOnKegiatan).Error

	return programOnKegiatan, err
}
