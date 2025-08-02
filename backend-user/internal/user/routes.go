package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, h *Handler) {
	group := r.Group("/users")
	group.GET("/", h.GetAll)
	group.GET("/:id", h.GetOne)
	group.POST("/", h.Create)
	group.PUT("/:id", h.Update)
	group.PUT("/:id/reset-password", h.ResetPassword)
}
