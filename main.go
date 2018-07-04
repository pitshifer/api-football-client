package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/pitshifer/api-football-client/client"
	"log"
)

type Config struct {
	ApiFootball struct {
		Url string
		Key string
	}
}

var config *Config

func main() {
	apiClient := client.Create(client.Params{
		ApiKey: config.ApiFootball.Key,
	})

	fmt.Printf("%T", apiClient)
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
