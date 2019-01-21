package releasebot

import (
	log "github.com/Sirupsen/logrus"
	"github.com/karriereat/release-bot/internal/pkg/config"
	"github.com/karriereat/release-bot/internal/pkg/handler"
	"github.com/karriereat/release-bot/pkg/notifier"
	"github.com/karriereat/release-bot/pkg/server"
)

// Bot struct
type Bot struct {
	server   *server.Server
	notifier notifier.Notifier
}

// NewBot returns an Bot instance
func NewBot(conf *config.Config) *Bot {

	notifier := notifier.NewSlackNotifier(conf)

	server := addHandler(conf, notifier)

	go notifier.Run()

	return &Bot{
		server:   server,
		notifier: notifier,
	}
}

// Run starts the server of the bot
func (bot *Bot) Run() {
	bot.server.Run()
}

func addHandler(conf *config.Config, notifier *notifier.SlackNotifier) *server.Server {
	server := server.NewServer(conf.Server.Port)

	server.AddRoute("/", &handler.IndexHandler{})

	gitlabHandler, err := handler.NewGitlabHandler(conf, notifier.NotifyChan)
	if err != nil {
		log.Errorf("Ommiting gitlab handler: %v", err)
	} else {
		server.AddRoute("/hooks/gitlab", gitlabHandler)
	}

	return server
}
