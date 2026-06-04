package main

import "net/http"

type application struct {
	config config
}

func (app *application) mount() http.Handler {
	
}

type config struct {
	addr string
	db   DBConfig
}

type DBConfig struct {
	dsn string
}