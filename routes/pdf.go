package routes

import (
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/gin-gonic/gin"
)

func Pdf(route *gin.Engine, PdfController controller.PdfController) {
	routes := route.Group("/api/pdf")
	{
		routes.POST("", PdfController.UploadPDF)
		routes.GET("get/:path/:dirname/:filename", PdfController.GetPdf)
	}

}
