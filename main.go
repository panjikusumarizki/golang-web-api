package main

import (
	"fmt"
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
	bookService := book.NewService(bookRepository)

	// Find all
	books, err := bookService.FindAll()

	for _, book := range books {
		fmt.Println("Title:", book.Title)
	}

	// Find by ID
	// book, err := bookService.FindByID(2)
	// fmt.Println("Title:", book.Title)

	// Create
	// bookRequest := book.BookRequest{
	// 	Title: "Hmmm",
	// 	Price: "100000",
	// }

	// bookService.Create(bookRequest)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/book/:id/:title", handler.BookHandler)
	v1.GET("/book", handler.BookQueryHandler)

	v1.POST("/book", handler.PostBookHandler)

	router.Run(":8000")
}
