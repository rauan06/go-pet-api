package models

import (
	"errors"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
}

func (b *Book) Validate() error {
	if b.Title == "" {
		return errors.New("book title cannot be empty")
	}

	if b.Author == "" {
		return errors.New("author name cannot be empty")
	}

	if b.ID != 0 {
		return errors.New("you cannot assign book ID by yourself")
	}

	if b.Year == 0 {
		return errors.New("year cannot be empty")
	}

	if b.Year < 0 {
		return errors.New("year cannot be less than 0")
	}

	return nil
}
