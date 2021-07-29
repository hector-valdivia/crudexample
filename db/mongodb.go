package db

import (
	"crudexample/utils"
	"github.com/goonode/mogo"
	"log"
)

var mongoConnection *mogo.Connection = nil

func GetConnection() *mogo.Connection {
	if mongoConnection == nil {
		connectionString := utils.GetEnv("DB_CONNECTION_STRING")
		dbName := utils.GetEnv("DB_NAME")
		config := &mogo.Config{
			ConnectionString: connectionString,
			Database:         dbName,
		}
		mongoConnection, err := mogo.Connect(config)
		if err != nil {
			log.Fatal(err)
		} else {
			return mongoConnection
		}
	}

	return mongoConnection
}
