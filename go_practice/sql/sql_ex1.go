package main

import (
	"database/sql"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	})
	db, err := sql.Open("mysql", "saechim$daeki@tcp(127.0.0.1:3306)/saechim")
	if err != nil {
		log.Fatal().Err(err).Msg("[sql.Open] failed")
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal().Err(err).Msg("[db.Close] failed")
		}
	}()

	var name, email string
	row := db.QueryRow("SELECT name, email from users where id = ?", 1)
	err = row.Scan(&name, &email)
	if err != nil {
		log.Fatal().Err(err).Msg("[db.Scan] failed")
	}
	fmt.Printf("name :%q, email :%q\n", name, email)
}
