package domain

import (
	"backend/errs"
	"backend/logger"
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

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{dbClient}
}
