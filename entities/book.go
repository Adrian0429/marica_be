package entities

import "github.com/google/uuid"

type (
	Book struct {
		ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
		Title     string    `gorm:"type:varchar(255);" json:"Title"`
		Thumbnail string    `json:"thumbnail_path"`
		Pages     []Pages   `json:"Pages,omitempty"`
	}

	Pages struct {
		ID       uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
		Filename string    `json:"filename"`
		Path     string    `json:"path"`

		BookID uuid.UUID `gorm:"type:uuid" json:"-"`
		Book   Book      `gorm:"foreignKey:BookID" json:"-"`
		
	}
)
