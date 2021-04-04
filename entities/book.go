package entities

import "time"

type Book struct {
	ID        int
	Name      string
	IsRented  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
