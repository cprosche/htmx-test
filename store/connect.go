package store

import (
	"errors"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectToDb() (*sqlx.DB, error) {
	dbUser := os.Getenv("PGUSER")
	dbPassword := os.Getenv("PGPASSWORD")
	dbName := os.Getenv("PGDATABASE")

	if dbUser == "" || dbPassword == "" || dbName == "" {
		return nil, errors.New("missing database credentials")
	}

	return sqlx.Connect(
		"postgres",
		"user="+dbUser+
			" password="+dbPassword+
			" dbname="+dbName+
			" sslmode=disable",
	)
}
