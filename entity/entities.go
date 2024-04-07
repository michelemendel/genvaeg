package entity

import (
	"github.com/michelemendel/genvaeg/util"
)

type User struct {
	UUID           string
	Name           string
	HashedPassword string
}

func NewUser(name, hashePassword string) User {
	return User{
		UUID:           util.GenerateUUID(),
		Name:           name,
		HashedPassword: hashePassword,
	}
}

type URLPair struct {
	FullURL  string
	ShortURL string
}

func NewURLPair(fullURL, shortURL string) *URLPair {
	return &URLPair{
		FullURL:  fullURL,
		ShortURL: shortURL,
	}
}
