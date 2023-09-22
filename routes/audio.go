package routes

import (
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/gin-gonic/gin"
)

func Audio(route *gin.Engine, AudioController controller.AudioController) {
	routes := route.Group("/api/audio")
	{
		routes.GET("get/:path/:dirname/:filename", AudioController.GetAudio)
	}

}
