package user

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/luanjsantos/backend-user/utils"
	"gorm.io/gorm"
)

// Create godoc
//
//	@Summary	Cria um usuário
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Security	BearerAuth
//	@Param		user	body		CreateUserDTO	true	"Dados do usuário"
//	@Success	201		{object}	User
//	@Failure	400		{object}	map[string]string
//	@Router		/users/ [post]
func (h *Handler) Create(c *gin.Context) {
	var dto CreateUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make(map[string]string)
			for _, fe := range ve {
				out[fe.Field()] = customErrorMessage(fe)
			}
			utils.LogError(err, map[string]interface{}{
				"action": "create_user",
				"errors": out,
			})
			c.JSON(http.StatusBadRequest, gin.H{"errors": out})
			return
		}
		utils.LogError(err, map[string]interface{}{
			"action": "create_user",
		})
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Create(dto)
	if err != nil {
		utils.LogError(err, map[string]interface{}{
			"action": "create_user",
			"email":  dto.Email,
		})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"})
		return
	}

	utils.LogInfo("Usuário criado com sucesso", map[string]interface{}{
		"action":  "create_user",
		"user_id": user.ID,
		"email":   user.Email,
	})
	c.JSON(http.StatusCreated, user)
}

// GetAll godoc
//
//	@Summary	Lista todos os usuários
//	@Tags		users
//	@Produce	json
//	@Security	BearerAuth
//	@Success	200	{array}	User
//	@Router		/users/ [get]
func (h *Handler) GetAll(c *gin.Context) {
	users, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuários"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetOne godoc
//
//	@Summary	Busca um usuário específico
//	@Tags		users
//	@Produce	json
//	@Security	BearerAuth
//	@Param		id	path		int	true	"ID do usuário"
//	@Success	200	{object}	User
//	@Failure	404	{object}	map[string]string
//	@Failure	500	{object}	map[string]string
//	@Router		/users/{id} [get]
func (h *Handler) GetOne(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID é obrigatório"})
		return
	}

	// Converter string para uint
	var userID uint
	if _, err := fmt.Sscanf(id, "%d", &userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	user, err := h.service.GetOne(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuário"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Update godoc
//
//	@Summary	Atualiza um usuário
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Security	BearerAuth
//	@Param		id		path		int				true	"ID do usuário"
//	@Param		user	body		UpdateUserDTO	true	"Dados do usuário"
//	@Success	200		{object}	User
//	@Failure	400		{object}	map[string]string
//	@Failure	404		{object}	map[string]string
//	@Failure	500		{object}	map[string]string
//	@Router		/users/{id} [put]
func (h *Handler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID é obrigatório"})
		return
	}

	var userID uint
	if _, err := fmt.Sscanf(id, "%d", &userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var dto UpdateUserDTO
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

	user, err := h.service.Update(userID, dto)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar usuário"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// ResetPassword godoc
//
//	@Summary	Redefine a senha de um usuário
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Security	BearerAuth
//	@Param		id		path		int					true	"ID do usuário"
//	@Param		user	body		ResetPasswordDTO	true	"Nova senha"
//	@Success	200		{object}	map[string]string
//	@Failure	400		{object}	map[string]string
//	@Failure	404		{object}	map[string]string
//	@Failure	500		{object}	map[string]string
//	@Router		/users/{id}/reset-password [put]
func (h *Handler) ResetPassword(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID é obrigatório"})
		return
	}

	var userID uint
	if _, err := fmt.Sscanf(id, "%d", &userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

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

	err := h.service.ResetPassword(userID, dto)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao redefinir senha"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Senha redefinida com sucesso"})
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
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
