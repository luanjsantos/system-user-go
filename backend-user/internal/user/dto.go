package user

import "github.com/luanjsantos/backend-user/shared"

type CreateUserDTO struct {
	Name     string            `json:"name" binding:"required,min=3"`
	Email    string            `json:"email" binding:"required,email"`
	Password string            `json:"password" binding:"required,min=6"`
	Status   shared.UserStatus `json:"status" binding:"omitempty,oneof=active inactive"`
}

type UpdateUserDTO struct {
	Name     string            `json:"name" binding:"omitempty,min=3"`
	Email    string            `json:"email" binding:"omitempty,email"`
	Password string            `json:"password" binding:"omitempty,min=6"`
	Status   shared.UserStatus `json:"status" binding:"omitempty,oneof=active inactive"`
}

type ResetPasswordDTO struct {
	Password string `json:"password" binding:"required,min=6"`
}
