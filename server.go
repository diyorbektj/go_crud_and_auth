package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"test/config"
	"test/controller"
	"test/repository"
	"test/service"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoute := r.Group("api/auth")
	{
		authRoute.POST("/login", authController.Login)
		authRoute.POST("/register", authController.Register)
	}
	userRoute := r.Group("api/user")
	{
		userRoute.GET("/profile", userController.Profile)
		userRoute.PUT("/profile", userController.Update)
	}

	r.Run(":3000")
}
