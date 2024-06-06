package bnrdbmigrator

import (
	"time"
)

type SchemaRevision struct {
	ID        string `gorm:"primary_key"`
	Revision  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (SchemaRevision) TableName() string {
	return "schema_revisions"
}
