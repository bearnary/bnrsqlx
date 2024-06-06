package bnrsqlx

func (c *defaultClient) DatabaseType() DatabaseType {
	return c.dbType
}
