package domain

import (
	"banking-resource-api/errs"
	"banking-resource-api/types"
	"banking-resource-api/utils"
	"fmt"
)

type CustomerRepoMySql struct{}

func (d CustomerRepoMySql) FindAll() ([]Customer, *errs.AppError) {
	return []Customer{}, nil
}

func NewCustomerRepoMySql(openSql types.OpenSqlxDB) CustomerRepository {
	user := utils.CheckMandatoryEnvVar("DB_USER")
	password := utils.CheckMandatoryEnvVar("DB_PASSWORD")
	host := utils.CheckMandatoryEnvVar("DB_HOST")
	port := utils.CheckMandatoryEnvVar("DB_PORT")
	db := utils.CheckMandatoryEnvVar("DB_NAME")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, db)
	openSql("mysql", connectionString)
	return CustomerRepoMySql{}
}
