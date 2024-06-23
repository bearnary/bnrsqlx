package bnrsqlx

import (
	"fmt"
)

// DeleteExec execute delete query with where clause
func (c *defaultClient) DeleteExec(arg interface{}, whereQuery string, values ...interface{}) error {

	tableName, vErr := parseTableName(arg)
	if vErr != nil {
		return vErr
	}

	query := fmt.Sprintf("delete from %v where %v", tableName, whereQuery)
	_, err := c.db.Exec(query, values...)
	if err != nil {
		return err
	}
	return nil
}
