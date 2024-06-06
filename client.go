package bnrsqlx

import (
	"github.com/jmoiron/sqlx"
	"github.com/oneononex/oolib/ooerrors"
)

type Client interface {
	DB() *sqlx.DB
	DatabaseType() DatabaseType
	ConnectPostgres() error
	ConnectMySQL() error

	NewUpdateHelper() *UpdateHelper

	DeleteExec(arg interface{}, whereQuery string, values ...interface{}) ooerrors.Error

	Get(dest interface{}, query string, args ...interface{}) ooerrors.Error
	Select(model interface{}, dest interface{}, qb SelectQueryBuilder) ooerrors.Error
	SelectOne(dest interface{}, qb SelectOneQueryBuilder) ooerrors.Error
	Count(model interface{}, whereQuery string, args []interface{}) (int64, ooerrors.Error)
	SelectWithCount(model interface{}, dest interface{}, qb SelectQueryBuilder, pagination *PaginationRequest) (int64, ooerrors.Error)

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
