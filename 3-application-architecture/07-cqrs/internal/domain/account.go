package domain

import "errors"

type Account struct {
	ID      int
	Balance float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("must greater than 0")
	}
	a.Balance += amount
	return nil
}
