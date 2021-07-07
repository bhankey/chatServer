package main

import (
	"chatServer/internal/app"
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/config.yaml", "path to configuration file")
	rand.Seed(time.Now().UnixNano())
}

// parseConfigFile returns config struct from .yaml file
func parseConfigFile(configPath string) *app.Config {
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Could not open config file: %v", err)
	}
	c := app.NewConfig()
	if err := yaml.Unmarshal(file, c); err != nil {
		log.Fatalf("Could not parse config file: %v", err)
	}
	return c
}

// @title chat server
// @version 1.0
// @description small chat server

// main configure and start URLShortener server
func main() {
	flag.Parse()
	config := parseConfigFile(configPath)

	srv, err := app.NewApp(config)
	if err != nil {
		log.Fatal(err)
	}

	if err := srv.Start(config); err != nil {
		log.Fatal(err)
	}
}
