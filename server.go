package main

import (
	"github.com/oat431/libmag/routes"
)

func main() {
	server := routes.SetupRouter()
	server.Listen(":8080")
}
