package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PgConn struct {
	db *sql.DB
}

func NewConn(c *inputConf) (*PgConn, error) {

	conn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		c.user, c.database, c.password, c.hostname, c.port)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PgConn{
		db: db,
	}, nil
}

