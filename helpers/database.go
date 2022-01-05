package database

import (
	"log"

	"github.com/jackc/pgx"
)

var pool *pgx.ConnPool

func connect() {
	pgxcfg, err := pgx.ParseURI("postgres://postgres:password@localhost:5432/go_chatapp")
	if err != nil {
		log.Fatalf("Parse URI error: %s", err)
	}

	pool, err = pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     pgxcfg,
		MaxConnections: 20,
		AfterConnect: func(conn *pgx.Conn) error {
			_, err := conn.Prepare("getPackage", `SELECT * FROM users`)
			return err
		},
	})
	if err != nil {
		log.Fatalf("Connection error: %s", err)
	}

}
