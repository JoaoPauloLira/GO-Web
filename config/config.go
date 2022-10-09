package config

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
)

func GetDB() (db *sql.DB, err error) {
	db, err = sql.Open("mssql", "serve=localhost;user id=sa;password=sa; database=Loja")
	return
}
