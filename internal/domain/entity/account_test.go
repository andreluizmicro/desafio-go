package entity

import (
	"errors"
	"github.com/andreluizmicro/desafio-go/internal/domain/value_object"
	"testing"
)

type testcase struct {
	ID            *value_object.ID
	User          *User
	Balance       *float64
	Credit        float64
	ExpectedError error
}

func TestCreateNewAccount(t *testing.T) {
	user, _ := createUser("André Silva", "andre@gmail.com", "045.454.888.84")

	t.Run("test should create new account", func(t *testing.T) {
		var balance float64 = 100
		testCases := []testcase{
			{ID: value_object.NewID(), User: user, Balance: &balance, ExpectedError: nil},
		}
		for _, item := range testCases {
			account, err := newAccount(*item.ID, item.User, *item.Balance)
			if err != nil && !errors.Is(err, item.ExpectedError) {
				t.Errorf("Expected %f but got %f", item.ExpectedError, err)
			}
			if account.ID().Value == "" {
				t.Errorf("Expected %f but got %f", item.ExpectedError, err)
			}
		}
	})

	t.Run("test should return error when try credit account", func(t *testing.T) {
		var balance float64 = 100
		var negativeBalance float64 = -100
		testCases := []testcase{
			{ID: value_object.NewID(), User: user, Balance: &balance, ExpectedError: ErrCreditValue},
			{ID: value_object.NewID(), User: user, Balance: &negativeBalance, ExpectedError: ErrCreditValue},
		}
		for _, item := range testCases {
			account, err := newAccount(*item.ID, item.User, *item.Balance)
			err = account.CreditAccount(0)
			if err != nil && !errors.Is(err, item.ExpectedError) {
				t.Errorf("Expected %f but got %f", item.ExpectedError, err)
			}
		}
	})

	t.Run("test should return error when try credit account", func(t *testing.T) {
		var balance float64 = 0
		testCases := []testcase{
			{ID: value_object.NewID(), User: user, Balance: &balance, ExpectedError: ErrInsufficientBalance},
		}
		for _, item := range testCases {
			account, err := newAccount(*item.ID, item.User, *item.Balance)
			err = account.DebitAccount(50)
			if err != nil && !errors.Is(err, item.ExpectedError) {
				t.Errorf("Expected %f but got %f", item.ExpectedError, err)
			}
			if account.Balance() != 0 {
				t.Errorf("Expected %f but got %f", item.ExpectedError, err)
			}
		}
	})

	t.Run("test should return error when try create account with negative balance", func(t *testing.T) {
		var balance float64 = -100
		testCases := []testcase{
			{ID: value_object.NewID(),
				User:          user,
				Balance:       &balance,
				ExpectedError: ErrCreateAccountWithNegativeBalance,
			},
		}
		for _, item := range testCases {
			_, err := newAccount(*item.ID, item.User, *item.Balance)
			if err != nil && !errors.Is(err, item.ExpectedError) {
				t.Errorf("Expected %f but got %f", item.ExpectedError, err)
			}
		}
	})
}
