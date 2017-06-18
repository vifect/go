package test

import (
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		fmt.Println("db error")
	}

}
