package bnrdbmigrator

import (
	"fmt"

	"gorm.io/gorm"
)

func (c *defaultClient) canRecordUpdatingRevision(revision int64) (tx *gorm.DB, err error) {
	defer func() {
		if err != nil && tx != nil {
			tx.Rollback()
			tx = nil
		}
	}()

	r := &SchemaRevision{}
	tx = c.db.Begin()
	statement := `
INSERT INTO %s (id, revision)
VALUES ('%s-updating', %d)
ON CONFLICT (id) DO UPDATE SET
    revision = EXCLUDED.revision
`
	err = tx.Exec(fmt.Sprintf(statement, r.TableName(), c.cfg.Prefix, revision)).
		Error
	return
}
