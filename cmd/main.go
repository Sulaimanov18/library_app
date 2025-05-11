package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/Sulaimanov18/library_app/internal/authors"
	"github.com/Sulaimanov18/library_app/internal/books"
	"github.com/Sulaimanov18/library_app/internal/common"
	"github.com/Sulaimanov18/library_app/internal/readers"
	"github.com/Sulaimanov18/library_app/pkg/db"
)

func main() {
	_ = godotenv.Load()

	// Создаем строку подключения (используется в goose)
	connStr := db.BuildConnString()

	// Открываем соединение с использованием sql.DB
	sqlDB, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Cannot open sql.DB: %v", err)
	}
	defer sqlDB.Close()

	// Выполняем миграции
	if err := runMigrations(sqlDB); err != nil {
		log.Fatalf("Migration error: %v", err)
	}
	log.Println("✅ Migrations completed")

	// Подключаемся через pgx.Conn для передачи в сервисы
	pgxConn, err := db.Connect()
	if err != nil {
		log.Fatalf("Cannot connect via pgx: %v", err)
	}
	defer pgxConn.Close(context.Background())

	// Запускаем Gin-сервер
	r := gin.Default()
	common.RegisterTestRoutes(r)
	books.RegisterBookRoutes(r, books.NewBookService(pgxConn))
	authors.RegisterAuthorRoutes(r, authors.NewAuthorService(pgxConn))
	readers.RegisterReaderRoutes(r, readers.NewReaderService(pgxConn))

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func runMigrations(db *sql.DB) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	return goose.Up(db, "./migrations")
}
