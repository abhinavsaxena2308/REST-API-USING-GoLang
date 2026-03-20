package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/abhinavsaxena2308/REST-API-USING-GoLang/cmd/api"
	"github.com/abhinavsaxena2308/REST-API-USING-GoLang/config"
	"github.com/abhinavsaxena2308/REST-API-USING-GoLang/db"
)

func main() {
	//initializr db
	dbConn, err := db.NewPostgreSql(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Envs.DBHost,
		config.Envs.DBPort,
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBName, 
	))

	if err != nil {
		log.Fatal("error in opening the database: ", err)
	}

	if err := initDB(dbConn); err != nil {
		log.Fatal("error in connecting to the database: ", err)
	}
	// start api server
	apiServer := api.NewAPIServer(":8080")
	if err := apiServer.Run(); err!= nil{
		log.Fatal("error in running the server")
	}
}

func initDB(db *sql.DB) error{
	err := db.Ping()
	if err != nil {
		return err
	}

	log.Println("Database connection established",config.Envs.DBName)
	return nil
}