package config

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config representation
type Config struct {
	Port           int
	SlackToken     string
	GitlabToken    string
	GitlabApiToken string
	LogLevel       string
}

// LoadConfig is used to create an Config struct
// by merging FS / ENV / FLAGS in that order
func LoadConfig() *Config {

	var conf Config

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

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
		return nil
	}
	log.SetLevel(lvl)

	log.Info(lvl)
	if lvl == log.DebugLevel {
		//log.SetReportCaller(true)
	}

	return &conf
}
