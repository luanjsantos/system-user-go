package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/luanjsantos/backend-user/utils"
)

// LoggingMiddleware registra informações sobre as requisições usando Logrus
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Tempo de início
		start := time.Now()

		// Processar requisição
		c.Next()

		// Calcular duração
		duration := time.Since(start).Seconds()

		// Logar informações da requisição
		utils.LogRequest(
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			c.Writer.Status(),
			duration,
		)
	}
}

// CORSMiddleware permite requisições CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
