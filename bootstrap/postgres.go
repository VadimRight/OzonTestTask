package bootstrap

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

func InitPostgresDatabase(cfg *Config)  {
	const op = "postgres.InitPostgresDatabase"

	dbHost := cfg.Postgres.PostgresHost
	dbPort := cfg.Postgres.PostgresPort
	dbUser := cfg.Postgres.PostgresUser
	dbPasswd := cfg.Postgres.PostgresPassword
	dbName := cfg.Postgres.DatabaseName

	postgresUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s TimeZone=Asia/Tehran",
		dbHost, dbPort, dbUser, dbPasswd, dbName)
	db, err := sql.Open("postgres", postgresUrl)
	if err != nil {
		log.Fatalf("Error while connecting to postgres database: %v", err)	
	}
	createDatabase, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "user" (
		id UUID PRIMARY KEY,
		username VARCHAR(20) NOT NULL UNIQUE,
		email VARCHAR(20) NOT NULL UNIQUE,
		password CHAR(60) NOT NULL UNIQUE,
		is_verified BOOL NOT NULL DEFAULT false,
		is_activate BOOL NOT NULL DEFAULT false
	);`)
	if err != nil {
		log.Fatalf("%s: %v", op, err)
	}
	_, err = createDatabase.Exec()
	if err != nil {
		log.Fatalf("%s: %v", op, err)
	}
	defer db.Close()
}
