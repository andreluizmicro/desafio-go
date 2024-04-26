package value_object

import (
	"errors"
	"regexp"
)

const cpfLength = 11

var (
	ErrInvalidCPF = errors.New("invalid cpf")
)

type CPF struct {
	value string
}

func NewCPF(value string) (*CPF, error) {
	cpf := CPF{value: value}
	err := cpf.validate()
	if err != nil {
		return nil, err
	}
	return &cpf, nil
}

func (c *CPF) validate() error {
	cpfRegex := regexp.MustCompile(`[^0-9]`)
	cpf := cpfRegex.ReplaceAllString(c.value, "")

	if len(cpf) != cpfLength {
		return ErrInvalidCPF
	}
	return nil
}
