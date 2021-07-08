package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/l3lackShark/binance-api-bot/envvars"
	"github.com/l3lackShark/binance-api-bot/web"
)

//Most of this code is from https://github.com/bwmarrin/discordgo/blob/master/examples/pingpong/main.go

// Variables used for command line parameters
var (
	Token   string
	BaseURL string = "http://localhost:24080"
)

func init() {
	envvars.LoadEnv()
	if os.Getenv("BOT_TOKEN") == "" {
		log.Fatalln("No bot token was provided (set the BOT_TOKEN env. var. '.env' file is supported)")
	}
	if os.Getenv("BASE_URL") != "" {
		BaseURL = os.Getenv("BASE_URL")
	}
	Token = os.Getenv("BOT_TOKEN")
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!btc" {
		price, err := web.SendBTCRequest(BaseURL)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("There was an error: %s", err.Error()))
			return
		}
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Average BTC Price in the last 5 MinutesL: %s", price))
	}
}
