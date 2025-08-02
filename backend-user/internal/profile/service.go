package profile

import (
	"errors"

	"github.com/luanjsantos/backend-user/internal/user"
)

type Service interface {
	GetProfile(userID uint) (ProfileResponseDTO, error)
	UpdateProfile(userID uint, dto UpdateProfileDTO) (ProfileResponseDTO, error)
}

type service struct {
	userService user.Service
	profileRepo Repository
}

func NewService(userService user.Service, profileRepo Repository) Service {
	return &service{
		userService: userService,
		profileRepo: profileRepo,
	}
}

func (s *service) GetProfile(userID uint) (ProfileResponseDTO, error) {
	// Buscar usuário
	users, err := s.userService.GetAll()
	if err != nil {
		return ProfileResponseDTO{}, err
	}

	var foundUser user.User
	for _, u := range users {
		if u.ID == userID {
			foundUser = u
			break
		}
	}

	if foundUser.ID == 0 {
		return ProfileResponseDTO{}, errors.New("usuário não encontrado")
	}

	// Buscar perfil
	profile, err := s.profileRepo.GetByUserID(userID)
	if err != nil && err.Error() != "record not found" {
		return ProfileResponseDTO{}, err
	}

	return ToProfileResponse(foundUser, profile), nil
}

func (s *service) UpdateProfile(userID uint, dto UpdateProfileDTO) (ProfileResponseDTO, error) {
	// Verificar se usuário existe
	users, err := s.userService.GetAll()
	if err != nil {
		return ProfileResponseDTO{}, err
	}

	var foundUser user.User
	for _, u := range users {
		if u.ID == userID {
			foundUser = u
			break
		}
	}

	if foundUser.ID == 0 {
		return ProfileResponseDTO{}, errors.New("usuário não encontrado")
	}

	// Buscar ou criar perfil
	profile, err := s.profileRepo.GetByUserID(userID)
	if err != nil {
		// Criar novo perfil
		profile = &Profile{
			UserID: userID,
		}
	}

	// Atualizar campos
	if dto.Bio != nil {
		profile.Bio = dto.Bio
	}
	if dto.Phone != nil {
		profile.Phone = dto.Phone
	}
	if dto.Address != nil {
		profile.Address = dto.Address
	}
	if dto.BirthDate != nil {
		profile.BirthDate = dto.BirthDate
	}

	// Salvar
	if profile.ID == 0 {
		err = s.profileRepo.Create(profile)
	} else {
		err = s.profileRepo.Update(profile)
	}

	if err != nil {
		return ProfileResponseDTO{}, err
	}

	return ToProfileResponse(foundUser, profile), nil
}
