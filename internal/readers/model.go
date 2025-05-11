package readers


type reader struct {
	ID	int		`json:"id"`
	Name string `json:"title"`
	Email string `json:""`
	JoinedAt int `json:"date"`
}

type CreateReaderRequest struct {
	Name string `json:"title"`
	Email string `json:""`
	JoinedAt int `json:"date"`
}


