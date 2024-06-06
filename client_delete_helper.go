package bnrsqlx

import (
	"fmt"

	"github.com/oneononex/oolib/ooerrors"
)

// DeleteExec execute delete query with where clause
func (c *defaultClient) DeleteExec(arg interface{}, whereQuery string, values ...interface{}) ooerrors.Error {

	tableName, vErr := parseTableName(arg)
	if vErr != nil {
		return vErr
	}

	query := fmt.Sprintf("delete from %v where %v", tableName, whereQuery)
	_, err := c.db.Exec(query, values...)
	if err != nil {
		return ooerrors.NewDatabaseError(err)
	}
	return nil
}
