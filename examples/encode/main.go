package main

import (
	"dagobert"
	"fmt"
	"log"
)

func main() {
	client, err := dagobert.NewClient("grpc://localhost:51000")
	if err != nil {
		log.Fatal(err)
	}

	docs, err := client.Encode(
		[]*dagobert.Document{
			dagobert.NewTextDocument("hello world"),
			dagobert.NewTextDocument("hi there"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, doc := range docs {
		fmt.Println(doc.GetEmbedding().GetDense().GetBuffer())
	}
}
