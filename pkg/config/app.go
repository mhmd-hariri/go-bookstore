package config

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB           *gorm.DB
	Port         int
	User         string
	Password     string
	DatabaseName string
	JwtKey       []byte
)

func init() {
	// initilize global variable
	Port = 9010
	User = "root"
	Password = "root"
	DatabaseName = "book_store"
	var err error
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", User, Password, DatabaseName))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection established")

	b := make([]byte, 32)
	_, err = rand.Read(b)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	secretKey := base64.URLEncoding.EncodeToString(b)
	JwtKey = []byte(secretKey)
}
