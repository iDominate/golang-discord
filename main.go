package main

import (
	"github.com/iDominate/golang-discord/bot"
)

func main() {
	bot.Start()

	<-make(chan struct{})
}
