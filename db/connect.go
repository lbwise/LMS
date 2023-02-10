package db

import (
	"errors"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

const (
	host	= "localhost"
	port	= 5432
	user	= "postgres"
	pass	= "Chromium24"
	dbname	= "lms"
)

func ConnectDB() (*sql.DB, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	DB, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err.Error())
	}
	err = DB.Ping()
	if err != nil {
		err = errors.New(err.Error())
		return nil, err
	}
	fmt.Println("----------- DB CONNECTED ------------")
	return DB, nil
}

func ResetDB(db *sql.DB) {
	execSQL("./db/sql/reset.sql", db)
	execSQL("./db/sql/user_seed_data.sql", db)
	execSQL("./db/sql/courses_seed_data.sql", db)
}