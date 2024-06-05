package db

import (
	"database/sql"
	"fmt"
	"log"
	"movie_crud/internal/config"

	_ "github.com/lib/pq"
)

func Connect(dbConf config.DB) (*sql.DB, error) {

	connstr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", dbConf.Driver, dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database)
	log.Println("Connecting db to :", connstr)
	conn, err := sql.Open(dbConf.Driver, connstr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
