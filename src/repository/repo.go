package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=26257 user=root dbname=bank sslmode=disable")
	if err != nil {
		fmt.Print(" There was an error ", err)
	}
	return db
}
