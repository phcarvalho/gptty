package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/phcarvalho/gptty/internal/models"
	_ "modernc.org/sqlite"
)

type application struct {
	systems *models.SystemModel
}

func main() {
	port := 4000
	addr := fmt.Sprintf(":%d", port)

	dsn := "./data.db"

	db, err := openDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	app := &application{
		systems: &models.SystemModel{DB: db},
	}

	srv := http.Server{
		Addr:    addr,
		Handler: app.routes(),
	}

	fmt.Printf("Starting server on port %d", port)
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
