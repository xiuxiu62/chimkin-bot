package commands

import (
	"fmt"
	"math/rand"
	"time"

	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/justincremer/chimkin-bot/pkg/logger"
)

func HandleInfo(s *discordgo.Session, m *discordgo.Message, t0 time.Time) {
	t1 := time.Now()
	channel, err := s.Channel(m.ChannelID)
	logger.Must("Unknown channel error: ", err)

	title := "Info Panel"
	channelName := channel.Name
	message := "```txt\n%s\n%s\n%-16s%-20s\n%-16s%-20s\n%-16s%-20s```"
	message = fmt.Sprintf(message, title, strings.Repeat("-", len(title)), "ChannelID", m.ChannelID, "Channel Name", channelName, "Uptime", (t1.Sub(t0).String()))
	s.ChannelMessageSend(m.ChannelID, message)
}

func HandleHelp(s *discordgo.Session, m *discordgo.Message) {
	title := "Help Panel"
	message := "```txt\n%s\n%s\n%s```"

	subMessage := strings.Repeat("%s :: %s\n", 3)
	subMessage = fmt.Sprintf(subMessage,
		"help  ", "List of commands",
		"info  ", "How is chimkin",
		"whois ", "Doxes the person who's name you provide [ sophie, justin, liana, sunny, angela, paul, joseph, siah, fluzz, kreiker]",
	)
	message = fmt.Sprintf(message, title, strings.Repeat("-", len(title)), subMessage)
	s.ChannelMessageSend(m.ChannelID, message)
}

var messageTable = map[string][]string{
	"sophie": {
		"Toasty",
		"Have you tried the oat milk?",
		"Hella",
	},
	"justin": {
		"Girl with basket of fruit",
		"Seeing lil ghosts everywhere",
		"Ricochet the pain in a bottle of rum",
		"Eating strawberries with you",
		"See you later space cowboy",
	},
	"liana": {
		"WHEN\nWENH\nWHEN YOU\nWHEN OU\nWHEN\nwHEN YOU",
		"BUNBUN",
		"If someone plays Hello World I will cry ",
		"abannanana.. seaanemanemane.. eminemineminem..",
	},
	"sunny": {
		"Sunnu nation must rise!",
		"┻━┻ ︵ ＼(’0’)/／ ︵ ┻━┻",
		"ᕕ(ᐛ)ᕗ",
	},
	"angela": {
		"gabagooey",
		"S tier troglodite",
		"Squaters are people too",
		"Not averse to bullying",
		"Shut the fuck up about zoo tycoon",
	},
	"paul": {
		"Paul is paulgers",
		"More sweet potato than person",
		"Secretly drowning in the sewer",
	},
	"joseph": {
		"Da Bling",
		"Sunday is Jesus' day to game",
		"I promise I'm not a barn owl",
	},
	"siah": {
		"SUNNU YOU NEED TO CALM DOWN NOW.",
		"Welcome galaxy, noice to have you join.",
		"Fun fact : siah loves to redeem owa owa channel points.",
		"SUNNU STOP FLIPPING TABLES!",
		"Fun fact : siah likes to make unofficial lanaplays0 memes in his spare time.",
		"Hello. siah at your service, how can i help you?",
		"yee to the haw.",
		"YEEEEEHAAWW!!! throws hat majestically into the air whilst the sunsets in the background.",
		"Fun fact : everything is bigger in texas.",
		"don’t you dare mess with texas. i know where you live…",
	},
	"fluzz": {
		"Just add Arc<Mutex<_>>",
		"Controls chutes and shoes alike",
	},
	"kreiker": {
		"Sunday is Jesus' day to game",
		"My pride is immeasurable, and my day is much better",
	},
}

func HandlePesonalMessage(s *discordgo.Session, m *discordgo.Message, name string) {
	messages := messageTable[name]
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(messages))
	s.ChannelMessageSend(m.ChannelID, messages[i])
}

func HandleUnknown(s *discordgo.Session, m *discordgo.Message, msg string) {
	c, err := s.UserChannelCreate(m.Author.ID)
	logger.Must("Unknown command error: ", err)
	s.ChannelMessageSend(c.ID, "The command \""+msg+"\" in not recognized. Try running `!help` for a list of commands.")
}
