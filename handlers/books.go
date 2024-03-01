package handlers

import (
	"strconv"

	"github.com/amjadjibon/golang-bdd-gherkin/dbx"
	"github.com/amjadjibon/golang-bdd-gherkin/models"
	"github.com/labstack/echo/v4"
)

// GetBooks godoc
// @Summary Get all books
// @Description Get all books
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {object} []models.Book
// @Router /books [get]
func GetBooks(c echo.Context) error {
	books := []models.Book{}
	dbx.DB.Find(&books)

	return c.JSON(200, books)
}

// GetBook godoc
// @Summary Get a book
// @Description Get a book
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Router /books/{id} [get]
func GetBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	book := models.Book{}
	dbx.DB.First(&book, id)

	return c.JSON(200, book)
}

// CreateBook godoc
// @Summary Create a book
// @Description Create a book
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.BookBase true "Book"
// @Success 201 {object} models.Book
// @Router /books [post]
func CreateBook(c echo.Context) error {
	bookBase := models.BookBase{}
	if err := c.Bind(&bookBase); err != nil {
		return err
	}

	if !bookBase.Validate() {
		return c.JSON(400, "invalid book")
	}

	book := models.Book{
		Title:  bookBase.Title,
		Author: bookBase.Author,
	}

	dbx.DB.Create(&book)
	return c.JSON(201, book)
}

// UpdateBook godoc
// @Summary Update a book
// @Description Update a book
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.BookBase true "Book"
// @Success 200 {object} models.Book
// @Router /books/{id} [put]
func UpdateBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	bookBase := models.BookBase{}
	if err := c.Bind(&bookBase); err != nil {
		return err
	}

	if !bookBase.Validate() {
		return c.JSON(400, "invalid book")
	}

	book := models.Book{}
	update := false
	dbx.DB.First(&book, id)

	if book.ID == 0 {
		return c.JSON(404, "book not found")
	}

	if bookBase.Title != "" && bookBase.Title != book.Title {
		book.Title = bookBase.Title
		update = true
	}

	if bookBase.Author != "" && bookBase.Author != book.Author {
		book.Author = bookBase.Author
		update = true
	}

	if update {
		dbx.DB.Save(&book)
	}

	return c.JSON(200, book)
}

// DeleteBook godoc
// @Summary Delete a book
// @Description Delete a book
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200
// @Router /books/{id} [delete]å
func DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	book := models.Book{}
	dbx.DB.First(&book, id)

	if book.ID == 0 {
		return c.JSON(404, "book not found")
	}

	dbx.DB.Delete(&book)

	return c.JSON(200, "book deleted")
}
