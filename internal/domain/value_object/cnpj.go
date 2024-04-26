package value_object

import (
	"errors"
	"regexp"
)

const cnpjLength = 14

var (
	ErrInvalidCNPJ = errors.New("invalid cnpj")
)

type CNPJ struct {
	value string
}

func NewCNPJ(value string) (*CNPJ, error) {
	cnpj := CNPJ{value: value}
	err := cnpj.validate()
	if err != nil {
		return nil, err
	}
	return &cnpj, nil
}

func (c *CNPJ) validate() error {
	cpfRegex := regexp.MustCompile(`[^0-9]`)
	cpf := cpfRegex.ReplaceAllString(c.value, "")

	if len(cpf) != cnpjLength {
		return ErrInvalidCNPJ
	}
	return nil
}
