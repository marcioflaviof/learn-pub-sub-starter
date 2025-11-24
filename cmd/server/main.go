package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")

	const CONNECTION = "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(CONNECTION)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Connection Successfull")

	// Wait for interrupt signal to gracefully shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Press Ctrl+C to exit...")
	<-sigs
	fmt.Println("Shutting down.")

	conn.Close()
}
