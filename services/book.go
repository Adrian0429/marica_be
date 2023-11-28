package services

import (
	"context"
	"strconv"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/google/uuid"
)

const (
	PATH       = "storage/"
	PAGES_PATH = "storage/pages"
)

type BookService interface {
	CreateBook(ctx context.Context, req dto.BookCreateRequest) (dto.BookCreateResponse, error)
	GetAllBooks(ctx context.Context) ([]dto.BookCreateResponse, error)
	GetBookPages(ctx context.Context, bookID string) (dto.BookCreateResponse, error)
	CheckTitle(ctx context.Context, Title string) (bool, error)
}

type bookService struct {
	br repository.BookRepository
	pr repository.PagesRepository
}

func NewBookService(br repository.BookRepository, pr repository.PagesRepository) *bookService {
	return &bookService{
		br: br,
		pr: pr,
	}
}

func (bs *bookService) CreateBook(ctx context.Context, req dto.BookCreateRequest) (dto.BookCreateResponse, error) {
	bookId := uuid.New()
	var resBooks dto.BookCreateResponse
	var mediaRequests []dto.MediaPathRequest
	var thumbnailPath string
	if req.Thumbnail != nil {
		thumbnailData, err := utils.IsBase64(*req.Thumbnail)
		if err != nil {
			return dto.BookCreateResponse{}, err
		}

		thumbnailFilename := utils.Getextension(req.Thumbnail.Filename)
		thumbnailPath = utils.GenerateFileName(PATH, req.Title+"/thumbnail", req.Title+thumbnailFilename)

		err = utils.UploadThumbnail(thumbnailData, PATH, req.Title, thumbnailFilename)
		if err != nil {
			return dto.BookCreateResponse{}, err
		}
	}

	book := entities.Book{
		ID:        bookId,
		Desc:      req.Desc,
		Title:     req.Title,
		UserID:    req.UserID,
		Thumbnail: thumbnailPath,
	}
	resBooks.ID = bookId.String()
	resBooks.Desc = req.Desc
	resBooks.Thumbnail = thumbnailPath
	resBooks.Title = req.Title

	createdBook, err := bs.br.CreateBook(ctx, book)
	if err != nil {
		return dto.BookCreateResponse{}, dto.ErrCreateBooks
	}

	for _, v := range req.MediaRequest {
		var mediaPath string

		if v.Media != nil {
			bookPages, err := utils.IsBase64(*v.Media)
			if err != nil {
				return dto.BookCreateResponse{}, dto.ErrToBase64
			}

			mediaPath = utils.GenerateFileName(PATH+createdBook.Title, "/Pages_"+strconv.Itoa(v.Index), v.Media.Filename)
			_ = utils.Upload(bookPages, PATH+createdBook.Title, "/Pages_"+strconv.Itoa(v.Index), v.Media.Filename)
		}

		pagesItem := entities.Pages{
			ID:       uuid.New(),
			Index:    v.Index,
			Page:     v.Page,
			Path:     mediaPath,
			FileName: v.Media.Filename,
			BookID:   createdBook.ID,
		}

		var medias dto.MediaPathRequest
		medias.Media = mediaPath
		medias.Index = v.Index
		medias.Page = v.Page
		mediaRequests = append(mediaRequests, medias)

		_, err := bs.pr.CreatePages(ctx, pagesItem)
		if err != nil {
			return dto.BookCreateResponse{}, dto.ErrCreatePages
		}
	}

	resBooks.MediaPathRequest = mediaRequests
	return resBooks, nil
}

func (bs *bookService) GetAllBooks(ctx context.Context) ([]dto.BookCreateResponse, error) {
	Books, err := bs.br.GetAllBooks(ctx)
	if err != nil {
		return nil, dto.ErrGetAllBooks
	}
	var allBooks []dto.BookCreateResponse
	for _, book := range Books {
		bookProps := dto.BookCreateResponse{
			ID:        book.ID.String(),
			Title:     book.Title,
			Desc:      book.Desc,
			Thumbnail: book.Thumbnail,
		}

		allBooks = append(allBooks, bookProps)
	}
	return allBooks, nil
}

func (bc *bookService) GetBookPages(ctx context.Context, bookID string) (dto.BookCreateResponse, error) {
	var mediaRequests []dto.MediaPathRequest
	books, err := bc.br.GetBookByID(ctx, bookID)
	if err != nil {
		return dto.BookCreateResponse{}, err
	}

	resBooks := dto.BookCreateResponse{
		ID:        books.ID.String(),
		Title:     books.Title,
		Desc:      books.Desc,
		Thumbnail: books.Thumbnail,
	}

	bookPages, err := bc.br.GetBookPages(ctx, bookID)
	if err != nil {
		return dto.BookCreateResponse{}, err
	}

	for _, page := range bookPages {
		pages := dto.MediaPathRequest{
			Index: page.Index,
			Page:  page.Page,
			Media: page.Path,
		}
		mediaRequests = append(mediaRequests, pages)
	}
	resBooks.MediaPathRequest = mediaRequests
	return resBooks, nil
}

func (bc *bookService) CheckTitle(ctx context.Context, Title string) (bool, error) {
	title, err := bc.br.GetBookByTitle(ctx, Title)
	if err != nil {
		return false, err
	}

	if title.Title == "" {
		return false, err
	}

	return true, nil
}
