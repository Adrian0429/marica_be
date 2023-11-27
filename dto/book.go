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
)

type (
	BookCreateRequest struct {
		Title        string                `form:"title" json:"title"`
		Desc         string                `form:"description" json:"description"`
		Thumbnail    *multipart.FileHeader `form:"thumbnail" json:"thumbnail"`
		UserID       uuid.UUID             `json:"user_id"`
		MediaRequest []MediaRequest        `json:"medias"`
	}

	BookCreateResponse struct {
		ID               string             `json:"id"`
		Title            string             `json:"title"`
		Desc             string             `json:"description"`
		Thumbnail        string             `json:"thumbnail_path"`
		MediaPathRequest []MediaPathRequest `json:"medias"`
	}

	MediaPathRequest struct {
		Index int    `json:"index"`
		Page  int    `json:"page"`
		Media string `json:"media"`
	}

	MediaRequest struct {
		Index int                   `json:"index"`
		Page  int                   `json:"page"`
		Media *multipart.FileHeader `json:"media"`
	}

	PagesResponse struct {
		ID       string `json:"id"`
		Pages    string `json:"pages"`
		Filename string `json:"filename"`
		Path     string `json:"path"`
	}

	BookPagesRequest struct {
		Path string `json:"pages_paths"`
	}
)
