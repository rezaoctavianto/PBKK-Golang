package main

import (
	"Authors/config"
	"Authors/controllers/authorcontroller"
	"Authors/controllers/bookcontroller"
	"Authors/controllers/collectioncontroller"
	"Authors/controllers/homecontroller"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	router := gin.Default()
	//1.Home
	router.GET("/", homecontroller.Welcome)

	//2.Books
	bookRoutes := router.Group("/books")
	{
		bookRoutes.GET("", bookcontroller.Index)         // List books
		bookRoutes.GET("/add", bookcontroller.Add)       // Add form
		bookRoutes.POST("/add", bookcontroller.Add)      // Add submission
		bookRoutes.GET("/edit", bookcontroller.Edit)     // Edit form
		bookRoutes.POST("/edit", bookcontroller.Edit)    // Edit submission
		bookRoutes.GET("/detail", bookcontroller.Detail) // View details
		bookRoutes.GET("/delete", bookcontroller.Delete) // Delete book
	}

	//3.Authors
	authorRoutes := router.Group("/authors")
	{
		authorRoutes.GET("", authorcontroller.Index)      // List authors
		authorRoutes.GET("/add", authorcontroller.Add)    // Show add form
		authorRoutes.POST("/add", authorcontroller.Add)   // Handle add form submission
		authorRoutes.GET("/edit", authorcontroller.Edit)  // Show edit form
		authorRoutes.POST("/edit", authorcontroller.Edit) // Handle edit form submission
		authorRoutes.GET("/delete", authorcontroller.Delete)
	}

	//4.Collection
	router.GET("/collections", collectioncontroller.Index)

	log.Println("Server running on port :8000")
	router.Run(":8000")
}
