package bnrdbmigrator

func (c *defaultClient) mustRecordInstalledRevision(revision int64) error {
	r := &SchemaRevision{}
	err := c.db.
		Where(&SchemaRevision{ID: c.cfg.Prefix}).
		Assign(&SchemaRevision{Revision: revision}).
		FirstOrCreate(&r).
		Error

	if err != nil {
		return err
	}
	return nil
}
