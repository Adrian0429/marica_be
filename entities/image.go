package entities

import "github.com/google/uuid"

type Image struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Path       string    `gorm:"type:varchar(100)" json:"path"`
	FileName   string    `gorm:"type:varchar(100)" json:"file_name"`
	Base64Data string    `gorm:"type:varchar(100)" json:"base64Data"`
}
