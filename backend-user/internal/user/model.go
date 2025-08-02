package user

import (
	"time"

	"github.com/luanjsantos/backend-user/shared"
)

type User struct {
	ID        uint              `gorm:"primaryKey" json:"id"`
	Name      string            `json:"name"`
	Email     string            `json:"email"`
	Password  string            `gorm:"not null" json:"-"`
	Status    shared.UserStatus `gorm:"type:user_status;default:'active'" json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}
