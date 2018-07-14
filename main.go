package main

import (
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

func main() {
	initLogger(config.Log)

	apiClient := client.Create(client.Params{
		ApiKey: config.ApiFootball.Key,
		Url:    config.ApiFootball.Url,
	})

	countries, err := apiClient.GetCountries()
	if err != nil {
		logger.Errorf("getting countries: %s", err)
	}

	leagues, err := apiClient.GetLeagues(0)
	if err != nil {
		logger.Errorf("getting leagues: %s", err)
	}

	leaguesByCountry, err := apiClient.GetLeagues(170)
	if err != nil {
		logger.Errorf("getting leagues by country with id = 79: %s", err)
	}

	standings, err := apiClient.GetStandings(79)
	if err != nil {
		logger.Errorf("getting standings: %s", err)
	}

	fmt.Printf("countries: %#v\n", countries)
	fmt.Printf("leagues: %#v\n", leagues)
	fmt.Printf("leagues by Italy: %#v\n", leaguesByCountry)
	fmt.Printf("standings by Serie A: %#v\n", standings)
}

func init() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("First argument must be path to config file")
	}

	config = &Config{}
	if _, err := toml.DecodeFile(flag.Arg(0), &config); err != nil {
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
