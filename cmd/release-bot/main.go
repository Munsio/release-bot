package main

import (
	"flag"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/karriereat/release-bot/internal/app/releasebot"
	"github.com/karriereat/release-bot/internal/pkg/config"
)

var configFlag string

func init() {
	flag.StringVar(&configFlag, "c", "./config.toml", "path to config file")

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetOutput(os.Stdout)
}

func main() {

	conf := config.LoadConfig()
	b := releasebot.NewBot(conf)

	b.Run()

}
