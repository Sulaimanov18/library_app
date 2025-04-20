package authors

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
)

type AuthorService struct {
	db *pgx.Conn
}

func NewAuthorService(db *pgx.Conn) *AuthorService {
	return &AuthorService{db: db}
}

func (s *AuthorService) GetAll() []Author {
	rows, err := s.db.Query(context.Background(), "SELECT id, name, bio, created_at FROM authors")
	if err != nil {
		return []Author{}
	}
	defer rows.Close()

	var authors []Author
	for rows.Next() {
		var a Author
		err := rows.Scan(&a.ID, &a.Name, &a.Bio, &a.CreateAt)
		if err == nil {
			authors = append(authors, a)
		}
	}
	return authors
}

func (s *AuthorService) GetByID(id int) (*Author, error) {
	var a Author
	err := s.db.QueryRow(context.Background(),
		"SELECT id, name, bio, created_at FROM authors WHERE id=$1", id).
		Scan(&a.ID, &a.Name, &a.Bio, &a.CreateAt)

	if err != nil {
		return nil, errors.New("author not found")
	}
	return &a, nil
}

func (s *AuthorService) Create(req CreateAuthorRequest) Author {
	var id int
	err := s.db.QueryRow(context.Background(),
		"INSERT INTO authors (name, bio, created_at) VALUES ($1, $2, $3) RETURNING id",
		req.Name, req.Bio, time.Now(),
	).Scan(&id)

	if err != nil {
		return Author{}
	}

	return Author{
		ID:       id,
		Name:     req.Name,
		Bio:      req.Bio,
		CreateAt: time.Now(),
	}
}
