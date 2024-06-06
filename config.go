package bnrsqlx

type Config struct {
	DatabaseUrl            string `env:"DATABASE_URL" mapstructure:"database_url"`
	DatabasePort           string `env:"DATABASE_PORT" mapstructure:"database_port"`
	DatabaseName           string `env:"DATABASE_NAME" mapstructure:"database_name"`
	DatabaseUsername       string `env:"DATABASE_USERNAME" mapstructure:"database_username"`
	DatabasePassword       string `env:"DATABASE_PASSWORD" mapstructure:"database_password"`
	DatabaseLogEnabled     bool   `env:"DATABASE_LOG_ENABLED" mapstructure:"database_log_enabled"`
	DatabaseDisableSSLMode bool   `env:"DATABASE_DISABLE_SSL_MODE" mapstructure:"database_disable_ssl_mode"`
}
