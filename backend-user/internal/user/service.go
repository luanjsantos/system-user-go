package user

import (
	"reflect"

	"github.com/luanjsantos/backend-user/shared"
	"github.com/luanjsantos/backend-user/utils"
)

type Service interface {
	GetAll() ([]User, error)
	GetOne(id uint) (User, error)
	Create(dto CreateUserDTO) (User, error)
	Update(id uint, dto UpdateUserDTO) (User, error)
	ResetPassword(id uint, dto ResetPasswordDTO) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAll() ([]User, error) {
	return s.repo.FindAll()
}

func (s *service) GetOne(id uint) (User, error) {
	return s.repo.FindOne(id)
}

func (s *service) Create(dto CreateUserDTO) (User, error) {
	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return User{}, err
	}

	// Definir status padrão se não fornecido
	status := dto.Status
	if !status.IsValid() {
		status = shared.StatusActive
	}

	user := User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: hashedPassword,
		Status:   status,
	}
	return s.repo.Create(user)
}

// applyUpdates aplica as atualizações fornecidas no usuário usando reflection
func (s *service) applyUpdates(user *User, dto UpdateUserDTO) error {
	userVal := reflect.ValueOf(user).Elem()
	dtoVal := reflect.ValueOf(dto)
	dtoType := dtoVal.Type()

	// Iterar sobre todos os campos do DTO
	for i := 0; i < dtoVal.NumField(); i++ {
		dtoField := dtoVal.Field(i)
		dtoFieldName := dtoType.Field(i).Name

		// Pular campos vazios
		if dtoField.IsZero() {
			continue
		}

		// Buscar campo correspondente no User
		userField := userVal.FieldByName(dtoFieldName)
		if !userField.IsValid() || !userField.CanSet() {
			continue
		}

		// Tratar password separadamente (precisa de hash)
		if dtoFieldName == "Password" {
			hashedPassword, err := utils.HashPassword(dtoField.String())
			if err != nil {
				return err
			}
			userField.SetString(hashedPassword)
			continue
		}

		// Tratar status (precisa de validação)
		if dtoFieldName == "Status" {
			status := shared.UserStatus(dtoField.String())
			if status.IsValid() {
				userField.SetString(string(status))
			}
			continue
		}

		// Para outros campos, aplicar diretamente
		userField.Set(dtoField)
	}

	return nil
}

func (s *service) Update(id uint, dto UpdateUserDTO) (User, error) {
	// Buscar usuário existente
	existingUser, err := s.repo.FindOne(id)
	if err != nil {
		return User{}, err
	}

	// Aplicar atualizações
	if err := s.applyUpdates(&existingUser, dto); err != nil {
		return User{}, err
	}

	return s.repo.Update(id, existingUser)
}

func (s *service) ResetPassword(id uint, dto ResetPasswordDTO) error {
	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return err
	}

	return s.repo.UpdatePassword(id, hashedPassword)
}
