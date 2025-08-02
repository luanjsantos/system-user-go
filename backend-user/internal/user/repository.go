package user

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]User, error)
	FindOne(id uint) (User, error)
	Create(user User) (User, error)
	Update(id uint, user User) (User, error)
	UpdatePassword(id uint, password string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) FindOne(id uint) (User, error) {
	var user User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *repository) Create(user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) Update(id uint, user User) (User, error) {
	err := r.db.Model(&User{}).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return User{}, err
	}
	return r.FindOne(id)
}

func (r *repository) UpdatePassword(id uint, password string) error {
	return r.db.Model(&User{}).Where("id = ?", id).Update("password", password).Error
}
