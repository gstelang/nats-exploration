package main

import (
	"fmt"
	"log"

	"github.com/gstelang/nats-exploration/config"
	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Publish a message to a subject
	subject := config.GetSubject()
	message := "Taylor Swift Concert!"
	err = nc.Publish(subject, []byte(message))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Published message: %s\n", message)
}
