package main

import (
	"crudexample/routers"
	"crudexample/utils"
)

func main() {
	router := routers.InitRoute()
	port := utils.GetEnv("SERVER_PORT", ":8080")
	router.Run(port)
}
