package bot

import (
	"fmt"
	"github.com/Fidel-wole/discord_bot/config"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	fmt.Println(config.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err.Error())
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println("error getting user,", err.Error())
		return
	}
	BotID = u.ID
	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println("error opening connection,", err.Error())
		return
	}
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

}

//func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
//	if m.Author.ID == BotID {
//		return
//	}
//
//	if m.Content == "Hi" {
//		_, _ = s.ChannelMessageSend(m.ChannelID, "Hello there, how can i help you today!")
//	}
//}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	// Respond dynamically based on the message content
	response := fmt.Sprintf("You said: %s. How can I help you?", m.Content)
	_, _ = s.ChannelMessageSend(m.ChannelID, response)
}
