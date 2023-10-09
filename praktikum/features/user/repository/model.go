package repository

import "gorm.io/gorm"

type User struct {
	*gorm.Model

	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created"`
	UpdatedAt string `json:"updated"`
	Token     string `json:"token"`
}
