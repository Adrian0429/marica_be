package entities

import (
	"github.com/Caknoooo/golang-clean_template/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Name       string    `gorm:"type:varchar(100)" json:"name"`
	Email      string    `gorm:"type:varchar(100)" json:"email"`
	TelpNumber string    `gorm:"type:varchar(14)" json:"telp_number"`
	Password   string    `gorm:"type:varchar(100)" json:"password"`
	Role       string    `gorm:"type:varchar(100)" json:"role"`

	Timestamp
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}
