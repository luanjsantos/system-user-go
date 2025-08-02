package profile

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, h *Handler) {
	group := r.Group("/profile")
	group.GET("", h.GetProfile)
	group.PUT("", h.UpdateProfile)
}
