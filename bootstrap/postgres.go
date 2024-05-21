package bootstrap

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func InitPostgresDatabase(cfg *Config) *Storage  {
	const op = "postgres.InitPostgresDatabase"

	dbHost := cfg.Postgres.PostgresHost
	dbPort := cfg.Postgres.PostgresPort
	dbUser := cfg.Postgres.PostgresUser
	dbPasswd := cfg.Postgres.PostgresPassword
	dbName := cfg.Postgres.DatabaseName

	postgresUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",dbHost, dbPort, dbUser, dbPasswd, dbName)
	db, err := sql.Open("postgres", postgresUrl)
	if err != nil {
		log.Fatalf("%s: %v", op, err)
	}
	createDatabase, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "user" (
		id UUID PRIMARY KEY,
		username VARCHAR(20) NOT NULL UNIQUE,
		password CHAR(60) NOT NULL UNIQUE
	);`)
	if err != nil {
		log.Fatalf("%s: %v", op, err)
	}
	_, err = createDatabase.Exec()
	if err != nil {
		log.Fatalf("%s: %v", op, err)
	}
	return &Storage{db: db}
}

func CloseDB(db *Storage) error {
	return db.db.Close()
}
