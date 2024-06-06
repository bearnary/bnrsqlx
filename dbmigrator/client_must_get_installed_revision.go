package bnrdbmigrator

import "gorm.io/gorm"

func (c *defaultClient) mustGetInstalledRevision() (int64, error) {
	version := &SchemaRevision{}
	err := c.db.Model(version).
		Where(&SchemaRevision{ID: c.cfg.Prefix}).
		First(version).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	return version.Revision, nil
}
