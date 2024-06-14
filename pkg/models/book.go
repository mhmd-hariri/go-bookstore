package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mhmd-hariri/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
func (b *Book) CreateBook() *Book {

	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllbooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db.Where("ID=?", Id).Find(&book)
	return &book, db
}
func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}
