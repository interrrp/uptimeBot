package main

import (
	"log"

	"github.com/Tnze/go-mc/chat"
)

func onGameStart() error {
	log.Println("Game started")
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
