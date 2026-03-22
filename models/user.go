package models

import (
	"github.com/nitesh-mhatre/go-rest-api/db"
	"github.com/nitesh-mhatre/go-rest-api/utils"
)


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

	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

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

func GetUserByUsername(username string) (*User, error) {
	quary := `SELECT id, username, password FROM users WHERE username = ?`
	row := db.DB.QueryRow(quary, username)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CheckPasswordHash(password, hash string) bool {
	return utils.CheckPasswordHash(password, hash)
}
