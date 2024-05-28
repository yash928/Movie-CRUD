package config

import "github.com/joho/godotenv"

type Config struct {
	DB     *DB
	Server *Server
}

type DB struct {
	Driver   string
	Database string
	Username string
	Password string
	Host     string
	Port     string
}

type Server struct {
	Env  string
	Port string
}

func GetConfig(envPath string) (*Config, error) {

	envMap, err := godotenv.Read(envPath)
	if err != nil {
		return nil, err
	}

	return &Config{
		DB: &DB{
			Driver:   envMap["DB_DRIVER"],
			Database: envMap["DB_DATABASE"],
			Username: envMap["DB_USERNAME"],
			Password: envMap["DB_PASSWORD"],
			Host:     envMap["DB_HOST"],
			Port:     envMap["DB_PORT"],
		},
		Server: &Server{
			Env:  envMap["ENVIRONMENT"],
			Port: envMap["SERVER_PORT"],
		},
	}, nil
}

type CtxKey struct {
	RequestID string `json:"request_id"`
	Session   string `json:"session"`
}
