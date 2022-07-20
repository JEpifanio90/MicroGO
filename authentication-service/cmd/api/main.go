package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	conn, err := openDB()

	if err != nil {
		log.Panic("Can't connect to Postgres!")
	}

	app := Config{
		DB: conn,
	}

	log.Printf("Starting Authentication service on port %s\n", webPort)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", os.Getenv("DSN"))

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
