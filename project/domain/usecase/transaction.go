package usecase

import (
	"context"
	"game-store-final-project/project/domain/entity/transaction"
)

type TransactionUseCase interface {
	UcGetAllTransaction(ctx context.Context) ([]*transaction.Transaction, error)
	UcStoreTransaction(ctx context.Context, dataTransaction *transaction.DTOTransaction) error
	UcGetAllTransactionByID(ctx context.Context, id string) ([]*transaction.Transaction, error)
	StoreTransaction(customer_id int, tanggal_pembelian string, voucher []string) (*transaction.Transaction, error)
}
