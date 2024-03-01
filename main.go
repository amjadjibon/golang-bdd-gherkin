package main

import (
	"github.com/amjadjibon/golang-bdd-gherkin/dbx"
	_ "github.com/amjadjibon/golang-bdd-gherkin/docs"
	"github.com/amjadjibon/golang-bdd-gherkin/handlers"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @Books Swagger Example API
// @version 1.0
// @description This is a sample server Books server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.books.io/support
// @contact.email support@books.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host http://localhost:8080
// @BasePath /v1
func main() {
	dbx.GetDB()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	booksRouter := e.Group("/v1/books")

	// Routes
	booksRouter.GET("", handlers.GetBooks)
	booksRouter.GET("/:id", handlers.GetBook)
	booksRouter.POST("", handlers.CreateBook)
	booksRouter.PUT("/:id", handlers.UpdateBook)
	booksRouter.DELETE("/:id", handlers.DeleteBook)

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
