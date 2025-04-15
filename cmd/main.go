package main

import (
	"github.com/pigen-dev/pigen-core/internal/api/routes"
)

func main() {
	server := routes.SetupRouter()
	server.Run(":5000")
}