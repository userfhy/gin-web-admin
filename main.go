package main

import (
	"log"

	"gin-web-admin/cmd/server"
)

func main() {
	if err := server.Run(server.Options{}); err != nil {
		log.Fatalf("server exited: %v", err)
	}
}
