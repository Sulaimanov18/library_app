package books

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
)

type BookService struct {
	conn *pgx.Conn
}

func NewBookService(conn *pgx.Conn) *BookService {
	return &BookService{conn: conn}
}

func (s *BookService) GetAll() ([]Book, error) {
	rows, err := s.conn.Query(context.Background(), "SELECT id, title, description, publisher, count FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Description, &b.Publisher, &b.Count); err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	return books, nil
}

func (s *BookService) GetByID(id int) (*Book, error) {
	var b Book
	err := s.conn.QueryRow(context.Background(),
		"SELECT id, title, description, publisher, count FROM books WHERE id=$1", id).
		Scan(&b.ID, &b.Title, &b.Description, &b.Publisher, &b.Count)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}

	return &b, nil
}

func (s *BookService) Create(req CreateBookRequest) (Book, error) {
	var id int
	err := s.conn.QueryRow(context.Background(),
		`INSERT INTO books (title, description, publisher, count, created_at)
         VALUES ($1, $2, $3, $4, $5)
         RETURNING id`,
		req.Title, req.Description, req.Publisher, req.Count, time.Now()).Scan(&id)

	if err != nil {
		return Book{}, err
	}

	return Book{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Publisher:   req.Publisher,
		Count:       req.Count,
	}, nil
}
