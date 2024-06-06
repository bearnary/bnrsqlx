package bnrsqlx

import (
	"fmt"

	"github.com/oneononex/oolib/ooerrors"
)

type SelectOneQueryBuilder struct {
	SelectQuery  string        `json:"select_query"`
	WhereQuery   string        `json:"where_query"`
	OrderByQuery *string       `json:"order_by_query"`
	Args         []interface{} `json:"args"`
	HasDeletedAt bool          `json:"has_deleted_at"`
}

func (c *defaultClient) SelectOne(dest interface{}, qb SelectOneQueryBuilder) ooerrors.Error {

	tableName, vErr := parseTableName(dest)
	if vErr != nil {
		return vErr
	}

	orderByQuery := ""
	if qb.OrderByQuery != nil {
		orderByQuery = fmt.Sprintf(" order by %v", *qb.OrderByQuery)
	}

	whereQuery := ""
	if qb.WhereQuery != "" {
		whereQuery = fmt.Sprintf(" where %v", qb.WhereQuery)
	}

	if qb.HasDeletedAt {
		if whereQuery != "" {
			whereQuery = fmt.Sprintf("%v and", whereQuery)
		} else {
			whereQuery = fmt.Sprintf("%v where ", whereQuery)
		}
		whereQuery = fmt.Sprintf("%v deleted_at is null", whereQuery)
	}

	query := fmt.Sprintf("select %v from %v%v%v limit 1", qb.SelectQuery, tableName, whereQuery, orderByQuery)

	err := c.db.Get(dest, query, qb.Args...)
	if err != nil {
		return ooerrors.NewDatabaseError(err)
	}

	return nil
}
