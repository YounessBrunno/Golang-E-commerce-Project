package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"github.com/YounessBrunno/Golang-E-commerce-Project/internals/env"
    "log/slog"
	_ "github.com/jackc/pgx/v5/stdlib"
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

	// Database
	connection, err := sql.Open("pgx", cfg.db.dsn)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	
    connection.PingContext(ctx)

	defer connection.Close()

	logger.Info("Database connection established successfully","dsn", cfg.db.dsn)

    //application
	api := application{
		config: cfg,
	}
    
	if err := api.serve(api.mount()); err != nil {
		
		logger.Error("server has failed to start", "error", err)

		os.Exit(1)
	}

}