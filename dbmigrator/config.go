package bnrdbmigrator

type Config struct {
	Prefix  string `mapstructure:"prefix"`
	SQLPath string `mapstructure:"sql_path"`
}
