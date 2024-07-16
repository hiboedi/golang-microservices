package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `json:"email" gorm:"uniqueIndex;not null;type:varchar(100)"`
	Password  string `json:"password" gorm:"not null"`
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Role      string `json:"role" gorm:"not null;sql:enum('admin','user');default:'user'"`
}

type CreateUser struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
}

type UserLogin struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}
