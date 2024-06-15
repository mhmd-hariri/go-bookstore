package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mhmd-hariri/go-bookstore/pkg/config"
)

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {

	config.DB.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	config.DB.NewRecord(b)
	config.DB.Create(&b)
	return b
}

func GetAllbooks() []Book {
	var books []Book
	config.DB.Find(&books)
	return books
}
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	config.DB.Where("ID=?", Id).Find(&book)
	return &book, config.DB
}
func DeleteBook(Id int64) Book {
	var book Book
	config.DB.Where("ID=?", Id).Delete(&book)
	return book
}
