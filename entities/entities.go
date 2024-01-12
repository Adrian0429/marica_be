package entities

import (
	"github.com/Caknoooo/golang-clean_template/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	User struct {
		ID         uuid.UUID `gorm:"type:uuid;primary_key;;default:uuid_generate_v4()" json:"id"`
		Name       string    `gorm:"type:varchar(100)" json:"name"`
		Email      string    `gorm:"type:varchar(100)" json:"email"`
		TelpNumber string    `gorm:"type:varchar(14)" json:"telp_number"`
		Password   string    `gorm:"type:varchar(100)" json:"password"`
		Role       string    `gorm:"type:varchar(100)" json:"role"`

		Timestamp
	}

	Book_User struct {
		UserID string `gorm:"type:varchar(255);" json:"user_id"`
		BookID string `gorm:"type:varchar(255);" json:"book_id"`
	}

	Book struct {
		ID          uuid.UUID `gorm:"type:uuid;primary_key;;default:uuid_generate_v4()" json:"id"`
		Title       string    `gorm:"type:varchar(255);" json:"title"`
		Tags        string    `gorm:"type:varchar(128);" json:"tags"`
		Desc        string    `gorm:"type:varchar(255);" json:"description"`
		Thumbnail   string    `json:"thumbnail"`
		Pages       []Pages   `json:"Pages,omitempty" gorm:"onDelete:CASCADE"`
		View        int       `json:"View_Count,omitempty"`
		Page_Count  int       `json:"Page_Count,omitempty"`
		Tokped_Link string    `json:"Tokped_Link,omitempty"`
	}

	Pages struct {
		ID        uuid.UUID `gorm:"type:uuid;primary_key;;default:uuid_generate_v4()" json:"id"`
		PageTitle string    `gorm:"type:string" json:"page_title"`
		Index     int       `gorm:"type:integer" json:"index"`
		BookID    uuid.UUID `gorm:"type:uuid" json:"-"`
		Book      Book      `gorm:"foreignKey:BookID" json:"-"`

		Files         []Files     `json:"Files,omitempty" gorm:"onDelete:CASCADE"`
		Iframes       []Iframes   `json:"Iframes,omitempty" gorm:"onDelete:CASCADE"`
		LiveWorksheet []Worksheet `json:"LiveWorksheet,omitempty" gorm:"onDelete:CASCADE"`
	}

	Files struct {
		ID    uuid.UUID `gorm:"type:uuid;primary_key;;default:uuid_generate_v4()" json:"id"`
		Path  string    `gorm:"type:varchar(255)" json:"path"`
		Index int       `gorm:"type:integer" json:"index"`

		PagesID uuid.UUID `gorm:"type:uuid" json:"-"`
		Pages   Pages     `gorm:"foreignKey:PagesID" json:"-"`
	}

	Iframes struct {
		ID    uuid.UUID `gorm:"type:uuid;primary_key;;default:uuid_generate_v4()" json:"id"`
		Path  string    `gorm:"type:varchar(255)" json:"path"`
		Index int       `gorm:"type:integer" json:"index"`

		PagesID uuid.UUID `gorm:"type:uuid" json:"-"`
		Pages   Pages     `gorm:"foreignKey:PagesID" json:"-"`
	}

	Worksheet struct {
		ID            uuid.UUID `gorm:"type:uuid;primary_key;;default:uuid_generate_v4()" json:"id"`
		Worksheet_ID  int       `gorm:"type:integer" json:"worksheet_id"`
		String_Code   string    `gorm:"type:varchar(255)" json:"String_Code"`
		Worksheet_ID2 int       `gorm:"type:integer" json:"worksheet_id2"`
		String_Code2  string    `gorm:"type:varchar(255)" json:"String_Code2"`
		String_Code3  string    `gorm:"type:varchar(255)" json:"String_Code3"`

		PagesID uuid.UUID `gorm:"type:uuid" json:"-"`
		Pages   Pages     `gorm:"foreignKey:PagesID" json:"-"`
	}
)

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}
