package bnrdbmigrator

import (
	"gorm.io/gorm"
)

type Client interface {
	MigrateDB(revisions []MigrationRevision) error
}

type defaultClient struct {
	cfg        Config
	db         *gorm.DB
	getBinData func(fileName string) ([]byte, error)
}

func NewClient(cfg Config, db *gorm.DB, getBinData func(path string) ([]byte, error)) Client {
	return &defaultClient{
		cfg:        cfg,
		db:         db,
		getBinData: getBinData,
	}
}
