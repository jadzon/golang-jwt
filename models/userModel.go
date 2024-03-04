package models

import (
	"errors"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}
type Users struct {
	Data []User
}

func (u *Users) GetUserByID(id int) (User, error) {
	for _, user := range u.Data {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}
func (u *Users) GetUserByUsername(username string) (User, error) {
	for _, user := range u.Data {
		if user.Username == username {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}
func (u *Users) AddUserToDB(user User) error {
	if _, err := u.GetUserByID(user.ID); err == nil {
		return errors.New("id has to be unique")
	}
	if _, err := u.GetUserByUsername(user.Username); err == nil {
		return errors.New("username has to be unique")
	}
	u.Data = append(u.Data, user)
	return nil
}
func (u *Users) NumberOfUsers() int {
	return len(u.Data)
}

// func CreateEmptyUserDB() Users {
// 	var temp Users
// 	return temp
// }

func CreateEmptyUserDB() *Users {
	return &Users{Data: make([]User, 0)}
}

type Database interface {
	GetUserByID(id int) (User, error)
	GetUserByUsername(username string) (User, error)
	AddUserToDB(user User) error
	NumberOfUsers() int
}
