package bnrsqlx

import "fmt"

type WhereQueryBuilder interface {
	Where(fieldName string, value interface{}, op WhereOperator)
	BuildQuery() (string, []interface{})
}

type defaultWhereQueryBuilder struct {
	whereQuery string
	args       []interface{}
}

// NewWhereQueryBuilder create where query builder
func NewWhereQueryBuilder() WhereQueryBuilder {
	return &defaultWhereQueryBuilder{
		args: []interface{}{},
	}
}

// Where add wher clause
func (b *defaultWhereQueryBuilder) Where(fieldName string, value interface{}, op WhereOperator) {
	if b.whereQuery != "" {
		b.whereQuery = fmt.Sprintf("%v and %v %v $%d", b.whereQuery, fieldName, op, len(b.args)+1)
	} else {
		b.whereQuery = fmt.Sprintf("%v %v $%d", fieldName, op, len(b.args)+1)
	}

	switch op {
	case WhereOperatorLike:
		value = fmt.Sprintf("%%%v%%", value)
	default:
	}

	b.args = append(b.args, value)
}

// BuildQuery create where query string and return with arguments
func (b *defaultWhereQueryBuilder) BuildQuery() (string, []interface{}) {
	return b.whereQuery, b.args
}
