package tests

import (
	"context"
	"fmt"
	"library-system/db"
	"library-system/entities"
	"library-system/repositories"
	"testing"
)

func TestFindAll(t *testing.T) {
	bookRepository := repositories.NewBookRepository(db.GetConnection())

	books, err := bookRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, book := range books {
		fmt.Println(book)
	}
}

func TestFindByID(t *testing.T) {
	bookRepository := repositories.NewBookRepository(db.GetConnection())

	book, err := bookRepository.FindByID(context.Background(), 4)
	if err != nil {
		panic(err)
	}

	fmt.Println(book)
}

func TestBookInsert(t *testing.T) {
	bookRepository := repositories.NewBookRepository(db.GetConnection())

	book := entities.Book{
		ID:       2,
		Name:     "Test Book 1",
		IsRented: false,
	}

	result, err := bookRepository.Insert(context.Background(), book)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestBookUpdateIsRent(t *testing.T) {
	bookRepository := repositories.NewBookRepository(db.GetConnection())

	book := entities.Book{
		ID:       2,
		IsRented: true,
	}

	result, err := bookRepository.UpdateIsRent(context.Background(), book)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindAllByIsRent(t *testing.T) {
	bookRepository := repositories.NewBookRepository(db.GetConnection())

	books, err := bookRepository.FindAllByIsRent(context.Background(), false)
	if err != nil {
		panic(err)
	}

	for _, book := range books {
		fmt.Println(book)
	}
}
