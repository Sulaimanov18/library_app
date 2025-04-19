package books

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Publisher   string `json:"publisher"`
	Count       int    `json:"count"`
}

type CreateBookRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Publisher   string `json:"publisher"`
	Count       int    `json:"count"`
}
