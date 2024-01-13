package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, user entities.User) (entities.User, error)
	GetAllUser(ctx context.Context) ([]entities.User, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (entities.User, error)
	GetUserByEmail(ctx context.Context, email string) (entities.User, error)
	UpdateUser(ctx context.Context, user entities.User) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	UpdateUserPassword(ctx context.Context, userID string, newPassword string) error
	GetAdminByEmail(ctx context.Context, email string) (entities.User, error)
	GetAdminByID(ctx context.Context, adminID uuid.UUID) (entities.User, error)
	GiveAccess(ctx context.Context, access entities.Book_User) (entities.Book_User, error)
	RemoveAccess(ctx context.Context, access entities.Book_User) error
	FindAccess(ctx context.Context, access entities.Book_User) (entities.Book_User, error)
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		connection: db,
	}
}

func (ur *userRepository) RegisterUser(ctx context.Context, user entities.User) (entities.User, error) {
	if err := ur.connection.Create(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (ur *userRepository) GetAllUser(ctx context.Context) ([]entities.User, error) {
	var users []entities.User
	if err := ur.connection.Not("name = ?", "Admin Sebangku").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) GetUserByID(ctx context.Context, userID uuid.UUID) (entities.User, error) {
	var user entities.User
	if err := ur.connection.Where("id = ?", userID).Take(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (ur *userRepository) GetUserByEmail(ctx context.Context, email string) (entities.User, error) {
	var user entities.User
	if err := ur.connection.Where("email = ?", email).Take(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, user entities.User) error {
	if err := ur.connection.Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) UpdateUserPassword(ctx context.Context, userID string, newPassword string) error {
	user := entities.User{
		Password: newPassword,
	}
	if err := ur.connection.Model(&entities.User{}).Where("id = ?", userID).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	if err := ur.connection.Delete(&entities.User{}, &userID).Error; err != nil {
		return err
	}
	return nil
}

func (ar *userRepository) GetAdminByEmail(ctx context.Context, email string) (entities.User, error) {
	var admin entities.User
	if err := ar.connection.Where("email = ?", email).Take(&admin).Error; err != nil {
		return entities.User{}, err
	}

	return admin, nil
}

func (ar *userRepository) GetAdminByID(ctx context.Context, adminID uuid.UUID) (entities.User, error) {
	var admin entities.User
	if err := ar.connection.Where("id = ?", adminID).Take(&admin).Error; err != nil {
		return entities.User{}, err
	}

	return admin, nil
}

func (ar *userRepository) GiveAccess(ctx context.Context, access entities.Book_User) (entities.Book_User, error) {
	if err := ar.connection.Create(&access).Error; err != nil {
		return entities.Book_User{}, err
	}
	return access, nil
}

func (ar *userRepository) FindAccess(ctx context.Context, access entities.Book_User) (entities.Book_User, error) {
	bookid := access.BookID
	userid := access.UserID

	result := entities.Book_User{}

	if err := ar.connection.Where("book_id = ? AND user_id = ?", bookid, userid).First(&result).Error; err != nil {
		return entities.Book_User{}, err
	}

	return result, nil
}

func (ur *userRepository) RemoveAccess(ctx context.Context, access entities.Book_User) error {
	if err := ur.connection.Where(&entities.Book_User{UserID: access.UserID, BookID: access.BookID}).Delete(&entities.Book_User{}).Error; err != nil {
		return err
	}
	return nil
}
