package container

import (
	"database/sql"
	"fmt"
	"log"
)

type container struct {
	postgres *sql.DB
}

var cont = container{}

func setDb(db *sql.DB) {
	if cont.postgres != nil {
		err := cont.postgres.Close()
		if err != nil {
			log.Println(err)
		}
	}

	cont.postgres = db
}

func DB() *sql.DB {
	return cont.postgres
}

func Disconnect() {
	setDb(nil)
}

func Connect(user string, password string, dbname string) bool {
	// Todo move to env config
	port := 5432
	host := "localhost"

	psqlconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Println(err)
		return false
	}

	// check db
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println("DB Connected!")
	setDb(db)

	return true
}
