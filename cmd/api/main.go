package main

import (
	"context"
	"log"
	"os"
	"github.com/YounessBrunno/Golang-E-commerce-Project/internals/env"
	"log/slog"
	"github.com/jackc/pgx/v5/pgxpool"
)


func main() {
	ctx := context.Background()
    
	cfg := config{
		addr: ":8000",
		db: DBConfig{
			dsn: env.GetEnvKey("DB_DSN", "user=postgres password=password dbname=ecommerce-db host=localhost port=5432 sslmode=disable"),
		},
	}

	// logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
    slog.SetDefault(logger)   

	// Database (pgxpool)
	pool, err := pgxpool.New(ctx, cfg.db.dsn)
	if err != nil {
		log.Fatal("failed to create pgx pool:", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("failed to ping database:", err)
	}

	defer pool.Close()

	logger.Info("Database connection established successfully", "dsn", cfg.db.dsn)

	// application
	api := application{
		config: cfg,
		db:     pool,
	}
    
	if err := api.serve(api.mount()); err != nil {
		
		logger.Error("server has failed to start", "error", err)

		os.Exit(1)
	}

}