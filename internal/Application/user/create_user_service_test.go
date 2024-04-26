package user

import (
	"github.com/andreluizmicro/desafio-go/internal/domain/value_object"
	repositorymock "github.com/andreluizmicro/desafio-go/test/mock/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateUser(t *testing.T) {
	repositoryMock := &repositorymock.UserRepositoryMock{}
	createUserService := NewCreateUserService(repositoryMock)

	t.Run("test should create user", func(t *testing.T) {
		repositoryMock.On("Create", mock.Anything).Return(value_object.NewID(), nil).Once()
		output, err := createUserService.Execute(
			CreateUserInputDto{
				Name:     "John Doe",
				Email:    "john.doe@example.com",
				Password: "Password123A@s",
				CPF:      "088.988.888-52",
				CNPJ:     nil,
				UserType: 1,
			},
		)
		assert.NoError(t, err)
		assert.NotNil(t, output)
	})
}
