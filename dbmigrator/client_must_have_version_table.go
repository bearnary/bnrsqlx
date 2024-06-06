package bnrdbmigrator

func (c *defaultClient) mustHaveVersionTable() error {
	err := c.db.AutoMigrate(&SchemaRevision{})
	if err != nil {
		return err
	}
	return nil
}
