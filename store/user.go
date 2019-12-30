package store

import (
	"fmt"
	db "todo_app/db"
	"todo_app/types"
)

func GetUserByEmailAddress(emailAddress string) types.User {
	var user types.User
	for _, usr := range db.UserDB {
		if usr.EmailAddress == emailAddress {
			user = usr
			break
		}
	}
	return user
}

func AddUser(newUser types.User) {
	newUser.ID = len(db.UserDB) + 1
	db.UserDB = append(db.UserDB, newUser)
	fmt.Println("After New user added: ", db.UserDB)
}

func UploadMockData() {
	db.SeedData()
}

type MyUser struct {
	*types.User
}

func (u *MyUser) GetUserByEmailAndPassword(emailAddress string, password string) types.User {
	var selUsr types.User
	for _, usr := range db.UserDB {
		if usr.EmailAddress == emailAddress && usr.Password == password {
			selUsr = usr
			break
		}
	}
	return selUsr
}