package releasebot

import (
	"github.com/karriereat/release-bot/internal/app/releasebot/handler"
	"github.com/karriereat/release-bot/internal/pkg/config"
	serv "github.com/karriereat/release-bot/internal/pkg/server"
)

// Bot struct
type Bot struct {
	config *config.Config
	server *serv.Server
}

// NewBot returns an Bot instance
func NewBot(conf *config.Config) *Bot {
	server := serv.NewServer(conf)

	server.AddRoute("/", &handler.IndexHandler{})
	server.AddRoute("/hooks/gitlab", &handler.GitlabHandler{Config: conf})

	return &Bot{
		config: conf,
		server: server,
	}
}

// Run starts the server of the bot
func (bot *Bot) Run() {
	bot.server.Run()
}
