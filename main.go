package main

import (
	"github.com/hawyar/fhir/server"
)

func main() {
	s := server.NewServer()
	s.Run()
}
