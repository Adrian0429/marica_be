package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetBookByTitle(ctx context.Context, title string) (entities.Book, error)
	CreateBook(ctx context.Context, book entities.Book) (entities.Book, error)
	GetAllBooks(ctx context.Context) ([]entities.Book, error)
	GetBookPages(ctx context.Context, bookID string) ([]entities.Pages, error)
	GetBookByID(ctx context.Context, title string) (entities.Book, error)
}

type bookRepository struct {
	connection *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		connection: db,
	}
}

func (br *bookRepository) GetBookByTitle(ctx context.Context, title string) (entities.Book, error) {
	var book entities.Book
	if err := br.connection.Where("title = ?", title).First(&book).Error; err != nil {
		return entities.Book{}, err
	}
	return book, nil
}

func (br *bookRepository) GetBookByID(ctx context.Context, title string) (entities.Book, error) {
	var book entities.Book
	if err := br.connection.Where("id = ?", title).First(&book).Error; err != nil {
		return entities.Book{}, err
	}
	return book, nil
}

func (br *bookRepository) CreateBook(ctx context.Context, book entities.Book) (entities.Book, error) {
	if err := br.connection.Create(&book).Error; err != nil {
		return entities.Book{}, err
	}
	return book, nil
}

func (br *bookRepository) GetAllBooks(ctx context.Context) ([]entities.Book, error) {
	var book []entities.Book
	if err := br.connection.Find(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (br *bookRepository) GetBookPages(ctx context.Context, bookID string) ([]entities.Pages, error) {
	var pages []entities.Pages
	if err := br.connection.Where("book_id = ?", bookID).Find(&pages).Error; err != nil {
		return []entities.Pages{}, err
	}

	return pages, nil
}
