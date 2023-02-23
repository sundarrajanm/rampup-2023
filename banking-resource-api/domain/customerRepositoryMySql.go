package domain

import (
	"banking-resource-api/errs"
	"banking-resource-api/types"
	"banking-resource-api/utils"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const MySqlDriver = "mysql"

type CustomerRepoMySql struct {
	client *sqlx.DB
}

func (d CustomerRepoMySql) FindAll() ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	d.client.Select(&customers,
		"select cust_id, name, city, zipcode, dob, status from customers")
	return customers, nil
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
	dbClient, _ := openSql("mysql", connectionString)
	return CustomerRepoMySql{dbClient}
}
