package services

import (
	"context"
	"errors"
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
	GetAllBooks(ctx context.Context) ([]dto.BooksRequest, error)
	GetAllBooksAdmin(ctx context.Context) ([]entities.Book, error)
	GetBookAllPages(ctx context.Context, bookID string) (dto.AllPagesRequest, error)
	GetTopBooks(ctx context.Context) ([]dto.BooksRequest, error)
	GetBookPages(ctx context.Context, bookID string, bookPage string) (dto.BookPageRequest, error)
	CheckTitle(ctx context.Context, Title string) (bool, error)
	DeleteBooks(ctx context.Context, BookID string) error
}

type bookService struct {
	br repository.BookRepository
	pr repository.PagesRepository
	fr repository.FilesRepository
}

func NewBookService(br repository.BookRepository, pr repository.PagesRepository, fr repository.FilesRepository) *bookService {
	return &bookService{
		br: br,
		pr: pr,
		fr: fr,
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
		ID:         bookId,
		Desc:       req.Desc,
		Title:      req.Title,
		UserID:     req.UserID,
		Page_Count: req.Page_Count,
		Tags:       req.Tags,
		View:       0,
		Thumbnail:  thumbnailPath,
	}

	resBooks.ID = bookId.String()
	resBooks.Desc = req.Desc
	resBooks.Thumbnail = thumbnailPath
	resBooks.Title = req.Title
	resBooks.Page_Count = req.Page_Count

	createdBook, err := bs.br.CreateBook(ctx, book)
	if err != nil {
		return dto.BookCreateResponse{}, dto.ErrCreateBooks
	}

	for _, v := range req.MediaRequest {
		var mediaPath string
		var pagesItem entities.Pages
		var medias dto.MediaPathRequest
		var medias_paths dto.Medias
		pagesId := uuid.New()

		pagesItem = entities.Pages{
			ID:        pagesId,
			Index:     v.Index,
			PageTitle: v.Title,
			BookID:    createdBook.ID,
		}

		_, err := bs.pr.CreatePages(ctx, pagesItem)
		if err != nil {
			return dto.BookCreateResponse{}, dto.ErrCreatePages
		}

		for _, w := range v.Files {
			if w.Images != nil {

				bookPages, err := utils.IsBase64(*w.Images)
				if err != nil {
					return dto.BookCreateResponse{}, dto.ErrToBase64
				}

				mediaPath = utils.GenerateFileName(PATH+createdBook.Title, "/Pages_"+strconv.Itoa(v.Index), w.Images.Filename)
				files := entities.Files{
					ID:      uuid.New(),
					Path:    mediaPath,
					Index:   w.Index,
					PagesID: pagesId,
				}
				medias_paths.Index = w.Index
				medias_paths.Path = mediaPath

				_, err = bs.fr.CreateFiles(ctx, files)
				if err != nil {
					return dto.BookCreateResponse{}, errors.New("error input files to db")
				}
				medias.Media = append(medias.Media, medias_paths)
				pagesItem.Files = append(pagesItem.Files, files)

				_ = utils.Upload(bookPages, PATH+createdBook.Title, "/Pages_"+strconv.Itoa(v.Index), w.Images.Filename)
			}
		}

		medias.Index = v.Index
		mediaRequests = append(mediaRequests, medias)

	}

	resBooks.MediaPathRequest = mediaRequests
	return resBooks, nil
}

func (bs *bookService) GetAllBooks(ctx context.Context) ([]dto.BooksRequest, error) {
	Books, err := bs.br.GetAllBooks(ctx)
	if err != nil {
		return nil, dto.ErrGetAllBooks
	}
	var allBooks []dto.BooksRequest
	for _, book := range Books {
		bookProps := dto.BooksRequest{
			ID:        book.ID.String(),
			Title:     book.Title,
			Tags:      book.Tags,
			Desc:      book.Desc,
			Thumbnail: book.Thumbnail,
		}

		allBooks = append(allBooks, bookProps)
	}
	return allBooks, nil
}

func (bs *bookService) GetAllBooksAdmin(ctx context.Context) ([]entities.Book, error) {
	Books, err := bs.br.GetAllBooks(ctx)
	if err != nil {
		return nil, dto.ErrGetAllBooks
	}
	return Books, nil
}

func (bs *bookService) GetTopBooks(ctx context.Context) ([]dto.BooksRequest, error) {
	Books, err := bs.br.GetTopBooks(ctx)
	if err != nil {
		return nil, dto.ErrGetAllBooks
	}

	var allBooks []dto.BooksRequest
	for _, book := range Books {
		bookProps := dto.BooksRequest{
			ID:        book.ID.String(),
			Title:     book.Title,
			Desc:      book.Desc,
			Thumbnail: book.Thumbnail,
			Tags:      book.Tags,
		}

		allBooks = append(allBooks, bookProps)
	}
	return allBooks, nil
}

func (bc *bookService) GetBookAllPages(ctx context.Context, bookID string) (dto.AllPagesRequest, error) {
	books, err := bc.br.GetBookByID(ctx, bookID)
	if err != nil {
		return dto.AllPagesRequest{}, err
	}

	resBooks := dto.AllPagesRequest{
		Title:      books.Title,
		Thumbnail:  books.Thumbnail,
		Desc:       books.Desc,
		UserID:     books.UserID,
		Page_Count: books.Page_Count,
		Tags:       books.Tags,
	}

	Pages, err := bc.br.GetBookAllPages(ctx, bookID)
	if err != nil {
		return dto.AllPagesRequest{}, err
	}

	for _, currPage := range Pages {
		Pages := dto.AllPagesMedia{
			Index: currPage.Index,
			Title: currPage.PageTitle,
		}

		PagePaths, err := bc.br.GetPagesPaths(ctx, currPage.ID.String())
		if err != nil {
			return dto.AllPagesRequest{}, err
		}

		for _, currFiles := range PagePaths {
			filePaths := dto.AllPagesFiles{
				Index: currFiles.Index,
				Path:  currFiles.Path,
			}
			Pages.Files = append(Pages.Files, filePaths)
		}
		resBooks.AllPagesMedia = append(resBooks.AllPagesMedia, Pages)
	}
	return resBooks, nil
}

func (bc *bookService) GetBookPages(ctx context.Context, bookID string, PageIndex string) (dto.BookPageRequest, error) {
	books, err := bc.br.GetBookByID(ctx, bookID)
	if err != nil {
		return dto.BookPageRequest{}, err
	}

	Page, err := bc.br.GetBookPage(ctx, bookID, PageIndex)
	if err != nil {
		return dto.BookPageRequest{}, err
	}

	resBooks := dto.BookPageRequest{
		BookID:     books.ID.String(),
		Title:      books.Title,
		Thumbnail:  books.Thumbnail,
		PageTitle:  Page.PageTitle,
		Page_Count: books.Page_Count,
	}

	PagePaths, err := bc.br.GetPagesPaths(ctx, Page.ID.String())
	if err != nil {
		return dto.BookPageRequest{}, err
	}

	for _, page := range PagePaths {
		pages := dto.PagePaths{
			Path: page.Path,
		}
		resBooks.PagePaths = append(resBooks.PagePaths, pages)
	}

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

func (bc *bookService) DeleteBooks(ctx context.Context, BookID string) error {

	return bc.br.DeleteBooks(ctx, BookID)
}
