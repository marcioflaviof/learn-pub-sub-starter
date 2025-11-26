package main

import (
	"fmt"
	"log"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/gamelogic"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
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

	publishCh, err := conn.Channel()
	if err != nil {
		log.Fatalf("could not create channel: %v", err)
	}

	gamelogic.PrintServerHelp()

	for {
		userInput := gamelogic.GetInput()
		if userInput[0] == "pause" {
			fmt.Println("Pausing game...")
			err = pubsub.PublishJSON(
				publishCh,
				routing.ExchangePerilDirect,
				routing.PauseKey,
				routing.PlayingState{
					IsPaused: true,
				},
			)

			if err != nil {
				log.Printf("could not pause the game: %v", err)
			}
		} else if userInput[0] == "resume" {
			fmt.Println("Resuming the game...")
			err = pubsub.PublishJSON(
				publishCh,
				routing.ExchangePerilDirect,
				routing.PauseKey,
				routing.PlayingState{
					IsPaused: false,
				},
			)
			if err != nil {
				log.Printf("could not resume the game: %v", err)
			}
		} else if userInput[0] == "quit" {
			fmt.Println("Quitting the server...")
			break
		} else {
			fmt.Println("Unknown command.")
		}
	}

}
