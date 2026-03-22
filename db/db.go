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

	user_table_quary := `CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL)`

	_, err := DB.Exec(user_table_quary)
	
	if err != nil {
		return err
	}

	events_table_quary := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time TEXT NOT NULL,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(events_table_quary)
	
	if err != nil {
		return err
	}

	return nil
}