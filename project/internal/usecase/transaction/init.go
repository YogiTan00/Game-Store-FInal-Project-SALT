package transaction

import (
	"context"
	"game-store-final-project/project/domain/repository"

	"github.com/stretchr/testify/mock"
)

type RepoTransaction struct {
	mock.Mock
}

type TransactionUseCaseInteractor struct {
	ctx             context.Context
	repoTransaction repository.TransactionRepository
	repoItem        repository.ItemRepository
	repoVoucher     repository.VoucherRepository
}

// ini ngeimplement domain usecase
func NewTransactionUseCaseInteractor(ctx context.Context, repoImplement repository.TransactionRepository, repoImplementItem repository.ItemRepository, repoImplementVoucher repository.VoucherRepository) *TransactionUseCaseInteractor {
	return &TransactionUseCaseInteractor{
		ctx:             ctx,
		repoTransaction: repoImplement,
		repoItem:        repoImplementItem,
		repoVoucher:     repoImplementVoucher,
	}
}
