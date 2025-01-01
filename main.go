package main

import (
	"fmt"
	"github.com/Fidel-wole/discord_bot/bot"
	"github.com/Fidel-wole/discord_bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	bot.Start()

	<-make(chan struct{})
	return
}
