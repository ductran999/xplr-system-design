package usecase

import "clean/internal/entity"

type AccountRepository interface {
	GetByID(id int) (*entity.Account, error)
	Update(account *entity.Account) error
}

type TransferUseCase struct {
	repo AccountRepository
}

func NewTransferUseCase(r AccountRepository) *TransferUseCase {
	if r == nil {
		panic("AccountRepository cannot be nil")
	}
	return &TransferUseCase{repo: r}
}

func (u *TransferUseCase) Execute(fromID, toID int, amount float64) error {
	from, _ := u.repo.GetByID(fromID)
	to, _ := u.repo.GetByID(toID)

	if err := from.Withdraw(amount); err != nil {
		return err
	}
	to.Balance += amount
	u.repo.Update(from)
	u.repo.Update(to)

	return nil
}
