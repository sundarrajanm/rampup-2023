package domain

import (
	"backend/errs"
	"backend/logger"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, *errs.AppError) {
	findAllSql := "select cust_id, name, city, zipcode, dob, status from customers"
	customers := make([]Customer, 0)
	err := d.client.Select(&customers, findAllSql)

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	findByIdSql := "select cust_id, name, city, zipcode, dob, status from customers where cust_id = ?"
	var c Customer
	err := d.client.Get(&c, findByIdSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer with Id: " + id + " not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWD")
	host := os.Getenv("DB_ADDR")
	port := os.Getenv("DB_PORT")
	db := os.Getenv("DB_NAME")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, db)
	client, err := sqlx.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client}
}
