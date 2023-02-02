package belajar_golang_datatabase

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql" // driver registration (only run init func)
)

func TestOpenConnection(t *testing.T) {
	// create database pool (pool of connection)
	// minimum mysql connection is ~150
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_golang_database")
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
