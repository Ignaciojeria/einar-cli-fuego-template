# PostgreSQL Repository Creation Rules

This document outlines the rules for creating PostgreSQL repositories in the Einar framework.

## 1. File Structure & Location

**Directory:** `/app/adapter/out/postgresql_repository/`

**File Naming:** `snake_case.go` (e.g., `save_customer.go` for component `save-customer`)

## 2. Code Structure

Every repository file must follow this pattern. Replace `RunPostgreSQLOperation` with the PascalCase name of the component (e.g., `SaveCustomer`) and update the constructor accordingly.

```go
package postgresql_repository

import (
	"<module-name>/app/shared/infrastructure/postgresql"
	"context"

	ioc "github.com/Ignaciojeria/einar-ioc/v2"
	"gorm.io/gorm"
)

// Replace RunPostgreSQLOperation with the component name (e.g. SaveCustomer)
type RunPostgreSQLOperation func(ctx context.Context, input interface{}) error

func init() {
	ioc.Registry(
		NewRunPostgreSQLOperation,
		postgresql.NewConnection)
}

func NewRunPostgreSQLOperation(connection *gorm.DB) RunPostgreSQLOperation {
	return func(ctx context.Context, input interface{}) error {
		return nil
	}
}
```

## 3. Infrastructure Dependencies

Repositories depend on the PostgreSQL connection and configuration. The LLM **MUST** ensure these files exist.

### A. Connection File

**Location:** `/app/shared/infrastructure/postgresql/connection.go`

**Rule:** Check if this file exists. If it does **NOT** exist, create it with the following content. If it **DOES** exist, **DO NOT** modify it.

```go
package postgresql

import (
	"<module-name>/app/shared/configuration"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

func init() {
	ioc.Registry(
		NewConnection,
		configuration.NewPostgreSQLConfiguration,
	)
}

func NewConnection(env configuration.PostgreSQLConfiguration) (*gorm.DB, error) {

	// 1️⃣ Si DATABASE_URL está seteado → usarlo sí o sí
	if env.DATABASE_URL != "" {

		// Warning elegante: ambos presentes
		if env.DATABASE_POSTGRES_HOSTNAME != "" ||
			env.DATABASE_POSTGRES_USERNAME != "" ||
			env.DATABASE_POSTGRES_PASSWORD != "" ||
			env.DATABASE_POSTGRES_NAME != "" {

			log.Println("[config warning] DATABASE_URL is set and overrides individual Postgres variables")
		}

		db, err := gorm.Open(postgres.Open(env.DATABASE_URL))
		if err != nil {
			return nil, err
		}

		_ = db.Use(tracing.NewPlugin())
		return db, nil
	}

	// 2️⃣ Resolver DSN manualmente si no tienes DATABASE_URL
	username := env.DATABASE_POSTGRES_USERNAME
	password := env.DATABASE_POSTGRES_PASSWORD
	host := env.DATABASE_POSTGRES_HOSTNAME
	port := env.DATABASE_POSTGRES_PORT
	dbname := env.DATABASE_POSTGRES_NAME
	sslMode := env.DATABASE_POSTGRES_SSL_MODE

	dsn := "postgres://" + username + ":" + password +
		"@" + host + ":" + port + "/" + dbname +
		"?sslmode=" + sslMode

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	_ = db.Use(tracing.NewPlugin())
	return db, nil
}
```

### B. Configuration File

**Location:** `/app/shared/configuration/postgresql_conf.go`

**Rule:** Check if this file exists. If it does **NOT** exist, create it with the following content. If it **DOES** exist, **DO NOT** modify it.

```go
package configuration

import (
	ioc "github.com/Ignaciojeria/einar-ioc/v2"
)

func init() {
	ioc.Registry(NewPostgreSQLConfiguration)
}

type PostgreSQLConfiguration struct {
	DATABASE_URL               string `env:"DATABASE_URL"`
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
```

## 4. Mandatory Registration in main.go

> [!IMPORTANT]
> **When creating a PostgreSQL Repository, the LLM MUST ensure that `main.go` contains a blank import for the postgresql infrastructure package.**

**Required Import:**

```go
_ "<module-name>/app/shared/infrastructure/postgresql"
```

**Why?**
Without this import, the connection provider will not be registered in the IoC system.
