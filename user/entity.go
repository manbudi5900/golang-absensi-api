package user

import (
	"time"
)

type User struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4();primary_key;"`
	Name      string
	Email     string
	RoleID    string
	Status    int
	Username  string
	Phone     string
	Alamat    string
	Avatar    string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
