package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/martensson/gramarr/radarr"
	"github.com/martensson/gramarr/sonarr"
)

type Config struct {
	Telegram TelegramConfig `json:"telegram"`
	Bot      BotConfig      `json:"bot"`
	Radarr   *radarr.Config `json:"radarr"`
	Sonarr   *sonarr.Config `json:"sonarr"`
}

type TelegramConfig struct {
	BotToken string `json:"botToken"`
}

type BotConfig struct {
	Name          string `json:"name"`
	Password      string `json:"password"`
	AdminPassword string `json:"adminPassword"`
}

func LoadConfig(configDir string) (*Config, error) {
	configPath := filepath.Join(configDir, "config.json")
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("config load error: %v", err)
	}
	var c Config
	json.Unmarshal(file, &c)
	return &c, nil
}

func ValidateConfig(c *Config) error {
	return nil
}
