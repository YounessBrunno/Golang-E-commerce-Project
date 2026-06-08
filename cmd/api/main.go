package main


import ( "log"
         "os" )



func main() {

	cfg := config{
		addr: ":8000",
		db: DBConfig{},
	}

	api := application{
		config: cfg,
	}
    
	if err := api.serve(api.mount()); err != nil {
		
		log.Printf("server has failed to start, error: %v", err)

		os.Exit(1)
	}

}