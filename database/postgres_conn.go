package app

import (
	"arabian-snooker/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectToPostgres() (*sql.DB, error) {
	log.Printf("Connecting to Postgres with parameters: host=%s port=%d user=%s dbname=%s sslmode=%s",
		config.Cfg.Host, config.Cfg.PGPort, config.Cfg.PGUser, config.Cfg.PGDB, config.Cfg.SSLMode)

	dbconnection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Cfg.Host, config.Cfg.PGPort, config.Cfg.PGUser, config.Cfg.PGPassword, config.Cfg.PGDB, config.Cfg.SSLMode)

	db, err := sql.Open("postgres", dbconnection)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	return db, nil
}
