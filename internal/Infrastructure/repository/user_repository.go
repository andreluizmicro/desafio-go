package repository

import (
	"database/sql"
	"github.com/andreluizmicro/desafio-go/internal/domain/entity"
	"github.com/andreluizmicro/desafio-go/internal/domain/value_object"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *entity.User) (*value_object.ID, error) {
	stmt, err := r.db.Prepare(
		"INSERT INTO users (id, user_type_id, name, email, cpf, cnpj, password) VALUES (?, ?, ?, ?, ?, ?, ?)",
	)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(
		user.ID.Value,
		user.UserType.Value,
		user.Name.Value,
		user.Email.Value,
		user.CPF.Value,
		nil,
		user.Password.Value,
	)
	if err != nil {
		return nil, err
	}

	return user.ID, nil
}
