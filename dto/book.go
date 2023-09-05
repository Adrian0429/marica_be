package dto

import (
	"errors"
	"mime/multipart"
)

var (
	ErrCreateBooks = errors.New("failed to create new books")
	ErrBookTitle   = errors.New("failed to fetch book titles")
	ErrCreatePages = errors.New("failed to create pages")
)

type (
	BookCreateRequest struct {
		Title        string                `form:"title" json:"Title"`
		Thumbnail    *multipart.FileHeader `form:"thumbnail" json:"Thumbnail"`
		PagesRequest []PagesRequest        `json:"Pages"`
	}

	BookTitleRequest struct {
		Title string `json:"Title"`
	}

	BookCreateResponse struct {
		ID        string `json:"id"`
		Title     string `json:"Book_Title"`
		Thumbnail string `json:"Thumbnail_Path"`
	}

	PagesRequest struct {
		Pages *multipart.FileHeader `json:"pages"`
	}

	PagesResponse struct {
		ID       string `json:"id"`
		Pages    string `json:"pages"`
		Filename string `json:"filename"`
		Path     string `json:"path"`
	}
)
