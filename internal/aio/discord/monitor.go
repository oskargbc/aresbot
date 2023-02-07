package discord

import (
	"aresbot/internal/aio/shared"
	"aresbot/pkg/logger"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"net/url"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
)

func NewDiscordMonitor(channelIds, guildIds, keywords []string, token string) Monitor {
	return Monitor{
		ChannelIds: channelIds,
		GuildIds:   guildIds,
		Keywords:   keywords,
		Token:      token,
	}
}

type Monitor struct {
	Keywords   []string
	GuildIds   []string
	ChannelIds []string
	Token      string
}

func (s *Monitor) Run() {
	dg, err := discordgo.New(s.Token)
	if err != nil {
		logger.ErrorLogger.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(s.messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsAll

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		logger.ErrorLogger.Println("error opening connection,", err)
		return
	}
	shared.OInfoF("Now monitoring channels.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func (s *Monitor) messageCreate(session *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	/*if m.Author.ID == session.State.User.ID {
		return
	}*/

	reUrl := regexp.MustCompile("^(?:https?:\\/\\/)?(?:[^@\\/\\n]+@)?(?:www\\.)?([^:\\/\\n]+)")

	channel := m.Message.ChannelID
	guild := m.Message.GuildID

	for _, v := range s.GuildIds {
		if v == guild {
			for _, v = range s.ChannelIds {
				if v == channel {
					for _, v = range s.Keywords {
						if strings.Contains(m.Message.Content, v) {
							if reUrl.MatchString(m.Message.Content) {
								go shared.OpenBrowser(0, "http://localhost:11914/quicktask?product="+url.QueryEscape(m.Message.Content)+"&store=OTTO&size=x")
							}
						}
						if m.Message.Embeds != nil {
							for _, v := range m.Embeds {
								fmt.Println(v.URL)
								if reUrl.MatchString(v.URL) {
									go shared.OpenBrowser(0, "http://localhost:11914/quicktask?product="+url.QueryEscape(v.URL)+"&store=OTTO&size=x")
								}
							}
						}
					}
				}
			}
		}
	}
}
