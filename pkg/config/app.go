package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

var (
	db           *gorm.DB
	Port         int
	User         string
	Password     string
	DatabaseName string
)

func init() {
	// initilize global variable
	Port = 9010
	User = "root"
	Password = "root"
	DatabaseName = "book_store"
}

func Connect() {
	d, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", User, Password, DatabaseName))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
