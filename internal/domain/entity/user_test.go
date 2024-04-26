package entity

import (
	"errors"
	valueobject "github.com/andreluizmicro/desafio-go/internal/domain/value_object"
	"testing"
)

func TestCreateUser(t *testing.T) {
	type testcase struct {
		ID            *valueobject.ID
		Name          *valueobject.Name
		Email         *valueobject.Email
		Password      *valueobject.Password
		CPF           *valueobject.CPF
		CNPJ          *valueobject.CNPJ
		UserType      *valueobject.UserType
		ExpectedError error
	}

	t.Run("test should create user", func(t *testing.T) {
		validName, _ := valueobject.NewName("John Doe")
		invalidName, _ := valueobject.NewName("AA")
		validEmail, _ := valueobject.NewEmail("johndoe@gmail.com")
		invalidEmail, _ := valueobject.NewEmail("aa")
		validPassword, _ := valueobject.NewPassword("12315Ab@dsa")
		invalidPassword, _ := valueobject.NewPassword("111111")
		invalidLowerLetterPassword, _ := valueobject.NewPassword("1223d12212")
		validCPF, _ := valueobject.NewCPF("207.275.320-14")
		invalidCPF, _ := valueobject.NewCPF("207.275.320")
		validCNPJ, _ := valueobject.NewCNPJ("20.188.512/0001-30")
		invalidCNPJ, _ := valueobject.NewCNPJ("20.188.512")
		validUserTypeDefault, _ := valueobject.NewUserType(1)
		validUserTypeShopkeeper, _ := valueobject.NewUserType(2)
		invalidUserType, _ := valueobject.NewUserType(3)

		testCases := []testcase{
			{
				Name:          validName,
				Email:         validEmail,
				Password:      validPassword,
				CPF:           validCPF,
				CNPJ:          nil,
				UserType:      validUserTypeDefault,
				ExpectedError: nil,
			},
			{
				Name:          invalidName,
				Email:         validEmail,
				Password:      validPassword,
				CPF:           validCPF,
				CNPJ:          nil,
				UserType:      validUserTypeDefault,
				ExpectedError: valueobject.ErrInvalidName,
			},
			{
				Name:          validName,
				Email:         invalidEmail,
				Password:      validPassword,
				CPF:           validCPF,
				CNPJ:          nil,
				UserType:      validUserTypeDefault,
				ExpectedError: valueobject.ErrInvalidEmail,
			},
			{
				Name:          validName,
				Email:         validEmail,
				Password:      invalidPassword,
				CPF:           validCPF,
				CNPJ:          nil,
				UserType:      validUserTypeDefault,
				ExpectedError: valueobject.ErrInvalidPassword,
			},
			{
				Name:          validName,
				Email:         validEmail,
				Password:      invalidLowerLetterPassword,
				CPF:           validCPF,
				CNPJ:          nil,
				UserType:      validUserTypeDefault,
				ExpectedError: valueobject.ErrInvalidPassword,
			},
			{
				Name:          validName,
				Email:         validEmail,
				Password:      validPassword,
				CPF:           invalidCPF,
				CNPJ:          nil,
				UserType:      validUserTypeDefault,
				ExpectedError: valueobject.ErrInvalidCPF,
			},
			{
				Name:          validName,
				Email:         validEmail,
				Password:      validPassword,
				CPF:           validCPF,
				CNPJ:          invalidCNPJ,
				UserType:      validUserTypeDefault,
				ExpectedError: valueobject.ErrInvalidCNPJ,
			},
			{
				Name:          validName,
				Email:         validEmail,
				Password:      validPassword,
				CPF:           validCPF,
				CNPJ:          validCNPJ,
				UserType:      validUserTypeShopkeeper,
				ExpectedError: nil,
			},
			{
				Name:          validName,
				Email:         validEmail,
				Password:      validPassword,
				CPF:           validCPF,
				CNPJ:          nil,
				UserType:      validUserTypeShopkeeper,
				ExpectedError: ErrInvalidNaturalPerson,
			},
			{
				Name:          validName,
				Email:         validEmail,
				Password:      validPassword,
				CPF:           validCPF,
				CNPJ:          validCNPJ,
				UserType:      validUserTypeDefault,
				ExpectedError: ErrInvalidLegalEntity,
			},
			{
				Name:          validName,
				Email:         validEmail,
				Password:      validPassword,
				CPF:           validCPF,
				CNPJ:          nil,
				UserType:      invalidUserType,
				ExpectedError: nil,
			},
		}

		for _, item := range testCases {
			_, err := NewUser(nil, item.Name, item.Email, item.Password, item.CPF, item.CNPJ, item.UserType)
			if err != nil && !errors.Is(err, item.ExpectedError) {
				t.Errorf("Expected %f but got %f", item.ExpectedError, err)
			}
		}
	})
}
