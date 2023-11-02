package main

import (
	"bytes"
	"fmt"

	phonedb "github.com/MrYanev/Gophercises/Exercise-8/phone/db"
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
	must(phonedb.Reset("postgres", psqlInfo, dbname))

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	must(phonedb.Migrate("postgres", psqlInfo))

	db, err := phonedb.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	err = db.Seed()
	must(err)

	phones, err := db.AllPhones()
	must(err)
	for _, p := range phones {
		fmt.Printf("Working on... %+v\n", p)
		number := normalize(p.Number)
		if number != p.Number {
			fmt.Println("Updating or removing...", number)
			existing, err := db.FindPhone(number)
			must(err)
			if existing != nil {
				must(db.DeletePhone(p.ID))
			} else {
				p.Number = number
				must(db.UpdatePhone(&p))
			}
		} else {
			fmt.Println("No changes required")
		}
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
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
