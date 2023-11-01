package main

import (
	"bytes"
	"database/sql"
	"fmt"

	"github.com/MrYanev/Gophercises/Exercise-8/phone/db"
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
	must(db.Reset("postgres", psqlInfo, dbname))

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	must(phonedb.Migrate("postgres", psqlInfo))

	db, err := sql.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	err = db.Seed()
	must(err)
	// must(db.Ping())
	// must(createPNTable(db))
	// id, err := insertPhone(db, "1234567890")
	// must(err)
	// fmt.Println("id=", id)

	// db, err := phonedb.Open("postgres")
	// number, err := getPhone(db, id)
	// must(err)
	// fmt.Println("Number is ...", number)

	// phones, err := getAllPhones(db)
	// must(err)
	// for _, p := range phones {
	// 	fmt.Printf("%+v\n", p)
	// 	number := normalize(p.number)
	// 	if number != p.number {
	// 		fmt.Println("Updating or removing...", number)
	// 		existing, err := findPhone(db, number)
	// 		must(err)
	// 		if existing != nil {
	// 			must(deletePhone(db, id))
	// 		} else {
	// 			p.number = number
	// 			must(updatePhone(db, p))
	// 		}
	// 	} else {
	// 		fmt.Println("No changes required")
	// 	}
	// }
}

func getPhone(db *sql.DB, id int) (string, error) {
	var number string
	err := db.QueryRow("SELECT * FROM phone_numbers WHERE id=$1", id).Scan(&id, &number)
	if err != nil {
		return "", err
	}
	return number, nil
}

func findPhone(db *sql.DB, number string) (*phone, error) {
	var p phone
	err := db.QueryRow("SELECT * FROM phone_numbers WHERE id=$1", number).Scan(&p.id, &p.number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &p, nil
}

func updatePhone(db *sql.DB, p phone) error {
	statement := `UPDATE phone_numbers SET value=$2 WHERE id=$1`
	_, err := db.Exec(statement, p.id, p.number)
	return err
}

func deletePhone(db *sql.DB, id int) error {
	statement := `DELETE FROM phone_numbers WHERE id=$1`
	_, err := db.Exec(statement, id)
	return err
}

type phone struct {
	id     int
	number string
}

func getAllPhones(db *sql.DB) ([]phone, error) {
	rows, err := db.Query("SELECT id, value FROM phone_numbers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ret []phone
	for rows.Next() {
		var p phone
		if err := rows.Scan(&p.id, &p.number); err != nil {
			return nil, err
		}
		ret = append(ret, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ret, nil
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
