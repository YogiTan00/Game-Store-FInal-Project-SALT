package transaction

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"game-store-final-project/project/domain/entity/transaction"
	"game-store-final-project/project/domain/repository"
	"game-store-final-project/project/internal/repository/mysql/mapper"
	"game-store-final-project/project/internal/repository/mysql/model"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

type TransactionRepositoryMysqlInteractor struct {
	dbConn *sql.DB
}

func NewTransactionMysqlInteractor(conndb *sql.DB) repository.TransactionRepository {
	return &TransactionRepositoryMysqlInteractor{dbConn: conndb}
}

func (t *TransactionRepositoryMysqlInteractor) GetTransaction(ctx context.Context) (*transaction.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s`, model.GetTableTransaction())
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: model.TransactionModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result := dbq.MustQ(ctx, t.dbConn, stmt, opts)
	if result == nil {
		return nil, errors.New("ITEM NOT FOUND")
	}

	dataTransaction, errMap := mapper.ModelToDomainTransaction(result.(*model.TransactionModel))
	if errMap != nil {
		return nil, errMap
	}

	return dataTransaction, nil
}

func (t *TransactionRepositoryMysqlInteractor) GetAllTransaction(ctx context.Context) ([]*transaction.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s`, model.GetTableTransaction())
	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: model.TransactionModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result := dbq.MustQ(ctx, t.dbConn, stmt, opts)
	if result == nil {
		return nil, errors.New("ITEM NOT FOUND")
	}

	dataTransaction, errMap := mapper.ListModelToDomainTransaction(result.([]*model.TransactionModel))
	if errMap != nil {
		return nil, errMap
	}

	return dataTransaction, nil
}

func (t *TransactionRepositoryMysqlInteractor) GetAllTransactionByID(ctx context.Context, id string) ([]*transaction.Transaction, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE id=?`, model.GetTableTransaction())
	rows, errMysql := t.dbConn.QueryContext(ctx, stmt, id)
	if errMysql != nil {
		return nil, errMysql
	}

	dataTransactionColletion := make([]*transaction.Transaction, 0)
	for rows.Next() {
		var (
			idTrans           int
			voucherCustomerId int
			customerId        int
			codeTransaction   string
			tanggalPembelian  time.Time
			total             int64
			hargaDiscount     int64
			totalHarga        int64
		)
		err := rows.Scan(
			&idTrans,
			&voucherCustomerId,
			&customerId,
			&codeTransaction,
			&tanggalPembelian,
			&total,
			&hargaDiscount,
			&totalHarga,
		)
		if err != nil {
			return nil, err
		}
		// trans, _ := time.Parse("INV02D01M2006Y15H04M05S", codeTransaction)

		dataTransaction, errTransaction := mapper.DataDbToEntityTransaction(transaction.DTOTransaction{
			Id:                idTrans,
			VoucherCustomerId: voucherCustomerId,
			CustomerId:        customerId,
			CodeTransaction:   codeTransaction,
			Tanggalpembelian:  tanggalPembelian.Format("02-01-2006 15:04:05"),
			Total:             total,
			HargaDiscount:     hargaDiscount,
			TotalHarga:        totalHarga,
		})

		if errTransaction != nil {
			return nil, errTransaction
		}

		dataTransactionColletion = append(dataTransactionColletion, dataTransaction)

	}
	defer rows.Close()

	return dataTransactionColletion, nil
}

func (t *TransactionRepositoryMysqlInteractor) StoreTransaction(ctx context.Context, dataTransaction *transaction.Transaction) error {
	fmt.Println(dataTransaction)
	return nil
}
