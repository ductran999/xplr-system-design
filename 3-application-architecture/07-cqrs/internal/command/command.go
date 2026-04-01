package command

import (
	"cqrs/internal/domain"
	"fmt"
)

type DepositCommand struct {
	AccountID int
	Amount    float64
}

type CommandHandler struct {
	//  Repository to load account from db
}

func (h *CommandHandler) Handle(cmd DepositCommand, acc *domain.Account) error {
	err := acc.Deposit(cmd.Amount)
	if err != nil {
		return err
	}

	fmt.Printf("[Command] %.2f send to %d\n", cmd.Amount, acc.ID)
	return nil
}
