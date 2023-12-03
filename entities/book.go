package entities

import "github.com/google/uuid"

type (
	Book struct {
		ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
		Title     string    `gorm:"type:varchar(255);" json:"title"`
		Desc      string    `gorm:"type:varchar(255);" json:"description"`
		Thumbnail string    `json:"thumbnail_path"`
		Pages     []Pages   `json:"Pages,omitempty"`
		View      int       `json:"View_Count,omitempty"`
		UserID    uuid.UUID `gorm:"type:uuid" json:"-"`
		User      User      `gorm:"foreignKey:UserID" json:"-"`
	}

	Pages struct {
		ID       uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
		Index    int       `gorm:"typeinteger" json:"index"`
		Page     int       `gorm:"typeinteger" json:"page"`
		Path     string    `gorm:"type:varchar(255)" json:"path"`
		FileName string    `gorm:"type:varchar(255)" json:"file_name"`

		BookID uuid.UUID `gorm:"type:uuid" json:"-"`
		Book   Book      `gorm:"foreignKey:BookID" json:"-"`
	}
)
