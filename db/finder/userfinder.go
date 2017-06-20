package finder

import (
	"gitlab.com/parkby/parkby-service/db"
	"gitlab.com/parkby/parkby-service/log"
	"gitlab.com/parkby/parkby-service/model"
)

// UserFinder user finder
type userFinder struct {
}

// UserFinder user finder
var UserFinder userFinder

// Find find user by ID
func (finder userFinder) Find(id int64) (user model.User, err error) {
	myDB := db.GetDB()

	user = model.User{}
	err = myDB.QueryRow(
		"select * from users where id=?", id,
	).Scan(
		&user.ID,
		&user.LastName,
		&user.FirstName,
		&user.PhoneNumber,
		&user.Email,
	)

	log.Error.PrintlnErr(err)

	return
}

func (finder userFinder) FindWithEmail(email string) model.User {
	myDB := db.GetDB()

	var user model.User

	err := myDB.QueryRow(
		"select * from users where email=?", email,
	).Scan(
		&user.ID,
		&user.LastName,
		&user.FirstName,
		&user.PhoneNumber,
		&user.Email,
		&user.Password,
	)

	log.Error.PrintlnErr(err)
	return user
}
