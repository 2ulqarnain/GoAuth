package main

import (
	"GoAuth/internal/auth"
	"GoAuth/internal/config"
	"GoAuth/internal/db"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	cfg := config.NewConfig()
	pool := db.NewPool(ctx, cfg.DatabaseURL)
	defer pool.Close()
	r := auth.NewRouter()

	fmt.Printf("Starting server on port %s...\n", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), r))
}
