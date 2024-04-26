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
	ID        *valueobject.ID       `json:"id"`
	Name      *valueobject.Name     `json:"name"`
	Email     *valueobject.Email    `json:"email"`
	Password  *valueobject.Password `json:"-"`
	CPF       *valueobject.CPF      `json:"cpf"`
	CNPJ      *valueobject.CNPJ     `json:"cnpj"`
	UserType  *valueobject.UserType `json:"user_type"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
	Deleted   bool                  `json:"deleted"`
}

func NewUser(
	id *valueobject.ID,
	name *valueobject.Name,
	email *valueobject.Email,
	password *valueobject.Password,
	cpf *valueobject.CPF,
	cnpj *valueobject.CNPJ,
	userType *valueobject.UserType,
) (*User, error) {
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

	err := user.validate()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (user *User) validate() error {
	if user.UserType.Value == 1 && user.CNPJ != nil {
		return ErrInvalidLegalEntity
	}
	if user.UserType.Value == 2 && user.CNPJ == nil {
		return ErrInvalidNaturalPerson
	}
	return nil
}
