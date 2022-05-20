package main

import (
	"awesomeProject/constants"
	"awesomeProject/controller"
)

func main() {
	server := controller.Server{}
	server.Initialize(constants.DB_USER, constants.DB_PASSWORD, constants.DB_HOST, constants.DB_PORT, constants.DB_NAME)
	server.Run(":8080")
}
