package main

import (
	"context"
	"log"
	"test/ent"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=mysecretpassword sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	CRUD05(context.Background(), client)

}
