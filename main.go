package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/fedelombar/starter-go/api"
	"github.com/fedelombar/starter-go/storage"
)

func main() {
	listenAddr := flag.String("listenadrr", ":3000", "the server address")
	flag.Parse()

	store := storage.NewMemoryStorage()

	server := api.NewServer(*listenAddr, store)
	fmt.Println("server is running on port:", *listenAddr)
	log.Fatal(server.Start())
}
