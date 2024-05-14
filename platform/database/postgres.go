package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	_ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
)

// PostgreSQLConnection func for connection to PostgreSQL database.
func PostgreSQLConnection() (*sqlx.DB, error) {
	// ini mau dipake kah? bebas dah lu atur, hardcode paling, contoh valuenya ada di git tutornye

	// Define database connection settings.
	// maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	// maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	// maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	// Define database connection for PostgreSQL.

	godotenv.Load(".env.local")

	UNAMEDB := os.Getenv("DB_USERNAME")
	PASSDB := os.Getenv("DB_PASSWORD")
	HOSTDB := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")
	DB_PARAMS := os.Getenv("DB_PARAMS")

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?%s", UNAMEDB, PASSDB, HOSTDB, DBNAME, DB_PARAMS)

	db, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	// Set database connection settings.
	// db.SetMaxOpenConns(maxConn)                           // the default is 0 (unlimited)
	// db.SetMaxIdleConns(maxIdleConn)                       // defaultMaxIdleConns = 2
	// db.SetConnMaxLifetime(time.Duration(maxLifetimeConn)) // 0, connections are reused forever

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}
