package main

import (
	"salotto/service"
)

func main() {
	service.ConnectDB()
	router := MapRoutes()
	router.Run(":8089")
}
