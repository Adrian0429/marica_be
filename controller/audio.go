package controller

import (
	"net/http"
	"os"

	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"

	"github.com/gin-gonic/gin"
)

type (
	AudioController interface {
		GetAudio(ctx *gin.Context)
	}

	audioController struct {
		audioService services.AudioService
	}
)


func NewAudioController(as services.AudioService) AudioController {
	return &audioController{
		audioService: as,
	}
}

func (ac audioController) GetAudio(ctx *gin.Context) {

	path := ctx.Param("path")
	dirName := ctx.Param("dirname")
	file := ctx.Param("filename")
	audioPath := "storage" + "/" + path + "/" + dirName + "/" + file

	_, err := os.Stat(audioPath)
	if err != nil {
		if os.IsNotExist(err) {
			response := utils.BuildResponseFailed("Failed to process request", err.Error(), utils.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}
	ctx.File(audioPath)

}
