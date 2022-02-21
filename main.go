package main

import (
	"golang-web-api/book"
	"golang-web-api/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/golearn?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection failed")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)

	// Find All
	// books, err := bookRepository.FindAll()

	// for _, book := range books {
	// 	fmt.Println("Title:", book.Title)
	// }

	// Find By ID
	// book, err := bookRepository.FindByID(2)
	// fmt.Println("Title:", book.Title)

	// Create
	// book := book.Book{
	// 	Title:       "Koplok",
	// 	Description: "Goblok looo!!!!!!",
	// 	Price:       200000,
	// 	Rating:      5,
	// 	Discount:    0,
	// }

	// bookRepository.Create(book)
	var book book.Book
	book, err = bookRepository.FindByID(3)

	// Update
	// book.Title = "Blokkkkk"
	// book, err = bookRepository.Update(book)

	book, err = bookRepository.Delete(book)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/book/:id/:title", handler.BookHandler)
	v1.GET("/book", handler.BookQueryHandler)

	v1.POST("/book", handler.PostBookHandler)

	router.Run(":8000")
}
