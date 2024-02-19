// This file contains the infrastructure implementation layer.
package infrastructure

import (
	"database/sql"
	"fmt"

	"userservice/config"

	_ "github.com/lib/pq"
)

type PostgreConnection struct {
	Db *sql.DB
}

func NewPostgreConnection(c *config.DBConfig) *PostgreConnection {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", c.User, c.Password, c.Hostname, c.Port, c.DB)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	return &PostgreConnection{
		Db: db,
	}
}
