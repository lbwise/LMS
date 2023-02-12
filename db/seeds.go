package db

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/bcrypt"
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

func CreateUsers(db *sql.DB) {
	filePath := "./db/sql/user_data.csv"
	fi, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
	}
	reader := csv.NewReader(fi)
	var (
		query, pwd string
	)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(records)
	for _, record := range records[1:] {
		pwd = record[2]
		password, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
		query = fmt.Sprintf(`INSERT INTO users (name, email, password, created_on) VALUES ('%s', '%s', '%s', date '%s')`,
			record[0],
			record[1],
			password,
			record[3],
		)
		res, err := db.Exec(query)	
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(res)
	}
	fmt.Println("COMPLETED")
}