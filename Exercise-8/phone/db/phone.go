package db

import "database/sql"

func Reset(driverName, dataSource, dbName string) error {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return err
	}
	err = resetDB(db, dbName)
	if err != nil {
		return err
	}
	return db.Close()
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

func Open(driverName, dataSource string) (*DB, error) {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

type DB struct {
	db *sql.DB
}

func (db *DB) Close() error {
	return db.db.Close()
}

func Migrate(driverName, dataSource string) error {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return err
	}
	err = createPNTable(db)
	if err != nil {
		return err
	}
	return db.Close()
}

func insertPhone(db *sql.DB, phone string) (int, error) {
	statement := `INSERT INTO phone_numbers(value) VALUES($1)`
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func createPNTable(db *sql.DB) error {
	statement := `
	CREATE TABLE IF NOT EXISTS phone_numbers (
		id SERIAL,
		value VARCHAR(255),
	)`
	_, err := db.Exec(statement)
	return err
}

func (db *DB) Seed() error {
	data := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}
	for _, number := range data {
		if _, err := insertPhone(db.db, number); err != nil {
			return err
		}
	}
	return nil
}

//Not being used yet
func getPhone(db *sql.DB, id int) (string, error) {
	var number string
	err := db.QueryRow("SELECT * FROM phone_numbers WHERE id=$1", id).Scan(&id, &number)
	if err != nil {
		return "", err
	}
	return number, nil
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
