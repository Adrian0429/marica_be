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
	PATH           = "storage/Pages"
	THUMBNAIL_PATH = "Thumbnail"
)

type BookService interface {
	CreateBook(ctx context.Context, req dto.BookCreateRequest) (dto.BookCreateResponse, error)
	GetAllBooks(ctx context.Context) ([]dto.BookCreateResponse, error)
	GetBookPages(ctx context.Context, bookID string) ([]dto.BookPagesRequest, error)
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
	existingBook, err := bs.br.GetBookByTitle(ctx, req.Title)
	if err != nil {
		return dto.BookCreateResponse{}, dto.ErrGetBookByTitle
	}

	if existingBook.Title != "" {
		return dto.BookCreateResponse{}, dto.ErrDuplicateTitle
	}

	bookId := uuid.New()

	var thumbnailPath string

	if req.Thumbnail != nil {
		thumbnailFilename := utils.Getextension(req.Thumbnail.Filename)
		thumbnailPath = utils.GenerateFileName("Thumbnail", req.Title, thumbnailFilename)
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
		var imagePath string

		if v.Pages != nil {
			bookPages, err := utils.IsBase64(*v.Pages)
			if err != nil {
				return dto.BookCreateResponse{}, dto.ErrToBase64
			}

			imageId := uuid.New()

			bookPagesSave := imageId.String() + utils.Getextension(v.Pages.Filename)
			imagePath = utils.GenerateFileName("Pages", createdBook.Title, bookPagesSave)
			_ = utils.SaveImage(bookPages, PATH, createdBook.Title, bookPagesSave)

		}

		pagesItem := entities.Pages{
			ID:       uuid.New(),
			Path:     imagePath,
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
			Thumbnail: book.Thumbnail,
		}
		allBooks = append(allBooks, bookProps)
	}
	return allBooks, nil
}

func (bc *bookService) GetBookPages(ctx context.Context, bookID string) ([]dto.BookPagesRequest, error) {

	Book, err := bc.br.GetBookPages(ctx, bookID)
	if err != nil {
		return []dto.BookPagesRequest{}, err
	}

	var allPages []dto.BookPagesRequest
	for _, page := range Book {
		pageProps := dto.BookPagesRequest{
			Path: page.Path,
		}
		allPages = append(allPages, pageProps)
	}
	return allPages, nil

}
