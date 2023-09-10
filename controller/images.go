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
	ImageController interface {
		UploadImage(ctx *gin.Context)
		GetImage(ctx *gin.Context)
	}

	imageController struct {
		imageService services.ImageService
	}
)

const PATH = "storage/"

func NewImageController(is services.ImageService) ImageController {
	return &imageController{
		imageService: is,
	}
}

func (ic *imageController) UploadImage(ctx *gin.Context) {
	file, err := ctx.FormFile("image_form")
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
	// fmt.Println(file)

	err = utils.SaveImage(bytes, "images", "UMUM", file.Filename)
	if err != nil {
		response := utils.BuildResponseFailed("Failed to process request", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	fmt.Print("Sukses")

}

func (ic *imageController) GetImage(ctx *gin.Context) {

	path := ctx.Param("path")
	dirName := ctx.Param("dirname")
	file := ctx.Param("filename")
	imagePath := "storage" + "/" + path + "/" + dirName + "/" + file

	_, err := os.Stat(imagePath)
	if err != nil {
		if os.IsNotExist(err) {
			response := utils.BuildResponseFailed("Failed to process request", err.Error(), utils.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}

	ctx.File(imagePath)
}