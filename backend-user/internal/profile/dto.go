package profile

import (
	"time"

	"github.com/luanjsantos/backend-user/internal/user"
)

// ProfileResponseDTO representa a resposta do perfil do usuário
type ProfileResponseDTO struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Status    string  `json:"status"`
	Bio       *string `json:"bio,omitempty"`
	AvatarURL *string `json:"avatar_url,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	Address   *string `json:"address,omitempty"`
	BirthDate *string `json:"birth_date,omitempty"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

// UpdateProfileDTO representa os dados para atualizar o perfil
type UpdateProfileDTO struct {
	Bio       *string    `json:"bio,omitempty"`
	Phone     *string    `json:"phone,omitempty"`
	Address   *string    `json:"address,omitempty"`
	BirthDate *time.Time `json:"birth_date,omitempty"`
}

// ToProfileResponse converte um User e Profile para ProfileResponseDTO
func ToProfileResponse(user user.User, profile *Profile) ProfileResponseDTO {
	response := ProfileResponseDTO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Status:    string(user.Status),
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	if profile != nil {
		response.Bio = profile.Bio
		response.AvatarURL = profile.AvatarURL
		response.Phone = profile.Phone
		response.Address = profile.Address
		if profile.BirthDate != nil {
			birthDate := profile.BirthDate.Format("2006-01-02")
			response.BirthDate = &birthDate
		}
	}

	return response
}
