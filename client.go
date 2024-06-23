package bnrsqlx

import (
	"github.com/jmoiron/sqlx"
)

type Client interface {
	DB() *sqlx.DB
	DatabaseType() DatabaseType
	ConnectPostgres() error
	ConnectMySQL() error

	NewUpdateHelper() *UpdateHelper

	DeleteExec(arg interface{}, whereQuery string, values ...interface{}) error

	Get(dest interface{}, query string, args ...interface{}) error
	Select(model interface{}, dest interface{}, qb SelectQueryBuilder) error
	SelectOne(dest interface{}, qb SelectOneQueryBuilder) error
	Count(model interface{}, whereQuery string, args []interface{}) (int64, error)
	SelectWithCount(model interface{}, dest interface{}, qb SelectQueryBuilder, pagination *PaginationRequest) (int64, error)

	Close() error
}

type defaultClient struct {
	cfg    *Config
	db     *sqlx.DB
	dbType DatabaseType
}

func NewClient(cfg *Config) Client {

	return &defaultClient{
		cfg: cfg,
	}
}
