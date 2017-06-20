package model

import (
	"database/sql"
	"errors"

	"gitlab.com/parkby/parkby-service/log"
)

// User model
type User struct {
	Model
	ID          int64
	LastName    string
	FirstName   string
	PhoneNumber string
	Email       string
	Password    string
}

var (
	insertstmt = "insert into users (last_name, first_name, phone_number, email, password) values (?,?,?,?,?)"
	// updatestmt = "update users set name = ? where id = ?"
)

// Save saves *user
func (user *User) Save(tx *sql.Tx) error {
	res, err := tx.Exec(insertstmt,
		user.LastName,
		user.FirstName,
		user.PhoneNumber,
		user.Email,
		user.Password,
	)

	if err != nil {
		log.Error.PrintlnErr(err)
		return errors.New("fail saving *User")
	}

	userID, _ := res.LastInsertId()
	user.ID = userID

	return nil
}

// Update updates user
func (user User) Update(tx *sql.Tx) error {
	log.Warn.Println(">>> NO Implementation")
	return nil
}

// EqualPassword equal passwords
func (user User) EqualPassword(password string) bool {
	return user.Password == password
}
