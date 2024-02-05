// This file contains the repository implementation layer.
package repository

import (
	"database/sql"
	"fmt"

	"github.com/SawitProRecruitment/UserService/config"
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
