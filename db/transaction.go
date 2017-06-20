package db

import (
	"database/sql"

	"gitlab.com/parkby/parkby-service/log"
)

type action func(tx *sql.Tx) error

type transactionInvoker struct {
}

// TransactionInvoker transaction invoker
var TransactionInvoker transactionInvoker

// Invoke invoke action
func (txInvoker transactionInvoker) Invoke(act action) error {
	var myDB = GetDB()
	tx, _ := myDB.Begin()

	err := act(tx)

	if err != nil {
		tx.Rollback()
		log.Error.PrintlnErr(err)
	} else {
		tx.Commit()
	}

	return err
}
