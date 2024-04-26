package entity

import (
	"errors"
	valueobject "github.com/andreluizmicro/desafio-go/internal/domain/value_object"
	"time"
)

var (
	ErrInvalidLegalEntity   = errors.New("legal entity must have the CNPJ")
	ErrInvalidNaturalPerson = errors.New("natural person does not have a CNPJ")
)

type User struct {
	ID        *valueobject.ID `json:"id"`
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	Password  string          `json:"-"`
	CPF       string          `json:"cpf"`
	CNPJ      *string         `json:"cnpj"`
	UserType  uint8           `json:"user_type"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Deleted   bool            `json:"deleted"`
}

func NewUser(id, cnpj *string, name, email, password, cpf string, userType uint8) (*User, error) {
	user := User{
		Name:     name,
		Email:    email,
		CPF:      cpf,
		CNPJ:     cnpj,
		UserType: userType,
		Password: password,
	}
	if id == nil {
		user.ID = valueobject.NewID()
	}
	return &user, nil
}
