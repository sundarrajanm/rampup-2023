package domain

import (
	"banking-resource-api/errs"
	"banking-resource-api/logger"
	"banking-resource-api/types"
	"banking-resource-api/utils"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const MySqlDriver = "mysql"

type CustomerRepoMySql struct {
	client *sqlx.DB
}

func (d CustomerRepoMySql) FindAll() ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)

	err := d.client.Select(&customers,
		"select cust_id, name, city, zipcode, dob, status from customers")
	if err != nil {
		logger.Error("Error while querying database: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepoMySql) FindById(id string) (*Customer, *errs.AppError) {
	findByIdSql := "select cust_id, name, city, zipcode, dob, status from customers where cust_id = ?"
	var c Customer
	err := d.client.Get(&c, findByIdSql, id)

	if err == sql.ErrNoRows {
		return nil, errs.NewNotFoundError(fmt.Sprintf("Customer with Id: '%s' not found", id))
	}

	if err != nil {
		logger.Error("Error while getting customer: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &c, nil
}

func GetConnectionString() string {
	user := utils.CheckMandatoryEnvVar("DB_USER")
	password := utils.CheckMandatoryEnvVar("DB_PASSWORD")
	host := utils.CheckMandatoryEnvVar("DB_HOST")
	port := utils.CheckMandatoryEnvVar("DB_PORT")
	db := utils.CheckMandatoryEnvVar("DB_NAME")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, db)
}

func NewCustomerRepoMySql(openSql types.OpenSqlxDB) CustomerRepository {
	connectionString := GetConnectionString()
	dbClient, err := openSql(MySqlDriver, connectionString)

	if err != nil {
		logger.Error("Panicing due to OpenSql failure: " + err.Error())
		panic(err.Error())
	}
	return CustomerRepoMySql{dbClient}
}
