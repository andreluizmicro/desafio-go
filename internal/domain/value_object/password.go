package value_object

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

const (
	passwordMinLength = 8
	passwordMaxLength = 100
)

var (
	ErrInvalidPassword = errors.New("password must be longer than 8 characters, contain numbers and lowercase and uppercase letters")
)

type Password struct {
	Value string
}

func NewPassword(value string) (*Password, error) {
	password := Password{Value: value}
	err := password.validate()
	if err != nil {
		return nil, err
	}

	err = password.HashPassword(password.Value)
	if err != nil {
		return nil, err
	}

	return &password, nil
}

func (p *Password) validate() error {
	if len(p.value) <= passwordMinLength || len(p.value) > passwordMaxLength {
		return ErrInvalidPassword
	}
	hasLower := regexp.MustCompile(`[a-z]`).MatchString
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString

	if !hasLower(p.value) || !hasUpper(p.value) {
		return ErrInvalidPassword
	}
	return nil
}

func (p *Password) HashPassword(rawPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	p.value = string(hash)
	return nil
}
