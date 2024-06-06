package bnrsqlx

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func (c *defaultClient) ConnectMySQL() error {

	var dataSource = fmt.Sprintf(`%v:%v@(%v:%v)/%v?parseTime=true&locAsia%%2FBangkok`, c.cfg.DatabaseUsername, c.cfg.DatabasePassword, c.cfg.DatabaseUrl, c.cfg.DatabasePort, c.cfg.DatabaseName)

	db, err := sqlx.Connect("mysql", dataSource)
	if err != nil {
		return err
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	c.db = db
	c.dbType = DatabaseTypeMySQL
	return nil
}
