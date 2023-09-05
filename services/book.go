package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/google/uuid"
)

const (
	PATH = "Storage/Pages"
)

type BookService interface {
	CreateBook(ctx context.Context, req dto.BookCreateRequest) (dto.BookCreateResponse, error)
	GetAllBookTitle(ctx context.Context) ([]dto.BookTitleRequest, error)
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

	var thumbnailPath string

	if req.Thumbnail != nil {
		thumbnailFilename := utils.Getextension(req.Thumbnail.Filename)
		thumbnailPath = utils.GenerateFileName(PATH, req.Title, thumbnailFilename)

		thumbnailData, err := utils.IsBase64(*req.Thumbnail)
		if err != nil {
			return dto.BookCreateResponse{}, err
		}
		err = utils.SaveImage(thumbnailData, "storage/Thumbnail", req.Title, thumbnailFilename)
	}

	book := entities.Book{
		ID:        bookId,
		Title:     req.Title,
		Thumbnail: thumbnailPath,
	}

	//create book
	createdBook, err := bs.br.CreateBook(ctx, book)
	if err != nil {
		return dto.BookCreateResponse{}, dto.ErrCreateBooks
	}

	for _, v := range req.PagesRequest {
		var imageName string

		if v.Pages != nil {
			bookPages, err := utils.IsBase64(*v.Pages)
			if err != nil {
				return dto.BookCreateResponse{}, dto.ErrToBase64
			}

			imageId := uuid.New()

			bookPagesSave := imageId.String() + utils.Getextension(v.Pages.Filename)
			_ = utils.SaveImage(bookPages, PATH, createdBook.Title, bookPagesSave)
			imageName = utils.GenerateFileName(PATH, createdBook.Title, bookPagesSave)
		}

		pagesItem := entities.Pages{
			ID:       uuid.New(),
			Path:     imageName,
			Filename: v.Pages.Filename,
			BookID:   bookId,
		}

		_, err := bs.pr.CreatePages(ctx, pagesItem)
		if err != nil {
			return dto.BookCreateResponse{}, dto.ErrCreatePages
		}

	}
	return dto.BookCreateResponse{
		ID:        createdBook.ID.String(),
		Title:     createdBook.Title,
		Thumbnail: createdBook.Thumbnail,
	}, nil

}

func (bs *bookService) GetAllBookTitle(ctx context.Context) ([]dto.BookTitleRequest, error) {
	Titles, err := bs.br.GetAllBookTitle(ctx)
	if err != nil {
		return nil, dto.ErrBookTitle
	}

	var bookTitleRequests []dto.BookTitleRequest
	for _, title := range Titles {
		bookTitle := dto.BookTitleRequest{
			Title: title.Title,
		}
		bookTitleRequests = append(bookTitleRequests, bookTitle)
	}
	return bookTitleRequests, nil
}
