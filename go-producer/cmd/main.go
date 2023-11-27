package main

import (
	"go-producer/internal/api/server"
	"go-producer/internal/producer"
	"os"
	"strconv"
)

func main() {
	port := os.Getenv("HTTP_PORT")
	p, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}
	s := server.NewServer(p)
	if err != nil {
		panic(err)
	}
	producer, err := producer.NewProducer()
	if err := s.Start(producer); err != nil {
		panic(err)
	}
}
