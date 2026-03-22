package db
 import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)


var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "./api.db")
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	return createTable()
}


func createTable() error {
	quary := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time TEXT NOT NULL,
		user_id INTEGER
	)`

	_, err := DB.Exec(quary)
	return err
}