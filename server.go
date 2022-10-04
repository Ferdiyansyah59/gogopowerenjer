package main

import (
	"bensi-api/config"
	"bensi-api/controller"
	"bensi-api/middleware"
	"bensi-api/repository"
	"bensi-api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
	// Repo
	userRepository    repository.UserRepository    = repository.NewUserRepository(db)
	articleRepository repository.ArticleRepository = repository.NewArticleRepository(db)
	// Service
	jwtService     service.JWTService     = service.NewJWTService()
	authService    service.AuthService    = service.NewAuthServie(userRepository)
	userService    service.UserService    = service.NewUserService(userRepository)
	articleService service.ArticleService = service.NewArticleService(articleRepository)

	// Controller
	authController    controller.AuthController    = controller.NewAuthController(authService, jwtService)
	userController    controller.UserController    = controller.NewUserController(userService, jwtService)
	articleController controller.ArticleContorller = controller.NewArticleController(articleService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
		authRoutes.GET("/all-user", userController.GetAllUser)
	}

	userRoutes := r.Group("api", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	articleRoutes := r.Group("api/article", middleware.AuthorizeJWT(jwtService))
	{
		articleRoutes.GET("/", articleController.GetAllArticle)
		articleRoutes.POST("/", articleController.InsertArticle)
		articleRoutes.DELETE("/:slug", articleController.DeleteArticle)
		articleRoutes.PUT("/:slug", articleController.UpdateArticle)
	}

	r.Run()
}
