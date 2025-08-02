package auth

type LoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseDTO struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type ResetPasswordDTO struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordResponseDTO struct {
	Message string `json:"message"`
}

type User struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}
