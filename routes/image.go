package routes

import (
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/gin-gonic/gin"
)

func Image(route *gin.Engine, ImageController controller.ImageController) {
	routes := route.Group("/image")
	{
		routes.POST("/AddImages", ImageController.UploadImage)
		routes.GET("/GetImages/:path/:dirname/:filename", ImageController.GetImage)
	}

}
