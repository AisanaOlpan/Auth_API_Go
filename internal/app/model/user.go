package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"password, omitempty"`
	EncryptedPassword string `json:"-"`
}

//Validation
func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(
			&u.Email,            //проверка значения
			validation.Required, //обязательное для заполнения
			is.Email),           //это поле является email

		validation.Field(
			&u.Password, //проверка значения
			validation.By(requiredIf(u.EncryptedPassword == "")), //обязательное для заполнения
			validation.Length(6, 100)),                           //это поле является email

	)
}

//BeforeCreate
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}
	return nil
}

//Sanitize
func (u *User) Sanitize() {
	u.Password = ""
}

//ComparePassword
func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}
func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
