package main

import (
	"log"
	"movie_crud/internal/adapter/persistance/db"
	"movie_crud/internal/config"
)

func main() {

	conf, err := config.GetConfig("./.env")
	if err != nil {
		log.Println("error while loading config", err)
		return
	}

	conn, err := db.Connect(*conf.DB)
	if err != nil {
		log.Println("error while connecting to db", err)
		return
	}

	defer conn.Close()

	_, err = conn.Exec("CREATE TABLE FRUITS(name VARCHAR(10));")
	if err != nil {
		log.Println("error while creating table", err)
		return
	}
}
