package main

import (
	"clean/internal/adapter/repository"
	"clean/internal/usecase"
	"fmt"
)

func main() {
	repo := repository.NewInMemoryAccountRepo()
	transferService := usecase.NewTransferUseCase(repo)

	from, _ := repo.GetByID(1)
	to, _ := repo.GetByID(2)
	fmt.Printf("Acc1: %.2f, Acc2: %.2f\n", from.Balance, to.Balance)

	fmt.Println("Transfer 200$ from 1 to 2...")
	err := transferService.Execute(1, 2, 200)
	if err == nil {
		from, _ := repo.GetByID(1)
		to, _ := repo.GetByID(2)
		fmt.Printf("Success! New balances: Acc1: %.2f, Acc2: %.2f\n", from.Balance, to.Balance)
	}
}
