package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/Sulaimanov18/library_app/internal/books"
	"github.com/Sulaimanov18/library_app/internal/common"
)

func main() {
	r := gin.Default()

	common.RegisterTestRoutes(r)
	books.RegisterBookRoutes(r, books.NewBookService())

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}