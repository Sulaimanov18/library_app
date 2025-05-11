package readers

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
)

type ReaderService struct {
	db *pgx.Conn
}

func NewReaderService(conn *pgx.Conn) *ReaderService {
	return &ReaderService{db: conn}
}

func (s *ReaderService) GetAll() ([]Reader, error) {
	rows, err := s.db.Query(context.Background(), "SELECT id, name, email, created_at FROM readers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var readers []Reader
	for rows.Next() {
		var r Reader
		if err := rows.Scan(&r.ID, &r.Name, &r.Email, &r.CreatedAt); err != nil {
			return nil, err
		}
		readers = append(readers, r)
	}
	return readers, nil
}

func (s *ReaderService) GetByID(id int) (*Reader, error) {
	var r Reader
	err := s.db.QueryRow(context.Background(), "SELECT id, name, email, created_at FROM readers WHERE id=$1", id).
		Scan(&r.ID, &r.Name, &r.Email, &r.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &r, nil
}

func (s *ReaderService) Create(req CreateReaderRequest) (*Reader, error) {
	var r Reader
	err := s.db.QueryRow(
		context.Background(),
		"INSERT INTO readers (name, email, created_at) VALUES ($1, $2, $3) RETURNING id, name, email, created_at",
		req.Name, req.Email, time.Now(),
	).Scan(&r.ID, &r.Name, &r.Email, &r.CreatedAt)

	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *ReaderService) Update(id int, req UpdateReaderRequest) error {
	_, err := s.db.Exec(
		context.Background(),
		"UPDATE readers SET name=$1, email=$2 WHERE id=$3",
		req.Name, req.Email, id,
	)
	return err
}

func (s *ReaderService) Delete(id int) error {
	_, err := s.db.Exec(context.Background(), "DELETE FROM readers WHERE id=$1", id)
	return err
}
