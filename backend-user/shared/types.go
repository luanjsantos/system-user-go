package shared

// UserStatus representa o status do usuário
type UserStatus string

const (
	StatusActive   UserStatus = "active"
	StatusInactive UserStatus = "inactive"
)

// String retorna a representação string do status
func (s UserStatus) String() string {
	return string(s)
}

// IsValid verifica se o status é válido
func (s UserStatus) IsValid() bool {
	switch s {
	case StatusActive, StatusInactive:
		return true
	default:
		return false
	}
}
