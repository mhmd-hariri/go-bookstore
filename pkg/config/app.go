package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:root@/book_store?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
