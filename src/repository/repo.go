package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=8080 user=gorm dbname=obas password=password")
	if err != nil {
		fmt.Print(" There was an error ", err)
	}
	return db
}
