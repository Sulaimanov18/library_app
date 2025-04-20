package authors

import (
	"time"
)
type Author struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Bio string `json:"bio"`
	CreateAt time.Time `json:"created_at"`
	
}

type CreateAuthorRequest struct {
	Name string `json:"name" binding:"required"`
	Bio string `json:"bio"`
}

type UpdateAuthorRequest struct{
	Name string `json:"name"`
	Bio string `json:"bio"`
}

