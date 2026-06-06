package main

import "net/http"

func main() {

	cfg := config{
		addr: ":8000",
		db: DBConfig{},
	}

	api := application{
		config: cfg,
	}
    
	http.ListenAndServe(cfg.addr, api.mount())

}

