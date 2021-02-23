package models

import (
    "database/sql"

    //mysql driver
    _ "github.com/go-sql-driver/mysql"
)

// DB contains connection to mysql database
var DB *sql.DB

// OpenConnectionDataBase makes connect to database
func OpenConnectionDataBase(dsn string) error {
    database, err := sql.Open(`mysql`, dsn)
    if err != nil {
        return err
    }

    DB = database
    return database.Ping()
}

// CloseConnectionDataBase closes connection to database
func CloseConnectionDataBase() error {
    return DB.Close()
}
