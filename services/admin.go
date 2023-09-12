package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/helpers"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/google/uuid"
)

type AdminService interface {
	RegisterAdmin()
	VerifyLogin(ctx context.Context, adminDTO dto.AdminLoginDTO) (bool, error)
	CheckAdminByEmail(ctx context.Context, email string) (entities.User, error)
	GetAdminByID(ctx context.Context, adminID uuid.UUID) (entities.User, error)
}

type adminService struct {
	adminRepository repository.UserRepository
}

func NewAdminService(ar repository.UserRepository) AdminService {
	return &adminService{
		adminRepository: ar,
	}
}

func (as *adminService) RegisterAdmin() {

}

func (as *adminService) CheckAdminByEmail(ctx context.Context, email string) (entities.User, error) {
	admin, err := as.adminRepository.GetAdminByEmail(ctx, email)
	if err != nil {
		return entities.User{}, err
	}

	return admin, nil
}

func (as *adminService) GetAdminByID(ctx context.Context, adminID uuid.UUID) (entities.User, error) {
	admin, err := as.adminRepository.GetAdminByID(ctx, adminID)
	if err != nil {
		return entities.User{}, dto.ErrNotAdminID
	}

	return admin, nil
}

func (as *adminService) VerifyLogin(ctx context.Context, adminDTO dto.AdminLoginDTO) (bool, error) {
	admin, err := as.adminRepository.GetAdminByEmail(ctx, adminDTO.Email)
	if err != nil {
		return false, dto.ErrEmailNotFound
	}

	if admin.Email == "" {
		return false, dto.ErrEmailNotFound
	}

	if checkPassword, _ := helpers.CheckPassword(admin.Password, []byte(adminDTO.Password)); !checkPassword {
		return false, dto.ErrPasswordNotMatch
	}

	return true, nil
}
