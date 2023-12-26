package entities

import "github.com/google/uuid"

type (
	Book struct {
		ID         uuid.UUID `gorm:"type:uuid;primary_key"`
		Title      string    `gorm:"type:varchar(255);" json:"title"`
		Desc       string    `gorm:"type:varchar(255);" json:"description"`
		Thumbnail  string    `json:"thumbnail"`
		Pages      []Pages   `json:"Pages,omitempty" gorm:"onDelete:CASCADE"`
		View       int       `json:"View_Count,omitempty"`
		Page_Count int       `json:"Page_Count,omitempty"`
		UserID     uuid.UUID `gorm:"type:uuid" json:"-"`
		User       User      `gorm:"foreignKey:UserID" json:"-"`
	}

	Pages struct {
		ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
		PageTitle string    `gorm:"type:string" json:"page_title"`
		Index     int       `gorm:"type:integer" json:"index"`
		BookID    uuid.UUID `gorm:"type:uuid" json:"-"`
		Book      Book      `gorm:"foreignKey:BookID" json:"-"`

		Files []Files `json:"Files,omitempty" gorm:"onDelete:CASCADE"`
	}

	Files struct {
		ID    uuid.UUID `gorm:"type:uuid;primary_key" json:"id" `
		Path  string    `gorm:"type:varchar(255)" json:"path"`
		Index int       `gorm:"type:integer" json:"index"`

		PagesID uuid.UUID `gorm:"type:uuid" json:"-"`
		Pages   Pages     `gorm:"foreignKey:PagesID" json:"-"`
	}
)
