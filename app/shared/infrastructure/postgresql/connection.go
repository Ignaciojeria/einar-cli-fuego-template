package postgresql

import (
	"archetype/app/shared/configuration"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

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
