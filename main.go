package main

import (
	"os"

	"github.com/evresi/api/server"
)

func main() {
	s, err := server.NewServer(os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	s.Serve()
}
