package profile

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id" gorm:"not null;uniqueIndex"`
	Bio       *string        `json:"bio"`
	AvatarURL *string        `json:"avatar_url"`
	Phone     *string        `json:"phone" gorm:"size:20"`
	Address   *string        `json:"address"`
	BirthDate *time.Time     `json:"birth_date"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
