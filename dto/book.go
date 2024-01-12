package dto

import (
	"errors"
	"mime/multipart"

	"github.com/google/uuid"
)

var (
	ErrCreateBooks    = errors.New("failed to create new books")
	ErrGetAllBooks    = errors.New("failed to fetch all Books")
	ErrCreatePages    = errors.New("failed to create pages")
	ErrAddAudio       = errors.New("failed to save audio")
	ErrNull           = errors.New("nothing")
	ErrDuplicateTitle = errors.New("duplicate title")
	ErrGetBookByTitle = errors.New("failed to get book by title")
	ErrGetBookByID    = errors.New("failed to get book by ID")
)

type (
	BookCreateRequest struct {
		Title        string                `form:"title" json:"title"`
		Desc         string                `form:"description" json:"description"`
		Thumbnail    *multipart.FileHeader `form:"thumbnail" json:"thumbnail"`
		UserID       uuid.UUID             `json:"user_id"`
		Page_Count   int                   `json:"page_count"`
		Tags         string                `json:"tags"`
		Tokped_Link  string                `json:"tokped_link"`
		MediaRequest []MediaRequest        `json:"medias"`
	}

	MediaRequest struct {
		Index     int         `json:"index"`
		Title     string      `json:"page_title"`
		Files     []Files     `json:"files"`
		IFrames   []IFrames   `json:"iframe"`
		Worksheet []Worksheet `json:"worksheet"`
	}

	IFrames struct {
		Index  int    `json:"index"`
		Iframe string `json:"iframe"`
	}

	Worksheet struct {
		Worksheet_ID  int    `gorm:"type:integer" json:"worksheet_id"`
		String_Code   string `gorm:"type:varchar(255)" json:"String_Code"`
		Worksheet_ID2 int    `gorm:"type:integer" json:"worksheet_id2"`
		String_Code2  string `gorm:"type:varchar(255)" json:"String_Code2"`
		String_Code3  string `gorm:"type:varchar(255)" json:"String_Code3"`
	}

	IframePaths struct {
		Iframe string `json:"iframe"`
	}

	Files struct {
		Index  int                   `json:"index"`
		Images *multipart.FileHeader `json:"images"`
	}

	AllPagesRequest struct {
		Title         string          `form:"title" json:"title"`
		Desc          string          `form:"description" json:"description"`
		Thumbnail     string          `form:"thumbnail" json:"thumbnail"`
		Page_Count    int             `json:"page_count"`
		Tags          string          `json:"tags"`
		AllPagesMedia []AllPagesMedia `json:"medias"`
	}

	AllPagesMedia struct {
		Index int             `json:"index"`
		Title string          `json:"page_title"`
		Files []AllPagesFiles `json:"files"`
	}

	AllPagesFiles struct {
		Index int    `json:"index"`
		Path  string `json:"images"`
	}

	BooksRequest struct {
		ID        string `json:"id"`
		Title     string `json:"title"`
		Tags      string `json:"tags"`
		Desc      string `json:"description"`
		Thumbnail string `json:"thumbnail_path"`
	}

	BookCreateResponse struct {
		ID               string             `json:"id"`
		Title            string             `json:"title"`
		Desc             string             `json:"description"`
		Tags             string             `json:"tags"`
		Thumbnail        string             `json:"thumbnail"`
		Page_Count       int                `json:"page_count"`
		Tokped_Link      string             `json:"tokped_link"`
		MediaPathRequest []MediaPathRequest `json:"medias"`
	}

	MediaPathRequest struct {
		Index     int         `json:"index"`
		Iframes   []IFrames   `json:"iframe"`
		Media     []Medias    `json:"media"`
		Worksheet []Worksheet `json:"worksheet"`
	}

	Medias struct {
		Index int    `json:"index"`
		Path  string `json:"path"`
	}

	PagePaths struct {
		Path string `json:"pages_paths"`
	}

	BookPageRequest struct {
		BookID     string        `json:"id"`
		Title      string        `json:"title"`
		Tags       string        `json:"tags"`
		PageTitle  string        `json:"page_title"`
		Page_Count int           `json:"page_count"`
		PagePaths  []PagePaths   `json:"page_paths"`
		Iframe     []IframePaths `json:"iframe_paths"`
		Worksheet  []Worksheet   `json:"worksheets"`
	}

	BookPreviewRequest struct {
		BookID      string      `json:"id"`
		Title       string      `json:"title"`
		Tags        string      `json:"tags"`
		Desc        string      `json:"description"`
		Thumbnail   string      `json:"thumbnail"`
		PageTitle   string      `json:"page_title"`
		Page_Count  int         `json:"page_count"`
		Tokped_Link string      `json:"tokped_link"`
		PagePaths   []PagePaths `json:"page_paths"`
	}
)
