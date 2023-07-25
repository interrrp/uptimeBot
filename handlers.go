package main

import (
	"log"

	"github.com/Tnze/go-mc/chat"
)

const SetupMsg = ("" +
	"uptimeBot is now up and running! " +
	"It's recommended to hide me in an unreachable area and vanish me via Essentials " +
	"so I don't get in the way.")

func onGameStart() error {
	log.Println("Game started")

	if !hasSentSetupMessage {
		err := chatHandler.SendMessage(SetupMsg)
		if err != nil {
			log.Println(err)
		}
		hasSentSetupMessage = true
	}

	return nil
}

func onDisconnect(reason chat.Message) error {
	log.Println("Rejoining due to disconnection:", reason)
	err := client.JoinServer(*addr)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func onDeath() error {
	log.Println("Died, respawning")
	err := player.Respawn()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
