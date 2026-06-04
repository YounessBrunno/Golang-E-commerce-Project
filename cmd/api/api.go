package main

type application struct {
  config config
}

type config struct {
  addr string
  db DBConfig
}

type DBConfig struct {
  dsn string
  
}