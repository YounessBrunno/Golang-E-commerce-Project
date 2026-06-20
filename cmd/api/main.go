package main


import ( "log"
         "os"
		 "database/sql"
		  _ "github.com/jackc/pgx/v5/stdlib"
	         )


func main() {
    
	cfg := config{
		addr: ":8000",
		db: DBConfig{
			dsn: "user=postgres password=password dbname=ecommerce-db host=localhost port=5432 sslmode=disable",
		},
	}

	// Database
	connection, err := sql.Open("pgx", cfg.db.dsn)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
  
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