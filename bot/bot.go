package bot

import (
	"fmt"
	"github.com/Fidel-wole/discord_bot/config"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	var err error
	goBot, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err.Error())
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println("Error fetching bot user:", err.Error())
		return
	}
	BotID = u.ID

	// Register commands after bot is ready
	goBot.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot is ready. Registering commands...")
		registerCommands(s)
	})

	goBot.AddHandler(interactionHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err.Error())
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
}

func registerCommands(s *discordgo.Session) {
	if s == nil || s.State == nil || s.State.User == nil {
		fmt.Println("Session or state is not initialized. Cannot register commands.")
		return
	}

	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "info",
			Description: "Get information about the bot",
		},
		{
			Name:        "greet",
			Description: "Send a greeting",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "The name of the person to greet",
					Required:    true,
				},
			},
		},
		{
			Name:        "math",
			Description: "Perform basic arithmetic",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "operation",
					Description: "The math operation (add, subtract, multiply, divide)",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionNumber,
					Name:        "num1",
					Description: "The first number",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionNumber,
					Name:        "num2",
					Description: "The second number",
					Required:    true,
				},
			},
		},
	}

	for _, cmd := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		if err != nil {
			fmt.Printf("Error registering command '%s': %v\n", cmd.Name, err)
		} else {
			fmt.Printf("Command '%s' registered successfully.\n", cmd.Name)
		}
	}
}

// Handle interactions (slash commands)
func interactionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Name {
	case "info":
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "I am a bot that provides useful information.",
			},
		})
	case "greet":
		name := i.ApplicationCommandData().Options[0].StringValue()
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Hello, %s!", name),
			},
		})
	case "math":
		opts := i.ApplicationCommandData().Options
		operation := opts[0].StringValue()
		num1 := opts[1].FloatValue()
		num2 := opts[2].FloatValue()
		result := 0.0

		switch operation {
		case "add":
			result = num1 + num2
		case "subtract":
			result = num1 - num2
		case "multiply":
			result = num1 * num2
		case "divide":
			if num2 != 0 {
				result = num1 / num2
			} else {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Division by zero is not allowed.",
					},
				})
				return
			}
		default:
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Invalid operation.",
				},
			})
			return
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Result: %.2f", result),
			},
		})
	}
}
