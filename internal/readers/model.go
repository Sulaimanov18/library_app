package readers

import "time"

type Reader struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateReaderRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type UpdateReaderRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
