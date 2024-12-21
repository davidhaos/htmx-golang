package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

type DBConn interface {
	Open() error
	Close() error
	Execute(query string) (*sql.Rows, error)
}

type MySQLConnection struct {
	db  *sql.DB
	dsn string //Data Source Name
}

func (conn *MySQLConnection) Open() error {
	var err error
	conn.db, err = sql.Open("mysql", conn.dsn)
	if err != nil {
		return fmt.Errorf("error opening mysql connection: %v", err)
	}
	if err := conn.db.Ping(); err != nil {
		return fmt.Errorf("error pinging database: %v", err)
	}
	fmt.Println("Successful connected to MySQL database")
	return nil
}

func (conn *MySQLConnection) Close() error {
	if err := conn.db.Close(); err != nil {
		return fmt.Errorf("error closing connection: %v", err)
	}
	fmt.Println("Connection closed!")
	return nil
}

func (conn *MySQLConnection) Execute(query string) (*sql.Rows, error) {
	rows, err := conn.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	return rows, nil
}
