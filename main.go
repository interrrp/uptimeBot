// A Minecraft bot that does nothing to keep your server running.
package main

import (
	"flag"
	"log"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
)

var (
	addr     = flag.String("address", "localhost:25565", "The server address in host:port format")
	username = flag.String("username", "uptimeBot", "The username to use when connecting to the server")
)

var (
	client *bot.Client
	player *basic.Player
)

func main() {
	flag.Parse()

	client = bot.NewClient()
	client.Auth = bot.Auth{
		Name: *username,
	}

	player = basic.NewPlayer(client, basic.DefaultSettings, basic.EventsListener{
		GameStart:  onGameStart,
		Disconnect: onDisconnect,
		Death:      onDeath,
	})

	err := client.JoinServer(*addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	for {
		err := client.HandleGame()
		if err != nil {
			log.Fatal(err)
		}
	}
}
