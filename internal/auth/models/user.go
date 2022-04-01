package models

import (
	"kurneo/internal/infrastructure/hash"
	"time"
)

type User struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Username    string    `gorm:"not null;unique" json:"-"`
	DOB         time.Time `gorm:"null" json:"dob"`
	About       string    `gorm:"null;size:255;column:about" json:"about"`
	Avatar      string    `gorm:"null;size:255" json:"avatar"`
	Name        string    `gorm:"not null;size:100" json:"name"`
	Email       string    `gorm:"not null;unique;size:100" json:"email"`
	Password    string    `gorm:"not null" json:"-"`
	Gender      string    `gorm:"null" json:"gender"`
	LastLoginAt time.Time `gorm:"null" json:"last_login_at"`
	CreatedAt   time.Time `gorm:"null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"null" json:"updated_at"`
}

func (user *User) VerifyPassword(password string) bool {
	return hash.Check(user.Password, password) == nil
}
