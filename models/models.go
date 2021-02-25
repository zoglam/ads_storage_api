package models

import (
    "database/sql"
    "fmt"

    config "github.com/zoglam/ads_storage_api/config"

    //mysql driver
    _ "github.com/go-sql-driver/mysql"
)

// DB contains connection to mysql database
var DB *sql.DB

// OpenConnectionDataBase makes connect to database
func OpenConnectionDataBase() error {
    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s",
        config.Params.Maria.DBUser,
        config.Params.Maria.DBPass,
        config.Params.Maria.DBHost,
        config.Params.Maria.DBPort,
        config.Params.Maria.DBName,
    )
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
