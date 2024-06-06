package bnrsqlx

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func (c *defaultClient) ConnectPostgres() error {
	sslMode := "require"
	if c.cfg.DatabaseDisableSSLMode {
		sslMode = "disable"
	}
	var dataSource = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", c.cfg.DatabaseUrl, c.cfg.DatabasePort, c.cfg.DatabaseUsername, c.cfg.DatabasePassword, c.cfg.DatabaseName, sslMode)
	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		return err
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	c.db = db
	c.dbType = DatabaseTypePostgres
	return c.db.Ping()
}
