package main

import (
	"flag"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/karriereat/release-bot/internal/app/releasebot"
	"github.com/karriereat/release-bot/internal/pkg/config"
)

var cfgFile string

func init() {
	flag.StringVar(&cfgFile, "c", "./config.toml", "path to config file")

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetOutput(os.Stdout)
}

func main() {
	flag.Parse()

	conf := config.LoadConfig(cfgFile)
	b := releasebot.NewBot(conf)

	b.Run()

}
