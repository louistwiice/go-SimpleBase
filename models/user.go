package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID				string		`json:"id"`  // ID will be in uuid format
	Email			string		`json:"email"`
	FirstName		string		`json:"first_name" binding:"required"`
	LastName		string		`json:"last_name" binding:"required"`
	Password		string		`json:"password"`
	IsActive		bool		`json:"is_active"`
	IsStaff			bool		`json:"is_staff"`
	IsSuperuser		bool		`json:"is_superuser"`
	CreatedAt		time.Time	`json:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at"`
}

func (u *User) SetPassword() error {
	if u.Password == ""{
		return ErrInvalidPassword
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err !=nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(plain_password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain_password))
}

func (u *User) Validate() error {
	if u.Email =="" || u.FirstName == "" || u.LastName == "" {
		return ErrNullField
	}
	return nil
}