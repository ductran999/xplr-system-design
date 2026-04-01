package entity

import "errors"

type Account struct {
	ID      int
	Balance float64
}

func (a *Account) Withdraw(amount float64) error {
	if a.Balance < amount {
		return errors.New("insufficient balance")
	}
	a.Balance -= amount

	return nil
}
