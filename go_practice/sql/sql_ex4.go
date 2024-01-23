package main

import (
	"database/sql"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

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

	var users []User
	err = db.Table("users").Where("id < ?", 10).Find(&users)
	if err != nil {
		log.Fatal().Err(err).Msg("[db.Find] failed")
	}

	spew.Dump(users)

	newUser := User{
		Name:     "saechimdd",
		Email:    "saechim@saeda.com",
		Password: "!31234124",
	}

	n, err := db.Table("users").Insert(&newUser)
	if err != nil {
		log.Fatal().Err(err).Msg("[db.insert] failed")
	}
	fmt.Printf("%d rows inserted", n)
}
