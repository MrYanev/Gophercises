package main

import (
	"bytes"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 8080
	user     = ""
	password = ""
	dbname   = "phonercise"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", host, port, user, password)
	db, err := sql.Open("postgres", psqlInfo)
	must(err)
	err = resetDB(db, dbname)
	must(err)
	db.Close()

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	must(db.Ping())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func resetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}
	return createDB(db, name)
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name) //Has a sql injection flow don't use
	if err != nil {
		return err
	}
	return nil
}

func normalize(phone string) string {
	var buf bytes.Buffer
	for _, i := range phone {
		if i >= '0' && i <= '9' {
			buf.WriteRune(i)
		}
	}
	return buf.String()
}
