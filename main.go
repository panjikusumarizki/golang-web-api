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

	// CRUD

	// ===========
	// CREATE data
	// ===========

	// book := book.Book{}
	// book.Title = "Dekat"
	// book.Price = 120000
	// book.Discount = 20
	// book.Rating = 5
	// book.Description = "Buku tentang arti dekat"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error creating book record")
	// 	fmt.Println("==========================")
	// }

	// ===========
	// GET data
	// ===========
	// var books []book.Book

	// err = db.Debug().Find(&books).Error // Get all
	// err = db.Debug().Where("rating = ?", 5).Find(&books).Error

	// if err != nil {
	// 	fmt.Println("=========================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("=========================")
	// }

	// for _, b := range books {
	// 	fmt.Println("Title:", b.Title)
	// 	fmt.Println("book object %v", b)
	// }

	var book book.Book

	err = db.Debug().Where("id = ?", 1).First(&book).Error
	if err != nil {
		fmt.Println("=========================")
		fmt.Println("Error finding book record")
		fmt.Println("=========================")
	}

	// ===========
	// UPDATE data
	// ===========
	// book.Title = "Jauh 2"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("=========================")
	// 	fmt.Println("Error updating book record")
	// 	fmt.Println("=========================")
	// }

	err = db.Delete(&book).Error
	if err != nil {
		fmt.Println("=========================")
		fmt.Println("Error deleting book record")
		fmt.Println("=========================")
	}

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/book/:id/:title", handler.BookHandler)
	v1.GET("/book", handler.BookQueryHandler)

	v1.POST("/book", handler.PostBookHandler)

	router.Run(":8000")
}
