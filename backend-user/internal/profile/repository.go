package profile

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetByUserID(userID uint) (*Profile, error)
	Create(profile *Profile) error
	Update(profile *Profile) error
	Delete(userID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetByUserID(userID uint) (*Profile, error) {
	var profile Profile
	err := r.db.Where("user_id = ?", userID).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *repository) Create(profile *Profile) error {
	return r.db.Create(profile).Error
}

func (r *repository) Update(profile *Profile) error {
	return r.db.Save(profile).Error
}

func (r *repository) Delete(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&Profile{}).Error
}
