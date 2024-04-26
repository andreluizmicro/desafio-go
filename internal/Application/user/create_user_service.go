package user

import (
	"github.com/andreluizmicro/desafio-go/internal/domain/contract"
)

type CreateUserService struct {
	userRepository contract.UserRepositoryInterface
}

func NewCreateUserService(userRepository contract.UserRepositoryInterface) *CreateUserService {
	return &CreateUserService{userRepository: userRepository}
}

func (s *CreateUserService) Execute(input CreateUserInputDto) (*CreateUserOutputDto, error) {
	//user, err := entity.NewUser(input.Name, input.Email, input.Password, input.CPF, *input.CNPJ, input.UserType)
	//if err != nil {
	//	return nil, err
	//}
	//
	//_, err = s.userRepository.Create(user)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &CreateUserOutputDto{Id: "123"}, nil
	return nil, nil
}
