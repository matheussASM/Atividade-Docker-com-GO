package main

import (
	"github.com/matheussASM/pre-atividade-02/database"
	"github.com/matheussASM/pre-atividade-02/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
