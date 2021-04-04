package tests

import (
	"fmt"
	"library-system/db"
	"library-system/repositories"
	"library-system/services"
	"testing"
)

func TestGetAll(t *testing.T) {
	bookRepository := repositories.NewBookRepository(db.GetConnection())
	bookService := services.NewBookService(bookRepository)

	books, err := bookService.GetAll()
	if err != nil {
		panic(err)
	}

	for _, book := range books {
		fmt.Println(book)
	}
}

func TestGetBookByID(t *testing.T) {
	bookRepository := repositories.NewBookRepository(db.GetConnection())
	bookService := services.NewBookService(bookRepository)

	book, err := bookService.GetBookByID(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(book)
}

func TestCreateBook(t *testing.T) {
	bookRepository := repositories.NewBookRepository(db.GetConnection())
	bookService := services.NewBookService(bookRepository)

	book, err := bookService.CreateBook(3, "Test Book Service 3", false)
	if err != nil {
		panic(err)
	}

	fmt.Println(book)
}

func TestFlagBookRentByIDAndIsRent(t *testing.T) {
	bookRepository := repositories.NewBookRepository(db.GetConnection())
	bookService := services.NewBookService(bookRepository)

	book, err := bookService.FlagBookRentByIDAndIsRent(3, true)
	if err != nil {
		panic(err)
	}

	fmt.Println(book)
}

func TestGetAllBookRented(t *testing.T) {
	bookRepository := repositories.NewBookRepository(db.GetConnection())
	bookService := services.NewBookService(bookRepository)

	books, err := bookService.GetAllBooksRented()
	if err != nil {
		panic(err)
	}

	for _, book := range books {
		fmt.Println(book)
	}
}
