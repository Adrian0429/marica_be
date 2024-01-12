package controller

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookController interface {
	CreateBook(c *gin.Context)
	// UpdateBook(c *gin.Context)
	GetAllBooks(c *gin.Context)
	GetBookPages(c *gin.Context)
	GetBookAllPages(c *gin.Context)
	GetTopBooks(c *gin.Context)
	GetImage(c *gin.Context)
	GetAllBooksAdmin(c *gin.Context)
	DeleteBooks(c *gin.Context)
	GetUserBooks(ctx *gin.Context)
	GetBookPreview(ctx *gin.Context)
}

type bookController struct {
	bookService services.BookService
	jwtService  services.JWTService
	userService services.UserService
}

func NewBookController(bs services.BookService, jwt services.JWTService, us services.UserService) BookController {
	return &bookController{
		bookService: bs,
		jwtService:  jwt,
		userService: us,
	}
}

func (bc *bookController) CreateBook(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := bc.jwtService.GetIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	user, err := bc.userService.GetUserByID(ctx, userID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "userID tidak valid", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if user.Role != "admin" {
		res := utils.BuildResponseFailed("Tidak memiliki Akses", "Role Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	title := ctx.PostForm("title")
	if checkTitle, _ := bc.bookService.CheckTitle(ctx.Request.Context(), title); checkTitle {
		res := utils.BuildResponseFailed("Judul Sudah Terdaftar", "failed", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	desc := ctx.PostForm("description")
	if desc == "" || title == "" {
		res := utils.BuildResponseFailed("Failed to retrieve title/desc", "", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	tags := ctx.PostForm("tags")
	tokped := ctx.PostForm("tokped_link")
	page_count := ctx.PostForm("page_count")
	parsedPageCount, err := strconv.Atoi(page_count)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid page_count"})
		return
	}

	thumbnail, err := ctx.FormFile("thumbnail")
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve thumbnail", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	req := dto.BookCreateRequest{
		UserID:      userID,
		Title:       title,
		Desc:        desc,
		Page_Count:  parsedPageCount,
		Tags:        tags,
		Tokped_Link: tokped,
		Thumbnail:   thumbnail,
	}

	for i := 0; ; i++ {
		filesUploaded := false
		var medias dto.MediaRequest
		medias.Index = i
		medias.Title = ctx.PostForm("Pages[" + strconv.Itoa(i) + "].Title")
		for j := 0; ; j++ {
			photo, err := ctx.FormFile("Pages[" + strconv.Itoa(i) + "].Files[" + strconv.Itoa(j) + "]")
			if err != nil {
				break
			}
			iframe := ctx.PostForm("Iframe[" + strconv.Itoa(i) + "].Files[" + strconv.Itoa(j) + "]")

			id1, err := strconv.Atoi(ctx.PostForm("WS[" + strconv.Itoa(i) + "].id[" + strconv.Itoa(j) + "]"))
			id2, err := strconv.Atoi(ctx.PostForm("WS[" + strconv.Itoa(i) + "].id2[" + strconv.Itoa(j) + "]"))
			WS_Input := dto.Worksheet{
				Worksheet_ID:  id1,
				String_Code:   ctx.PostForm("WS[" + strconv.Itoa(i) + "].code[" + strconv.Itoa(j) + "]"),
				Worksheet_ID2: id2,
				String_Code2:  ctx.PostForm("WS[" + strconv.Itoa(i) + "].code2[" + strconv.Itoa(j) + "]"),
				String_Code3:  ctx.PostForm("WS[" + strconv.Itoa(i) + "].code3[" + strconv.Itoa(j) + "]"),
			}

			Iframes := dto.IFrames{
				Index:  j,
				Iframe: iframe,
			}

			files := dto.Files{
				Index:  j,
				Images: photo,
			}

			medias.Files = append(medias.Files, files)

			if iframe != "" {
				medias.IFrames = append(medias.IFrames, Iframes)
			}

			if WS_Input.String_Code != "" {
				medias.Worksheet = append(medias.Worksheet, WS_Input)
			}

			filesUploaded = true
		}

		if !filesUploaded {
			break
		}

		req.MediaRequest = append(req.MediaRequest, medias)
	}

	Page, err := bc.bookService.CreateBook(ctx, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to create Books", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Successfully create Books", Page)
	ctx.JSON(http.StatusOK, res)
}

func (bc *bookController) GetAllBooks(c *gin.Context) {
	result, err := bc.bookService.GetAllBooks(c.Request.Context())

	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan List Buku", err.Error(), utils.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("Berhasil Mendapatkan List Buku", result)
	c.JSON(http.StatusOK, res)
}

func (bc *bookController) GetTopBooks(c *gin.Context) {
	result, err := bc.bookService.GetTopBooks(c.Request.Context())

	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan List Buku", err.Error(), utils.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan List Buku", result)
	c.JSON(http.StatusOK, res)
}

func (bc *bookController) GetBookPages(ctx *gin.Context) {
	id := ctx.Param("book_id")
	page := ctx.Query("page")
	if page == "" {
		page = "0"
	}

	Books, err := bc.bookService.GetBookPages(ctx, id, page)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Detail Buku", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Buku", Books)
	ctx.JSON(http.StatusOK, res)
}

func (bc *bookController) GetBookAllPages(ctx *gin.Context) {
	id := ctx.Param("book_id")

	Books, err := bc.bookService.GetBookAllPages(ctx, id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan all page Buku", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan all page Buku", Books)
	ctx.JSON(http.StatusOK, res)
}

func (bc *bookController) GetImage(ctx *gin.Context) {
	path := ctx.Param("path")
	dirname := ctx.Param("dirname")
	filename := ctx.Param("filename")

	imagePath := "storage/" + path + "/" + dirname + "/" + filename

	_, err := os.Stat(imagePath)
	if os.IsNotExist(err) {
		ctx.JSON(400, gin.H{
			"message": "image not found",
		})
		return
	}

	ctx.File(imagePath)
}

func (bc *bookController) DeleteBooks(ctx *gin.Context) {
	BookId := ctx.Param("book_id")

	Pages, err := bc.bookService.GetBookPages(ctx.Request.Context(), BookId, "0")
	if err != nil {
		return
	}

	utils.DeleteFiles(Pages.Title)

	if err := bc.bookService.DeleteBooks(ctx.Request.Context(), BookId); err != nil {
		res := utils.BuildResponseFailed("Gagal Menghapus Buku", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menghapus Buku", utils.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (bc *bookController) GetAllBooksAdmin(c *gin.Context) {
	result, err := bc.bookService.GetAllBooksAdmin(c.Request.Context())

	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan List Buku admin", err.Error(), utils.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("Berhasil Mendapatkan List Buku admin", result)
	c.JSON(http.StatusOK, res)
}

func (bc *bookController) GetBookPreview(c *gin.Context) {
	id := c.Param("book_id")

	result, err := bc.bookService.GetBookPreview(c.Request.Context(), id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Preview Buku", err.Error(), utils.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Preview Buku", result)
	c.JSON(http.StatusOK, res)
}

func (bc *bookController) GetUserBooks(ctx *gin.Context) {
	uid := ctx.Query("uid")
	if uid == "" {
		token := ctx.MustGet("token").(string)
		userID, err := bc.jwtService.GetIDByToken(token)

		if err != nil {
			res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		result, err := bc.bookService.GetUserBooks(ctx.Request.Context(), userID)
		if err != nil {
			res := utils.BuildResponseFailed("Gagal Mendapatkan Buku yang dimiliki User", err.Error(), utils.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, res)
			return
		}
		res := utils.BuildResponseSuccess("Berhasil Mendapatkan List Buku yang dimiliki User", result)
		ctx.JSON(http.StatusOK, res)
	} else {
		parsedUUID, err := uuid.Parse(uid)
		if err != nil {
			fmt.Println("Error parsing UUID:", err)
			return
		}
		result, err := bc.bookService.GetUserBooks(ctx.Request.Context(), parsedUUID)
		if err != nil {
			res := utils.BuildResponseFailed("Gagal Mendapatkan Buku yang dimiliki User", err.Error(), utils.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, res)
			return
		}
		res := utils.BuildResponseSuccess("Berhasil Mendapatkan List Buku yang dimiliki User", result)
		ctx.JSON(http.StatusOK, res)
	}

}
