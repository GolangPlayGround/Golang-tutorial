package main

import (
	"database/sql"
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

	result, err := db.Exec(
		"insert into users VALUES (?,?,?,?)",
		3,
		"saechimdaeki",
		"saechim@daeki.com",
		"131224124124141",
	)
	if err != nil {
		log.Fatal().Err(err).Msg("[db.Exec] failed")
	}
	n, err := result.LastInsertId()
	if err != nil {
		log.Fatal().Err(err).Msg("[result.LastInsertId] failed")
	}
	log.Info().Int64("last_insert_id", n).Msg("[insert] succeed")
}
