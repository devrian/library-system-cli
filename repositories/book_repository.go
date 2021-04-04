package repositories

import (
	"context"
	"database/sql"
	"library-system/entities"
	"time"
)

type BookRepository interface {
	FindAll(ctx context.Context) ([]entities.Book, error)
	FindByID(ctx context.Context, ID int) (entities.Book, error)
	Insert(ctx context.Context, book entities.Book) (entities.Book, error)
	UpdateIsRent(ctx context.Context, book entities.Book) (entities.Book, error)
	FindAllByIsRent(ctx context.Context, statusRented bool) ([]entities.Book, error)
}

type repository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(ctx context.Context) ([]entities.Book, error) {
	script := "SELECT id, name, is_rented, created_at, updated_at FROM books"
	rows, err := r.db.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []entities.Book
	for rows.Next() {
		book := entities.Book{}
		rows.Scan(&book.ID, &book.Name, &book.IsRented, &book.CreatedAt, &book.UpdatedAt)
		books = append(books, book)
	}

	return books, nil
}

func (r *repository) FindByID(ctx context.Context, ID int) (entities.Book, error) {
	script := "SELECT id, name, is_rented, created_at, updated_at FROM books WHERE id = ? LIMIT 1"
	rows, err := r.db.QueryContext(ctx, script, ID)
	book := entities.Book{}
	if err != nil {
		return book, err
	}

	defer rows.Close()

	if rows.Next() {
		rows.Scan(&book.ID, &book.Name, &book.IsRented, &book.CreatedAt, &book.UpdatedAt)
		return book, nil
	}

	return book, nil
}

func (r *repository) Insert(ctx context.Context, book entities.Book) (entities.Book, error) {
	script := "INSERT INTO books(id, name, is_rented) VALUES(?, ?, ?)"
	_, err := r.db.ExecContext(ctx, script, book.ID, book.Name, book.IsRented)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *repository) UpdateIsRent(ctx context.Context, book entities.Book) (entities.Book, error) {
	script := "UPDATE books SET is_rented = ?, updated_at = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, script, book.IsRented, time.Now(), book.ID)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *repository) FindAllByIsRent(ctx context.Context, statusRented bool) ([]entities.Book, error) {
	script := "SELECT id, name, is_rented, created_at, updated_at FROM books WHERE is_rented = ?"
	rows, err := r.db.QueryContext(ctx, script, statusRented)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []entities.Book
	for rows.Next() {
		book := entities.Book{}
		rows.Scan(&book.ID, &book.Name, &book.IsRented, &book.CreatedAt, &book.UpdatedAt)
		books = append(books, book)
	}

	return books, nil
}
