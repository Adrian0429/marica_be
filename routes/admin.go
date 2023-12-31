package routes

import (
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/Caknoooo/golang-clean_template/middleware"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/gin-gonic/gin"
)

func Admin(route *gin.Engine, AdminController controller.AdminController, BookController controller.BookController, jwtService services.JWTService) {
	routes := route.Group("/api/admin")
	{
		routes.POST("/login", AdminController.LoginAdmin)
		routes.GET("/me", middleware.Authenticate(jwtService), AdminController.MeAdmin)

		routes.POST("/addBooks", middleware.Authenticate(jwtService), BookController.CreateBook)
		routes.GET("/allPages/:book_id", middleware.Authenticate(jwtService), BookController.GetBookAllPages)
		routes.DELETE("/books/:book_id", middleware.Authenticate(jwtService), BookController.DeleteBooks)
		routes.GET("/books", middleware.Authenticate(jwtService), BookController.GetAllBooksAdmin)
		routes.POST("/:user_id/:book_id", middleware.Authenticate(jwtService), AdminController.GiveAccess)
		routes.DELETE("/:user_id/:book_id", middleware.Authenticate(jwtService), AdminController.GiveAccess)
	}

}
