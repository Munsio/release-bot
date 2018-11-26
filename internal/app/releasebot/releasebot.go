package releasebot

import (
	"github.com/karriereat/release-bot/internal/app/releasebot/handler"
	"github.com/karriereat/release-bot/internal/pkg/config"
	"github.com/karriereat/release-bot/internal/pkg/notifier"
	serv "github.com/karriereat/release-bot/internal/pkg/server"
)

// Bot struct
type Bot struct {
	config   *config.Config
	server   *serv.Server
	notifier *notifier.SlackNotifier
}

// NewBot returns an Bot instance
func NewBot(conf *config.Config) *Bot {

	notifyChan := make(chan notifier.Message)

	server := addHandler(conf, notifyChan)

	notifier := notifier.NewSlackNotifier(conf)

	go notifier.Run(notifyChan)

	return &Bot{
		config:   conf,
		server:   server,
		notifier: notifier,
	}
}

// Run starts the server of the bot
func (bot *Bot) Run() {
	bot.server.Run()
}

func addHandler(conf *config.Config, notifyChan chan notifier.Message) *serv.Server {
	server := serv.NewServer(conf)

	server.AddRoute("/", &handler.IndexHandler{})
	server.AddRoute("/hooks/gitlab", &handler.GitlabHandler{Config: conf, NotifyChan: notifyChan})

	return server
}
