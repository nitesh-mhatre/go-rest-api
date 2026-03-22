package models

import "github.com/nitesh-mhatre/go-rest-api/db"

type User struct {
	ID int
	Username string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	quary := `INSERT INTO users (username, password) VALUES (?, ?)`
	stmt , err := db.DB.Prepare(quary)
	
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.Username, u.Password)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = int(id)

	return nil

}

