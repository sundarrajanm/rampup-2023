package types

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type HttpListenAndServe func(string, http.Handler) error
type OpenSqlxDB func(string, string) (*sqlx.DB, error)
