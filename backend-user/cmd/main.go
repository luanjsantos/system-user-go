package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/luanjsantos/backend-user/config"
	_ "github.com/luanjsantos/backend-user/docs"
	"github.com/luanjsantos/backend-user/internal/auth"
	"github.com/luanjsantos/backend-user/internal/profile"
	"github.com/luanjsantos/backend-user/internal/user"
	"github.com/luanjsantos/backend-user/middleware"
	"github.com/luanjsantos/backend-user/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title						API Usuário
// @version					1.0
// @description				API REST em Go com Gin, GORM e Swagger
// @host						localhost:8080
// @BasePath					/
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				Digite "Bearer" seguido de um espaço e o token JWT
func main() {
	// Inicializar logger
	utils.InitLogger()
	utils.LogInfo("Iniciando aplicação", map[string]interface{}{
		"port": "8080",
		"env":  "development",
	})

	// Inicializar banco de dados
	config.InitDB()
	db := config.DB
	utils.LogInfo("Banco de dados conectado", nil)

	// db.AutoMigrate(&user.User{})

	// User services
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	// Auth services
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		secretKey = "your-secret-key-change-in-production"
	}
	authService := auth.NewService(userService, secretKey)
	authHandler := auth.NewHandler(authService)

	// Profile services
	profileRepo := profile.NewRepository(db)
	profileService := profile.NewService(userService, profileRepo)
	profileHandler := profile.NewHandler(profileService)

	r := gin.New() // Usar gin.New() em vez de gin.Default() para mais controle

	// Middlewares globais
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.LoggingMiddleware())
	r.Use(gin.Recovery()) // Middleware de recovery do gin

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Rotas públicas
	auth.RegisterRoutes(r, authHandler)

	// Rotas protegidas
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware(authService))
	user.RegisterRoutes(protected, userHandler)
	profile.RegisterRoutes(protected, profileHandler)

	utils.LogInfo("Servidor iniciado com sucesso", map[string]interface{}{
		"port": "8080",
	})

	r.Run(":8080")
}
