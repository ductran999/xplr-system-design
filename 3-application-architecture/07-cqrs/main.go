package main

import (
	"cqrs/internal/command"
	"cqrs/internal/domain"
	"cqrs/internal/query"
	"fmt"
	"log/slog"
)

func main() {
	myAccount := &domain.Account{ID: 1, Balance: 1000}

	cmdHandler := &command.CommandHandler{}
	queryHandler := &query.QueryHandler{}

	fmt.Println("--- Write flow ---")
	err := cmdHandler.Handle(command.DepositCommand{AccountID: 1, Amount: 500}, myAccount)
	if err != nil {
		slog.Error("write failed", "error", err)
	}

	fmt.Println("--- Read flow ---")
	display := queryHandler.GetBalanceDisplay(myAccount.ID, myAccount.Balance)
	fmt.Printf("Output: %+v\n", display)
}
