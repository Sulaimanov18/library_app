package main

import (
	"log"

	"context"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Sulaimanov18/library_app/internal/authors"
	"github.com/Sulaimanov18/library_app/internal/books"
	"github.com/Sulaimanov18/library_app/internal/common"
	"github.com/Sulaimanov18/library_app/pkg/db"
)

func main() {
	_ = godotenv.Load() // загружаем .env

	conn, err := db.Connect() // подключение к базе
	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}
	defer func() {
		if err := conn.Close(context.Background()); err != nil {
			log.Printf("Error closing DB connection: %v", err)
		}
	}()

	r := gin.Default()

	common.RegisterTestRoutes(r)
	books.RegisterBookRoutes(r, books.NewBookService(conn))             // временно
	authors.RegisterAuthorRoutes(r, authors.NewAuthorService(conn)) // <-- ВАЖНО

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
