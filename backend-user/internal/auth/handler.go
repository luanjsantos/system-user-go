package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

// Login godoc
//
//	@Summary	Faz login do usuário
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		credentials	body		LoginDTO	true	"Credenciais de login"
//	@Success	200			{object}	LoginResponseDTO
//	@Failure	400			{object}	map[string]string
//	@Failure	401			{object}	map[string]string
//	@Router		/auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	var dto LoginDTO
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

	response, err := h.service.Login(dto)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Logout godoc
//
//	@Summary	Faz logout do usuário
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Security	BearerAuth
//	@Success	200	{object}	map[string]string
//	@Failure	401	{object}	map[string]string
//	@Router		/auth/logout [post]
func (h *Handler) Logout(c *gin.Context) {
	// Extrair token do header Authorization
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
		return
	}

	// Verificar se o header começa com "Bearer "
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
		return
	}

	token := tokenParts[1]
	err := h.service.Logout(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout realizado com sucesso"})
}

// ResetPassword godoc
//
//	@Summary	Redefine a senha do usuário para "123"
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		email	body		ResetPasswordDTO	true	"Email do usuário"
//	@Success	200		{object}	ResetPasswordResponseDTO
//	@Failure	400		{object}	map[string]string
//	@Failure	404		{object}	map[string]string
//	@Router		/auth/reset-password [post]
func (h *Handler) ResetPassword(c *gin.Context) {
	var dto ResetPasswordDTO
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

	response, err := h.service.ResetPassword(dto)
	if err != nil {
		if err.Error() == "usuário não encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// ValidateToken godoc
//
//	@Summary	Valida se o token JWT é válido
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Security	BearerAuth
//	@Success	200	{object}	map[string]interface{}
//	@Failure	401	{object}	map[string]string
//	@Router		/auth/validate [get]
func (h *Handler) ValidateToken(c *gin.Context) {
	// Extrair token do header Authorization
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
		return
	}

	// Verificar se o header começa com "Bearer "
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
		return
	}

	token := tokenParts[1]
	claims, err := h.service.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Token válido",
		"user_id": claims.UserID,
		"email":   claims.Email,
		"exp":     claims.ExpiresAt,
	})
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
