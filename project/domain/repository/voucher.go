package repository

import (
	"context"
	"game-store-final-project/project/domain/entity/voucher"
)

type VoucherRepository interface {
	GetVoucherByCode(ctx context.Context, code string) (*voucher.Voucher, error)
	StoreVoucher(ctx context.Context, dataVoucher *voucher.Voucher) error
}
