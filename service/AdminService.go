package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type AdminService interface {
	FindAll() ([]model.Admin, error)
	FindById(id int) (model.Admin, error)
	FindByUserId(userId int) (model.Admin, error)
	FindAllRelation() ([]model.Admin, error)
	Create(adminRequest request.CreateAdminRequest) (model.Admin, error)
	Update(id int, adminRequest request.UpdateAdminRequest) (model.Admin, error)
	Delete(id int) (model.Admin, error)
}

type adminService struct {
	adminRepository repository.AdminRepository
}

func NewAdminService(adminRepository repository.AdminRepository) *adminService {
	return &adminService{adminRepository}
}

func (s *adminService) FindAll() ([]model.Admin, error) {
	var admins, err = s.adminRepository.FindAll()

	return admins, err
}

func (s *adminService) FindById(id int) (model.Admin, error) {
	var admin, err = s.adminRepository.FindById(id)

	return admin, err
}

func (s *adminService) FindByUserId(userId int) (model.Admin, error) {
	var admin, err = s.adminRepository.FindByUserId(userId)

	return admin, err
}

func (s *adminService) FindAllRelation() ([]model.Admin, error) {
	var admins, err = s.adminRepository.FindAllRelation()

	return admins, err
}

func (s *adminService) Create(adminRequest request.CreateAdminRequest) (model.Admin, error) {
	var admin = model.Admin{
		UserId:      adminRequest.UserId,
		NamaLengkap: adminRequest.NamaLengkap,
	}

	var newAdmin, err = s.adminRepository.Create(admin)

	return newAdmin, err
}

func (s *adminService) Update(id int, adminrequest request.UpdateAdminRequest) (model.Admin, error) {
	var admin, err = s.adminRepository.FindById(id)

	admin.UserId = adminrequest.UserId
	admin.NamaLengkap = adminrequest.NamaLengkap

	updatedAdmin, err := s.adminRepository.Update(admin)

	return updatedAdmin, err
}

func (s *adminService) Delete(id int) (model.Admin, error) {
	var instansi, err = s.adminRepository.FindById(id)

	deletedInstansi, err := s.adminRepository.Delete(instansi)

	return deletedInstansi, err
}
