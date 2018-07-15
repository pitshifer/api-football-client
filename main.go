package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/pitshifer/api-football-client/client"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

const (
	DefaultLogFile  = "/var/log/api-football-client.log"
	DefaultLogLevel = logrus.ErrorLevel
)

type Config struct {
	Log         LoggerConfig
	ApiFootball struct {
		Url string
		Key string
	}
}

type LoggerConfig struct {
	File  string
	Level string
}

var config *Config
var logger = logrus.New()
var action string
var countryId int
var leagueId int

func main() {
	var err error
	var result interface{}

	initLogger(config.Log)

	apiClient := client.Create(client.Params{
		ApiKey: config.ApiFootball.Key,
		Url:    config.ApiFootball.Url,
	})

	switch action {
	case "countries":
		result, err = apiClient.GetCountries()
		break

	case "leagues":
		result, err = apiClient.GetLeagues(countryId)
		break

	case "standings":
		result, err = apiClient.GetStandings(leagueId)

	default:
		fmt.Println("Not specified an action")
		os.Exit(0)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(b))
	os.Exit(0)
}

func init() {
	var configFile string

	flag.StringVar(&configFile, "c", "", "path to config file")
	flag.StringVar(&action, "action", "asd", "action name: countries, leagues")
	flag.IntVar(&countryId, "country", 0, "country ID")
	flag.IntVar(&leagueId, "league", 0, "league ID")

	flag.Parse()

	config = &Config{}
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatalf("Cannot read config file %s: %s", flag.Arg(0), err)
	}
}

func initLogger(config LoggerConfig) {
	logger.Formatter = &logrus.JSONFormatter{}

	logger.SetOutput(os.Stdout)
	if config.File != "" {
		file, err := os.OpenFile(config.File, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			logger.Errorf("cannot open log file: %s", err)
		} else {
			logger.SetOutput(file)
		}
	}
}
