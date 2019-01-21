package config

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config representation
type Config struct {
	LogLevel string
	Slack    struct {
		Token     string
		Channel   string
		IconEmoji string
	}
	Gitlab struct {
		WebhookSecret string
		APIToken      string
		BaseURL       string
		SkipSSLVerify bool
	}
	Server struct {
		Port int
	}
}

// LoadConfig is used to create an Config struct
func LoadConfig(cfgFile string) *Config {

	var conf Config

	viper.SetConfigFile(cfgFile)
	viper.SetDefault("Gitlab.SkipSSLVerify", false)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&conf)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	lvl, err := log.ParseLevel(conf.LogLevel)
	if err != nil {
		log.Error("Invalid log level:" + conf.LogLevel)
	}
	log.SetLevel(lvl)

	return &conf
}
