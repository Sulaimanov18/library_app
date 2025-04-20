package authors

import (
	"errors"
	"time"
)

type AuthorService struct {
	authors []Author
	lastID  int
}

func NewAuthorService() *AuthorService {
	return &AuthorService{
		authors: []Author{},
		lastID:  0,
	}
}

func (s *AuthorService) GetAll() []Author {
	return s.authors
}

func (s *AuthorService) GetByID(id int) (*Author, error) {
	for _, a := range s.authors {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("author not found")
}

func (s *AuthorService) Create(req CreateAuthorRequest) Author {
	s.lastID++
	author := Author{
		ID:        s.lastID,
		Name:      req.Name,
		Bio:       req.Bio,
		CreateAt: time.Now(),
	}
	s.authors = append(s.authors, author)
	return author
}
