package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

func GetConnection() *sqlx.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	dataSource := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, name,
	)
	log.Info("connecting to db: ", dataSource)
	conn, err := sqlx.Open("postgres", dataSource)
	if err != nil {
		log.Fatal(err)
	}

	if err := conn.PingContext(context.Background()); err != nil {
		log.Fatal(err)
	}

	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(time.Minute * 5)

	return conn
}
