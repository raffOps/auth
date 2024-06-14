package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func sanityCheck() {
	err := godotenv.Load("config/dev.env")
	if err != nil {
		panic(err)
	}
	if _, ok := os.LookupEnv("POSTGRES_HOST"); !ok {
		panic("POSTGRES_HOST environment variable is not set")
	}

	if _, ok := os.LookupEnv("POSTGRES_PORT"); !ok {
		panic("POSTGRES_HOST environment variable is not set")
	}

	if _, ok := os.LookupEnv("POSTGRES_USER"); !ok {
		panic("POSTGRES_USER environment variable is not set")
	}

	if _, ok := os.LookupEnv("POSTGRES_PASSWORD"); !ok {
		panic("POSTGRES_PASSWORD environment variable is not set")
	}

	if _, ok := os.LookupEnv("POSTGRES_DB"); !ok {
		panic("POSTGRES_DB environment variable is not set")
	}
}

func GetPostgresConn(setLimits bool) *sql.DB {
	sanityCheck()
	err := godotenv.Load("config/dev.env")
	if err != nil {
		panic(err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if setLimits {
		db.SetMaxOpenConns(5)
		db.SetMaxIdleConns(5)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	return db
}
