package models

import (
	"rest-api/db"
	"rest-api/utils"
)

type User struct {
	ID       int
	Username string `binding:"required"`
	Password string `binding:"required"`
	Email    string
}

func (u User) Save() error {
	query := `INSERT INTO users (username, password, email) VALUES (?, ?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Username, hashedPassword, u.Email)

	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = int(userId)
	return err
}

func (u User) ValidateCredentials() (bool, error) {
	query := `SELECT id, email, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return false, err
	}

	return utils.CheckPasswordHash(u.Password, retrievedPassword), nil
}
