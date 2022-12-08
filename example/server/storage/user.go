package storage

import (
	"crypto/rsa"

	"golang.org/x/text/language"
)

type User struct {
	ID                string
	Username          string
	Password          string
	FirstName         string
	LastName          string
	Email             string
	EmailVerified     bool
	Phone             string
	PhoneVerified     bool
	PreferredLanguage language.Tag
}

type Service struct {
	keys map[string]*rsa.PublicKey
}

type UserStore interface {
	GetUserByID(string) *User
	GetUserByUsername(string) *User
	ExampleClientID() string
}

type userStore struct {
	users map[string]*User
}

func NewUserStore() UserStore {
	return userStore{
		users: map[string]*User{
			"id1": {
				ID:                "id1",
				Username:          "matthew.cuthbert@demo.localhost",
				Password:          "password",
				FirstName:         "Matthew",
				LastName:          "Cuthbert",
				Email:             "matthew.cuthbert@demo.localhost",
				EmailVerified:     true,
				Phone:             "",
				PhoneVerified:     false,
				PreferredLanguage: language.English,
			},
			"id2": {
				ID:                "id2",
				Username:          "anne.shirley@demo.localhost",
				Password:          "password",
				FirstName:         "Anne",
				LastName:          "Shirley",
				Email:             "anne.shirley@demo.localhost",
				EmailVerified:     true,
				Phone:             "",
				PhoneVerified:     false,
				PreferredLanguage: language.English,
			},
			"id3": {
				ID:                "id3",
				Username:          "diana.barry@demo.localhost",
				Password:          "password",
				FirstName:         "Diana",
				LastName:          "Barry",
				Email:             "diana.barry@demo.localhost",
				EmailVerified:     true,
				Phone:             "",
				PhoneVerified:     false,
				PreferredLanguage: language.English,
			},
			"id4": {
				ID:                "id4",
				Username:          "gilbert.blythe@demo.localhost",
				Password:          "password",
				FirstName:         "Gilbert",
				LastName:          "Blythe",
				Email:             "gilbert.blythe@demo.localhost",
				EmailVerified:     true,
				Phone:             "",
				PhoneVerified:     false,
				PreferredLanguage: language.English,
			},
			"id5": {
				ID:                "id5",
				Username:          "marilla.cuthbert@demo.localhost",
				Password:          "password",
				FirstName:         "Marilla",
				LastName:          "Cuthbert",
				Email:             "marilla.cuthbert@demo.localhost",
				EmailVerified:     true,
				Phone:             "",
				PhoneVerified:     false,
				PreferredLanguage: language.English,
			},
		},
	}
}

// ExampleClientID is only used in the example server
func (u userStore) ExampleClientID() string {
	return "cargo-mailbox"
}

func (u userStore) GetUserByID(id string) *User {
	return u.users[id]
}

func (u userStore) GetUserByUsername(username string) *User {
	for _, user := range u.users {
		if user.Username == username {
			return user
		}
	}
	return nil
}
