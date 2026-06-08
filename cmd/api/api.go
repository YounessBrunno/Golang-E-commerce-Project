package main

import "net/http"

type application struct {
	config config
}

func (app *application) mount() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Health check passed!"))
	})
    
	return mux

}

func (app *application) serve(h http.Handler) error {

	server := http.Server{
		Addr:    app.config.addr,
		Handler: h,
	}
	
	return server.ListenAndServe()
}

type config struct {
	addr string
	db   DBConfig
}

type DBConfig struct {
	dsn string
}

