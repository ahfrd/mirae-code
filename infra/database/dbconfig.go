package database

import (
	"database/sql"
	"errors"
	"mirae-code/env"

	_ "github.com/go-sql-driver/mysql"
)

type Env struct {
	Params *env.EnvironmentParameters
	Error  error
}

func NewMySQLDB(config env.EnvironmentParameters) (*sql.DB, error) {
	conn, err := sql.Open(config.Database.MySQLDB.DBType, config.SetupMySQLDBConnection().Database.MySQLDB.DBConfig)
	if err != nil {
		return conn, errors.New("failed connecting to database : " + err.Error())
	}

	if err = conn.Ping(); err != nil {
		return conn, errors.New("failed pinging to database : " + err.Error())
	}

	return conn, err
}
