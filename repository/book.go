package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetBookByTitle(ctx context.Context, title string) (entities.Book, error)
	CreateBook(ctx context.Context, book entities.Book) (entities.Book, error)
	GetAllBooks(ctx context.Context) ([]entities.Book, error)
	GetUserBooksID(ctx context.Context, userID uuid.UUID) ([]entities.Book_User, error)
	GetTopBooks(ctx context.Context) ([]entities.Book, error)
	GetBookPage(ctx context.Context, bookID string, bookPage string) (entities.Pages, error)
	GetBookAllPages(ctx context.Context, bookID string) ([]entities.Pages, error)
	GetBookByID(ctx context.Context, title string) (entities.Book, error)
	GetPagesPaths(ctx context.Context, pagesID string) ([]entities.Files, error)
	GetPagesIframe(ctx context.Context, pagesID string) ([]entities.Iframes, error)
	GetPagesWorksheets(ctx context.Context, pagesID string) ([]entities.Worksheet, error)
	DeleteBooks(ctx context.Context, BookID string) error
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

func (br *bookRepository) GetBookPage(ctx context.Context, bookID string, bookPage string) (entities.Pages, error) {
	var pages entities.Pages
	if err := br.connection.Where("book_id = ? AND index= ?", bookID, bookPage).Find(&pages).Error; err != nil {
		return entities.Pages{}, err
	}

	return pages, nil
}

func (br *bookRepository) GetBookAllPages(ctx context.Context, bookID string) ([]entities.Pages, error) {
	var pages []entities.Pages
	if err := br.connection.Where("book_id = ?", bookID).Find(&pages).Error; err != nil {
		return []entities.Pages{}, err
	}

	return pages, nil
}

func (br *bookRepository) GetPagesPaths(ctx context.Context, pagesID string) ([]entities.Files, error) {
	var files []entities.Files
	if err := br.connection.Where("pages_id", pagesID).Find(&files).Error; err != nil {
		return []entities.Files{}, err
	}

	return files, nil
}

func (br *bookRepository) GetPagesIframe(ctx context.Context, pagesID string) ([]entities.Iframes, error) {
	var Iframes []entities.Iframes
	if err := br.connection.Where("pages_id", pagesID).Find(&Iframes).Error; err != nil {
		return []entities.Iframes{}, err
	}

	return Iframes, nil
}

func (br *bookRepository) GetPagesWorksheets(ctx context.Context, pagesID string) ([]entities.Worksheet, error) {
	var files []entities.Worksheet
	if err := br.connection.Where("pages_id", pagesID).Find(&files).Error; err != nil {
		return []entities.Worksheet{}, err
	}
	return files, nil
}

func (br *bookRepository) DeleteBooks(ctx context.Context, BookID string) error {

	var book entities.Book
	if err := br.connection.Preload("Pages.Files").First(&book, "id = ?", BookID).Error; err != nil {
		return err
	}

	for _, page := range book.Pages {
		for _, file := range page.Files {
			if err := br.connection.Delete(&file).Error; err != nil {
				return err
			}
		}
		if err := br.connection.Delete(&page).Error; err != nil {
			return err
		}
	}

	if err := br.connection.Delete(&entities.Book_User{}, "book_id = ?", BookID).Error; err != nil {
		return err
	}

	if err := br.connection.Delete(&book, "id = ?", BookID).Error; err != nil {
		return err
	}

	return nil
}

func (br *bookRepository) GetUserBooksID(ctx context.Context, userID uuid.UUID) ([]entities.Book_User, error) {
	var userBooks []entities.Book_User
	if err := br.connection.Where("user_id = ?", userID).Find(&userBooks).Error; err != nil {
		return []entities.Book_User{}, err
	}
	return userBooks, nil
}
