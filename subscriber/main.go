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

	// Subscribe to a subject
	subject := config.GetSubject()
	_, err = nc.Subscribe(subject, func(msg *nats.Msg) {
		fmt.Printf("Received message: %s\n", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	// Keep the subscriber alive
	select {}
}
