package bot

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/iDominate/golang-discord/config"
)

var botID string
var goBot *discordgo.Session

func Start() {
	config, err := config.GetConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	goBot, err := discordgo.New(fmt.Sprintf("Bot %s", config.Token))

	if err != nil {
		log.Fatal(err.Error())
	}
	user, err := goBot.User("@me")

	if err != nil {
		log.Fatal(err.Error())
	}
	botID = user.ID

	goBot.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "ping" {
			s.ChannelMessageSend(m.ChannelID, "pong")
		}
	})
	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Bot is running...")
}
