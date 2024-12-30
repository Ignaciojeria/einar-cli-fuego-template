package configuration

import (
	ioc "github.com/Ignaciojeria/einar-ioc/v2"
)

func init() {
	ioc.Registry(NewPostgreSQLConfiguration)
}

type PostgreSQLConfiguration struct {
	DATABASE_POSTGRES_HOSTNAME string `env:"DATABASE_POSTGRES_HOSTNAME,required"`
	DATABASE_POSTGRES_PORT     string `env:"DATABASE_POSTGRES_PORT,required"`
	DATABASE_POSTGRES_NAME     string `env:"DATABASE_POSTGRES_NAME,required"`
	DATABASE_POSTGRES_USERNAME string `env:"DATABASE_POSTGRES_USERNAME,required"`
	DATABASE_POSTGRES_PASSWORD string `env:"DATABASE_POSTGRES_PASSWORD,required"`
	DATABASE_POSTGRES_SSL_MODE string `env:"DATABASE_POSTGRES_SSL_MODE,required"`
}

func NewPostgreSQLConfiguration() (PostgreSQLConfiguration, error) {
	return Parse[PostgreSQLConfiguration]()
}
