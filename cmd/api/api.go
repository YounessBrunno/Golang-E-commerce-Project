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
    
	return mux

}

func (app *application) serve() http.Handler {

	server := http.Server{
		Addr:    app.config.addr,
		Handler: app.mount(),
	}

	return server.Handler
}
type config struct {
	addr string
	db   DBConfig
}

type DBConfig struct {
	dsn string
}

