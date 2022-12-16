package repository

import (
	"context"
	"game-store-final-project/project/domain/entity/customer"
)

type CustomerRepository interface {
	StoreCustomer(ctx context.Context, dataCustomer *customer.Customer) error
	GetCustomerByNik(ctx context.Context, nik string) (*customer.Customer, error)
	IndexCustomerWithTransaction(ctx context.Context, nik string) (*customer.Customer, error)
	GetVoucherByCustomerId(ctx context.Context, id string) (*customer.Customer, error)
}
