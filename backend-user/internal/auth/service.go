package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/luanjsantos/backend-user/internal/user"
	"github.com/luanjsantos/backend-user/shared"
	"github.com/luanjsantos/backend-user/utils"
)

type Service interface {
	Login(dto LoginDTO) (LoginResponseDTO, error)
	Logout(token string) error
	ValidateToken(token string) (*Claims, error)
	ResetPassword(dto ResetPasswordDTO) (ResetPasswordResponseDTO, error)
}

type service struct {
	userService user.Service
	secretKey   []byte
}

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func NewService(userService user.Service, secretKey string) Service {
	return &service{
		userService: userService,
		secretKey:   []byte(secretKey),
	}
}

func (s *service) Login(dto LoginDTO) (LoginResponseDTO, error) {
	// Buscar usuário por email
	users, err := s.userService.GetAll()
	if err != nil {
		return LoginResponseDTO{}, err
	}

	var foundUser user.User
	for _, u := range users {
		if u.Email == dto.Email {
			foundUser = u
			break
		}
	}

	if foundUser.ID == 0 {
		return LoginResponseDTO{}, errors.New("usuário não encontrado")
	}

	// Verificar se usuário está ativo
	if foundUser.Status != shared.StatusActive {
		return LoginResponseDTO{}, errors.New("usuário inativo")
	}

	// Verificar senha
	if !utils.CheckPassword(dto.Password, foundUser.Password) {
		return LoginResponseDTO{}, errors.New("senha incorreta")
	}

	// Gerar token JWT
	token, err := s.generateToken(foundUser)
	if err != nil {
		return LoginResponseDTO{}, err
	}

	// Converter para response DTO
	responseUser := User{
		ID:        foundUser.ID,
		Name:      foundUser.Name,
		Email:     foundUser.Email,
		Status:    string(foundUser.Status),
		CreatedAt: foundUser.CreatedAt.Format(time.RFC3339),
	}

	return LoginResponseDTO{
		Token: token,
		User:  responseUser,
	}, nil
}

func (s *service) Logout(token string) error {
	// Em uma implementação real, você poderia invalidar o token
	// Por exemplo, adicionando à uma blacklist
	// Por simplicidade, apenas validamos o token
	_, err := s.ValidateToken(token)
	return err
}

func (s *service) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return s.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token inválido")
}

func (s *service) generateToken(user user.User) (string, error) {
	claims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24 horas
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}

func (s *service) ResetPassword(dto ResetPasswordDTO) (ResetPasswordResponseDTO, error) {
	// Buscar usuário por email
	users, err := s.userService.GetAll()
	if err != nil {
		return ResetPasswordResponseDTO{}, err
	}

	var foundUser user.User
	for _, u := range users {
		if u.Email == dto.Email {
			foundUser = u
			break
		}
	}

	if foundUser.ID == 0 {
		return ResetPasswordResponseDTO{}, errors.New("usuário não encontrado")
	}

	// Verificar se usuário está ativo
	if foundUser.Status != shared.StatusActive {
		return ResetPasswordResponseDTO{}, errors.New("usuário inativo")
	}

	// Resetar senha para "123"
	newPassword := "123"

	// Atualizar senha no banco
	err = s.userService.ResetPassword(foundUser.ID, user.ResetPasswordDTO{Password: newPassword})
	if err != nil {
		return ResetPasswordResponseDTO{}, err
	}

	return ResetPasswordResponseDTO{
		Message: "Senha alterada com sucesso",
	}, nil
}
