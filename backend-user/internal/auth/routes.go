package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {
	group := r.Group("/auth")
	group.POST("/login", h.Login)
	group.POST("/logout", h.Logout)
	group.POST("/reset-password", h.ResetPassword)
	group.GET("/validate", h.ValidateToken)
}
