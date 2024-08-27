package common

import (
	"log"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

type environment struct {
	Port             int    `env:"PORT,default=8080"`
	PostgresHost     string `env:"POSTGRES_HOST,default=localhost"`
	PostgresPort     string `env:"POSTGRES_PORT,default=443"`
	PostgresDB       string `env:"POSTGRES_DB,default=postgres"`
	PostgresUser     string `env:"POSTGRES_USER,default=postgres"`
	PostgresPassword string `env:"POSTGRES_PASSWORD,default=postgres"`
	Extras           env.EnvSet
}

var Env environment

func SetupEnvironmentVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	es, error := env.UnmarshalFromEnviron(&Env)
	if error != nil {
		log.Fatal(error)
	}
	Env.Extras = es
}