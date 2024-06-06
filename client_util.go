package bnrsqlx

import "github.com/jmoiron/sqlx"

func (c *defaultClient) DB() *sqlx.DB {
	return c.db
}

func (c *defaultClient) Close() error {
	return c.db.Close()
}
