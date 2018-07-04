package main

import (
	"log"
	"github.com/BurntSushi/toml"
	"fmt"
	"flag"
)

type Config struct {
	ApiFootball struct{
		Key string
	}
}

var config *Config

func main() {
	fmt.Printf("%#v", config)
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
