package domain

import (
	"backend/errs"
	"backend/logger"
	"database/sql"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (d AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "insert into accounts (cust_id, opening_date, account_type, amount, status) values (?,?,?,?,?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("an error occured while inserting into account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected datbase error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("an error occured while getting account id from result: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected datbase error")
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func (d AccountRepositoryDB) FindBy(id string) (*Account, *errs.AppError) {
	findAccountByIdSql := "select * from accounts where account_id = ?"
	var a Account
	err := d.client.Get(&a, findAccountByIdSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("account with Id: " + id + " not found")
		} else {
			logger.Error("Error while scanning account " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &a, nil
}

func (d AccountRepositoryDB) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while creating new transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected databse error")
	}

	insertTransactionSql := "insert into transactions (account_id, amount, tx_type, tx_date) values (?,?,?,?)"
	result, _ := tx.Exec(insertTransactionSql, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	if t.IsWithdrawal() {
		updateAccountSql := "update accounts set amount = amount - ? where account_id = ?"
		_, err = tx.Exec(updateAccountSql, t.Amount, t.AccountId)
	} else {
		updateAccountSql := "update accounts set amount = amount + ? where account_id = ?"
		_, err = tx.Exec(updateAccountSql, t.Amount, t.AccountId)
	}

	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected databse error")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while committing transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected databse error")
	}

	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	account, appError := d.FindBy(t.AccountId)
	if appError != nil {
		return nil, appError
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount
	return &t, nil
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{dbClient}
}
