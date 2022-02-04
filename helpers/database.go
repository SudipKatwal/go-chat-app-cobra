package helpers

import (
	"fmt"
	"log"

	"github.com/jackc/pgx"
)

var pool *pgx.ConnPool

type (
	LoginUser struct {
		Id       int
		Name     string
		Email    string
		Password string
	}
)

func Connect() {
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

func Login(email, password string) (LoginUser, error) {
	stmt := "SELECT * FROM users WHERE email=$1 AND password=$2;"
	rows := pool.QueryRow(stmt, email, password)
	var id int
	var name, dbEmail, dbPassword string
	switch err := rows.Scan(&id, &name, &dbEmail, &dbPassword); err {
	case pgx.ErrNoRows:
		return LoginUser{}, err
	case nil:
		return LoginUser{Id: id, Name: name, Email: email, Password: password}, nil
	default:
		// panic(err.Error())
		return LoginUser{}, err
	}
}

func Register(name, email, password string) {
	var id int

	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`

	err := pool.QueryRow(query, name, email, password).Scan(&id)

	if err != nil {
		panic(err)
	}
	fmt.Println("Register successful. Please login now!")
}
