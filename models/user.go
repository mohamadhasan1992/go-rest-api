package models

import (
	"errors"

	"github.com/mohamadhasan1992/go-rest-api.git/db"
	"github.com/mohamadhasan1992/go-rest-api.git/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users(email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userId, err := res.LastInsertId()
	u.Id = userId
	return err
}

func (user *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)
	var retrievedPassword string
	err := row.Scan(&user.Id, &retrievedPassword)
	if err != nil {
		return err
	}
	isPasswrodValid := utils.CompareHash(user.Password, retrievedPassword)
	if isPasswrodValid {
		return nil
	} else {
		return errors.New("Credentials are not valid!")
	}
}
