package bnrsqlx

import (
	"fmt"
)

type SelectQueryBuilder struct {
	SelectQuery  string        `json:"select_query"`
	WhereQuery   string        `json:"where_query"`
	OrderByQuery *string       `json:"order_by_query"`
	Limit        *int64        `json:"limit"`
	Page         *int64        `json:"page"`
	Args         []interface{} `json:"args"`
	HasDeletedAt bool          `json:"has_deleted_at"`
}

// Get query normal select query and return first value found in 'dest' struct
func (c *defaultClient) Get(dest interface{}, query string, args ...interface{}) error {
	err := c.db.Get(dest, query, args...)
	if err != nil {
		return err
	}
	return nil
}

// Select do select query with select query builder
func (c *defaultClient) Select(model interface{}, dest interface{}, qb SelectQueryBuilder) error {

	tableName, err := parseTableName(model)
	if err != nil {
		return err
	}

	orderByQuery := ""
	if qb.OrderByQuery != nil {
		orderByQuery = fmt.Sprintf(" order by %v", *qb.OrderByQuery)
	}

	limitQuery := ""
	if qb.Page != nil && qb.Limit != nil {
		limit := *qb.Limit
		offset := ((*qb.Page - 1) * *qb.Limit)
		limitQuery = fmt.Sprintf(" limit %v offset %v", limit, offset)
	} else if qb.Limit != nil {
		limitQuery = fmt.Sprintf(" limit %v", *qb.Limit)
	}

	whereQuery := ""
	if qb.WhereQuery != "" {
		whereQuery = fmt.Sprintf(" where %v", qb.WhereQuery)
	}

	query := fmt.Sprintf("select %v from %v%v%v%v", qb.SelectQuery, tableName, whereQuery, orderByQuery, limitQuery)
	err = c.db.Select(dest, query, qb.Args...)
	if err != nil {
		return err
	}

	return nil
}

// Count count all data in database with where clause
func (c *defaultClient) Count(model interface{}, whereQuery string, args []interface{}) (int64, error) {

	if whereQuery != "" {
		whereQuery = fmt.Sprintf(" where %v", whereQuery)
	}

	tableName, err := parseTableName(model)
	if err != nil {
		return 0, err
	}

	var count int64
	query := fmt.Sprintf("select count(*) as count from %v%v", tableName, whereQuery)
	rows, err := c.db.Query(query, args...)
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}
	}

	return count, nil
}

func (c *defaultClient) SelectWithCount(model interface{}, dest interface{}, qb SelectQueryBuilder, pagination *PaginationRequest) (int64, error) {

	if qb.HasDeletedAt {
		if qb.WhereQuery != "" {
			qb.WhereQuery = fmt.Sprintf("%v and", qb.WhereQuery)
		}
		qb.WhereQuery = fmt.Sprintf("%v deleted_at is null", qb.WhereQuery)
	}

	count, err := c.Count(model, qb.WhereQuery, qb.Args)
	if err != nil {
		return 0, err
	}

	// if len(values) == 2 && values[0] != 0 && values[1] != 0 {
	// 	qb.Page = &values[0]
	// 	qb.Limit = &values[1]
	// }
	if pagination != nil {
		if pagination.Limit != 0 {
			qb.Limit = &pagination.Limit
		}
		if pagination.Page != 0 {
			qb.Page = &pagination.Page
		}
	}

	err = c.Select(model, dest, qb)
	if err != nil {
		return 0, err
	}

	return count, nil
}
