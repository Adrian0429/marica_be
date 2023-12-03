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
	GetTopBooks(ctx context.Context) ([]entities.Book, error)
	GetBookPages(ctx context.Context, bookID string, bookPage string) ([]entities.Pages, error)
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

func (br *bookRepository) GetBookByID(ctx context.Context, id string) (entities.Book, error) {
	var book entities.Book
	if err := br.connection.Where("id = ?", id).First(&book).Error; err != nil {
		return entities.Book{}, err
	}
	if err := br.connection.Model(&book).Update("View", book.View+1).Error; err != nil {
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

func (br *bookRepository) GetTopBooks(ctx context.Context) ([]entities.Book, error) {
	var books []entities.Book
	if err := br.connection.Order("View desc").Limit(10).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (br *bookRepository) GetBookPages(ctx context.Context, bookID string, bookPage string) ([]entities.Pages, error) {
	var pages []entities.Pages
	if err := br.connection.Where("book_id = ? AND index= ?", bookID, bookPage).Find(&pages).Error; err != nil {
		return []entities.Pages{}, err
	}

	return pages, nil
}
