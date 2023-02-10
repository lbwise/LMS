package db

import (
	"database/sql"
	"io/ioutil"
	"fmt"
)

func setUp() error {
	panic("HELP")
}

func TestQuery() {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	var out string
	testStmt := `SELECT name FROM users WHERE name='liam wise'`
	err = db.QueryRow(testStmt).Scan(&out)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}


func execSQL(filePath string, db *sql.DB) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err.Error())
	}
	query := string(data)
	res, err := db.Exec(query)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(res)
}
