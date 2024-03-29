package main

import (
	"log"
	"os"

	"github.com/Caknoooo/golang-clean_template/config"
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/Caknoooo/golang-clean_template/middleware"
	"github.com/Caknoooo/golang-clean_template/migrations"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/Caknoooo/golang-clean_template/routes"
	"github.com/Caknoooo/golang-clean_template/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	var (
		db                  *gorm.DB                       = config.SetUpDatabaseConnection()
		jwtService          services.JWTService            = services.NewJWTService()
		pageRepository      repository.PagesRepository     = repository.NewPagesRepository(db)
		filesRepository     repository.FilesRepository     = repository.NewFilesRepository(db)
		iframesRepository   repository.IframesRepository   = repository.NewIframesRepository(db)
		worksheetRepository repository.WorksheetRepository = repository.NewWorksheetRepository(db)
		passwordRepository  repository.PasswordRepository  = repository.NewPasswordRepository(db)

		userRepository repository.UserRepository = repository.NewUserRepository(db)
		userService    services.UserService      = services.NewUserService(userRepository, passwordRepository)
		userController controller.UserController = controller.NewUserController(userService, jwtService)

		bookRepository repository.BookRepository = repository.NewBookRepository(db)
		bookService    services.BookService      = services.NewBookService(bookRepository, pageRepository, filesRepository, iframesRepository, worksheetRepository)
		bookController controller.BookController = controller.NewBookController(bookService, jwtService, userService)

		adminRepository repository.UserRepository  = repository.NewUserRepository(db)
		adminService    services.AdminService      = services.NewAdminService(adminRepository)
		adminController controller.AdminController = controller.NewAdminController(adminService, jwtService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
	routes.User(server, userController, jwtService, bookController)
	routes.Admin(server, adminController, bookController, jwtService)

	if err := migrations.Seeder(db); err != nil {
		log.Fatalf("error migration seeder: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	server.Run(":" + port)
}
