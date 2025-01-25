package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username" gorm:"column:username;type:varchar(20);unique" validate:"required"`
	Email       string `json:"email" gorm:"column:email;type:varchar(100);unique" validate:"required,email"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number;type:varchar(15)" validate:"required"`
	FullName    string `json:"full_name" gorm:"column:full_name;type:varchar(100)" validate:"required"`
	Address     string `json:"address" gorm:"column:address;type:text" validate:"required"`
	Dob         string `json:"dob" gorm:"column:dob;type:date" validate:"required"`
	Password    string `json:"password,omitempty" gorm:"column:password;type:varchar(100)" validate:"required"`
	CreatedAt   string `json:"-" gorm:"column:created_at"`
	UpdatedAt   string `json:"-"`
}

func (*User) TableName() string {
	return "users"
}

func (l User) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type UserSession struct {
	ID                  int `gorm:"primary_key"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	UserID              int       `json:"user_id" gorm:"type:int" validate:"required"`
	Token               string    `json:"token" gorm:"type:varchar(512)" validate:"required"`
	RefreshToken        string    `json:"refresh_token" gorm:"type:varchar(512)" validate:"required"`
	TokenExpired        time.Time `json:"-" validate:"required"`
	RefreshTokenExpired time.Time `json:"-" validate:"required"`
}

func (*UserSession) TableName() string {
	return "user_sessions"
}

func (l UserSession) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
