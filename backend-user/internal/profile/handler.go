package profile

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

// GetProfile godoc
//
//	@Summary	Obtém o perfil do usuário logado
//	@Tags		profile
//	@Accept		json
//	@Produce	json
//	@Security	BearerAuth
//	@Success	200	{object}	ProfileResponseDTO
//	@Failure	401	{object}	map[string]string
//	@Failure	404	{object}	map[string]string
//	@Router		/profile [get]
func (h *Handler) GetProfile(c *gin.Context) {
	// O middleware de auth já validou o token e colocou o user_id no contexto
	userID := c.GetInt("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	profile, err := h.service.GetProfile(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

// UpdateProfile godoc
//
//	@Summary	Atualiza o perfil do usuário logado
//	@Tags		profile
//	@Accept		json
//	@Produce	json
//	@Security	BearerAuth
//	@Param		profile	body		UpdateProfileDTO	true	"Dados do perfil"
//	@Success	200		{object}	ProfileResponseDTO
//	@Failure	400		{object}	map[string]string
//	@Failure	401		{object}	map[string]string
//	@Failure	404		{object}	map[string]string
//	@Router		/profile [put]
func (h *Handler) UpdateProfile(c *gin.Context) {
	// O middleware de auth já validou o token e colocou o user_id no contexto
	userID := c.GetInt("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	var dto UpdateProfileDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make(map[string]string)
			for _, fe := range ve {
				out[fe.Field()] = customErrorMessage(fe)
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": out})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := h.service.UpdateProfile(uint(userID), dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func customErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " é obrigatório"
	case "email":
		return "E-mail inválido"
	case "min":
		return fe.Field() + " precisa ter no mínimo " + fe.Param() + " caracteres"
	}
	return fe.Field() + " inválido"
}
