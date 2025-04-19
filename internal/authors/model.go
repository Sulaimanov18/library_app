package authors


type Authors struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Bio string `json:"bio"`
	Created string `json:"date"`
	
}

type CreateAuthors struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Publisher   string `json:"publisher"`
	Count       int    `json:"count"`
}