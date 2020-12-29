package controllers

import (
	domain2 "obas/domain/users"
	"obas/io/users"
)

func GetUser(userId string) domain2.User {
	var myuser domain2.User
	user, err := users.GetUser(userId)
	if err != nil {
		return myuser
	}
	return user
}
