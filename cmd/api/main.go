package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"github.com/YounessBrunno/Golang-E-commerce-Project/internals/env"
    "github.com/YounessBrunno/Golang-E-commerce-Project/internals/middlewares"
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

	// logger middleware

	// Database
	connection, err := sql.Open("pgx", cfg.db.dsn)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	
    connection.PingContext(ctx)

	defer connection.Close()

    //application
	api := application{
		config: cfg,
	}
    
	if err := api.serve(api.mount()); err != nil {
		
		log.Printf("server has failed to start, error: %v", err)

		os.Exit(1)
	}

}