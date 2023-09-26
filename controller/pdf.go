package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"

	"github.com/gin-gonic/gin"
)

type (
	PdfController interface {
		GetPdf(ctx *gin.Context)
		UploadPDF(ctx *gin.Context)
	}

	pdfController struct {
		pdfService services.PdfService
	}
)

func NewPdfController(ps services.PdfService) PdfController {
	return &pdfController{
		pdfService: ps,
	}
}

func (pc pdfController) UploadPDF(ctx *gin.Context) {
	file, err := ctx.FormFile("pdf_form")
	if err != nil {
		response := utils.BuildResponseFailed("Failed to process request", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	bytes, err := utils.IsBase64(*file)
	if err != nil {
		response := utils.BuildResponseFailed("Failed to process request", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = utils.SavePDF(bytes, "pdf", "UMUM", file.Filename)
	if err != nil {
		response := utils.BuildResponseFailed("Failed to process request", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	fmt.Print("Sukses")
}

func (pc pdfController) GetPdf(ctx *gin.Context) {

	path := ctx.Param("path")
	dirName := ctx.Param("dirname")
	file := ctx.Param("filename")
	pdfPath := "storage" + "/" + path + "/" + dirName + "/" + file

	_, err := os.Stat(pdfPath)
	if err != nil {
		if os.IsNotExist(err) {
			response := utils.BuildResponseFailed("Failed to process request", err.Error(), utils.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}
	ctx.File(pdfPath)

}
