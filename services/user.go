package services

import (
	"context"
	"time"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/helpers"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type UserService interface {
	RegisterUser(ctx context.Context, userDTO dto.UserCreateRequest) (entities.User, error)
	GetAllUser(ctx context.Context) ([]dto.UserRequest, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (entities.User, error)
	GetUserByEmail(ctx context.Context, email string) (entities.User, error)
	ForgotPassword(ctx context.Context, user entities.User) (bool, error)
	UpdateUserPassword(ctx context.Context, userID string, newPassword string) error
	CheckUser(ctx context.Context, email string) (bool, error)
	CheckToken(ctx context.Context, token string) (entities.Password, error)
	UpdateUser(ctx context.Context, userDTO dto.UserUpdateRequest) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	Verify(ctx context.Context, email string, password string) (bool, error)
}

type userService struct {
	userRepository     repository.UserRepository
	passwordRepository repository.PasswordRepository
}

func NewUserService(ur repository.UserRepository, pr repository.PasswordRepository) UserService {
	return &userService{
		userRepository:     ur,
		passwordRepository: pr,
	}
}

func (us *userService) RegisterUser(ctx context.Context, userDTO dto.UserCreateRequest) (entities.User, error) {
	user := entities.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(userDTO))
	user.Role = helpers.USER
	user.ID = uuid.New()
	if err != nil {
		return entities.User{}, err
	}
	return us.userRepository.RegisterUser(ctx, user)
}

func (us *userService) GetAllUser(ctx context.Context) ([]dto.UserRequest, error) {
	user, err := us.userRepository.GetAllUser(ctx)
	res := []dto.UserRequest{}
	if err != nil {
		return []dto.UserRequest{}, err
	}
	for _, users := range user {
		userlist := dto.UserRequest{
			ID:    users.ID.String(),
			Name:  users.Name,
			Email: users.Email,
		}
		res = append(res, userlist)
	}

	return res, nil
}

func (us *userService) GetUserByID(ctx context.Context, userID uuid.UUID) (entities.User, error) {
	return us.userRepository.GetUserByID(ctx, userID)
}

func (us *userService) GetUserByEmail(ctx context.Context, email string) (entities.User, error) {
	return us.userRepository.GetUserByEmail(ctx, email)
}

func (us *userService) CheckUser(ctx context.Context, email string) (bool, error) {
	res, err := us.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	if res.Email == "" {
		return false, err
	}
	return true, nil
}

func (us *userService) CheckToken(ctx context.Context, token string) (entities.Password, error) {
	res, err := us.passwordRepository.GetToken(ctx, token)
	if err != nil {
		return entities.Password{}, err
	}

	if time.Now().After(res.Expiry) {
		return entities.Password{}, err
	}
	return res, nil
}

func (us *userService) ForgotPassword(ctx context.Context, user entities.User) (bool, error) {
	expiry := time.Now().Add(5 * time.Minute)
	token := utils.GenerateOTP()

	forgot := entities.Password{
		ID:     uuid.New(),
		UserID: user.ID,
		Token:  token,
		Expiry: expiry,
	}

	res, err := us.passwordRepository.CreatePassword(ctx, forgot)
	if err != nil {
		return false, err
	}
	if res.UserID == uuid.Nil {
		return false, err
	}
	utils.SendRequestEmail(user.Email, res.Token)

	return true, nil
}

func (us *userService) UpdateUserPassword(ctx context.Context, userID string, newPassword string) error {
	newpass, _ := helpers.HashPassword(newPassword)
	return us.userRepository.UpdateUserPassword(ctx, userID, newpass)
}

func (us *userService) UpdateUser(ctx context.Context, userDTO dto.UserUpdateRequest) error {
	user := entities.User{}
	if err := smapping.FillStruct(&user, smapping.MapFields(userDTO)); err != nil {
		return nil
	}
	return us.userRepository.UpdateUser(ctx, user)
}

func (us *userService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return us.userRepository.DeleteUser(ctx, userID)
}

func (us *userService) Verify(ctx context.Context, email string, password string) (bool, error) {
	res, err := us.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	checkPassword, err := helpers.CheckPassword(res.Password, []byte(password))
	if err != nil {
		return false, err
	}

	if res.Email == email && checkPassword {
		return true, nil
	}
	return false, nil
}
