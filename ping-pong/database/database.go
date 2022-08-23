package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/hiiamtrong/ping-ping/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", config.Cfg.PGUser, config.Cfg.PGPassword, config.Cfg.PGHost, config.Cfg.PGPort, config.Cfg.PGDB)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	createCounterQuery := "CREATE TABLE IF NOT EXISTS counter(id SERIAL primary key, count int)"
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	_, err = db.ExecContext(ctx, createCounterQuery)

	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
