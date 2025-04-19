package books

import (
	"errors"
)

type BookService struct {
	books  []Book
	lastID int
}

func NewBookService() *BookService {
	return &BookService{
		books:  []Book{},
		lastID: 0,
	}
}

func (s *BookService) GetAll() []Book {
	return s.books
}

func (s *BookService) Create(req CreateBookRequest) (Book, error) {
	if req.Title == "" {
		return Book{}, errors.New("title is required")
	}
	book := Book{
		ID:          s.lastID + 1,
		Title:       req.Title,
		Description: req.Description,
		Publisher:   req.Publisher,
		Count:       req.Count,
	}
	s.lastID++
	s.books = append(s.books, book)
	return book, nil
}