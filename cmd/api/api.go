package main

import (
	"log"
	"net/http"
	"time"
	"github.com/YounessBrunno/Golang-E-commerce-Project/internals/middlewares"
	"github.com/YounessBrunno/Golang-E-commerce-Project/internals/products"
	repo "github.com/YounessBrunno/Golang-E-commerce-Project/internals/adapters/postgresql/sqlc"
)


type application struct {
	config config
	db     repo.DBTX
}

func (app *application) mount() http.Handler {

	mux := http.NewServeMux()
	
	
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Health check passed!"))
	})

	ProductHandler := products.NewHandler(products.NewProductService(repo.New(app.db)))

	mux.HandleFunc("/products", ProductHandler.ListProducts)

	return mux

}

func (app *application) serve(h http.Handler) error {
    
	server := http.Server{
		Addr:         app.config.addr,
		Handler:      middleware.LoggingMiddleware(h),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	log.Printf("Starting server on %s", app.config.addr)

	return server.ListenAndServe()
}

type config struct {
	addr string
	db   DBConfig
}

type DBConfig struct {
	dsn string
}
