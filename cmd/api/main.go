package main



func main() {

	cfg := config{
		addr: ":8000",
		db: DBConfig{},
	}

	api := application{
		config: cfg,
	}
    
	api.serve(api.mount())
	
}

