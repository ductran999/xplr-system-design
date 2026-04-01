package query

import "fmt"

type AccountReadModel struct {
	AccountID   int    `json:"id"`
	BalanceText string `json:"balance_display"`
}

type QueryHandler struct{}

func (h *QueryHandler) GetBalanceDisplay(id int, balance float64) AccountReadModel {
	fmt.Printf("[Query] Query data for account %d...\n", id)

	return AccountReadModel{
		AccountID:   id,
		BalanceText: fmt.Sprintf("$%.2f USD", balance),
	}
}
