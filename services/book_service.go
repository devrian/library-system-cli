package services

import (
	"context"
	"errors"
	"library-system/entities"
	"library-system/repositories"
	"strconv"
)

type BookService interface {
	GetAll() ([]entities.Book, error)
	GetBookByID(ID int) (entities.Book, error)
	CreateBook(ID int, name string, isRented bool) (entities.Book, error)
	FlagBookRentByIDAndIsRent(ID int, isRented bool) (entities.Book, error)
	GetAllBooksRented() ([]entities.Book, error)
}

type service struct {
	repository repositories.BookRepository
}

func NewBookService(repository repositories.BookRepository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]entities.Book, error) {
	books, err := s.repository.FindAll(context.Background())
	if err != nil {
		return books, err
	}

	return books, nil
}

func (s *service) GetBookByID(ID int) (entities.Book, error) {
	book, err := s.repository.FindByID(context.Background(), ID)
	if err != nil {
		return book, err
	}
	if book.ID == 0 {
		return book, errors.New("Book not found")
	}

	return book, nil
}

func (s *service) CreateBook(ID int, name string, isRented bool) (entities.Book, error) {
	book := entities.Book{
		ID:       ID,
		Name:     name,
		IsRented: isRented,
	}

	checkBookId, err := s.repository.FindByID(context.Background(), book.ID)
	if err != nil {
		return book, err
	}
	if checkBookId.ID != 0 {
		return book, errors.New("code of book has been registered " + strconv.Itoa(checkBookId.ID))
	}

	bookCreated, err := s.repository.Insert(context.Background(), book)
	if err != nil {
		return bookCreated, err
	}

	return bookCreated, nil
}

func (s *service) FlagBookRentByIDAndIsRent(ID int, isRented bool) (entities.Book, error) {
	book, err := s.repository.FindByID(context.Background(), ID)
	if err != nil {
		return book, err
	}
	if book.ID == 0 {
		return book, errors.New("Book not found")
	}

	book.IsRented = isRented
	bookUpdated, err := s.repository.UpdateIsRent(context.Background(), book)
	if err != nil {
		return bookUpdated, err
	}

	return bookUpdated, nil
}

func (s *service) GetAllBooksRented() ([]entities.Book, error) {
	books, err := s.repository.FindAllByIsRent(context.Background(), true)
	if err != nil {
		return books, err
	}

	return books, nil
}
