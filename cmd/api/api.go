package main

import "net/http"

type application struct {
	config config
}

type config struct {
	addr string
	db   DBConfig
}

type DBConfig struct {
	dsn string
}

func (app *application) mount() http.Handler {
	
}