package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type (
	PasswordRepository interface {
		CreatePassword(ctx context.Context, pass entities.Password) (entities.Password, error)
		GetToken(ctx context.Context, token string) (entities.Password, error)
	}

	passwordRepository struct {
		db *gorm.DB
	}
)

func NewPasswordRepository(db *gorm.DB) *passwordRepository {
	return &passwordRepository{
		db: db,
	}
}

func (r *passwordRepository) CreatePassword(ctx context.Context, password entities.Password) (entities.Password, error) {
	if err := r.db.Create(&password).Error; err != nil {
		return entities.Password{}, err
	}

	return password, nil
}

func (r *passwordRepository) GetToken(ctx context.Context, token string) (entities.Password, error) {
	var pass entities.Password
	if err := r.db.Where("token = ?", token).Take(&pass).Error; err != nil {
		return entities.Password{}, err
	}
	return pass, nil
}
