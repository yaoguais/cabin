package main

import "github.com/BurntSushi/toml"

// Config configurations for application
type Config struct {
	Logger struct {
		Level      string `json:"level" toml:"level" yaml:"level"`
		Formatter  string `json:"formatter" toml:"formatter" yaml:"formatter"`
		File       string `json:"file" toml:"file" yaml:"file"`
		MaxSize    int    `json:"maxSize" toml:"maxSize" yaml:"maxSize"`
		MaxAge     int    `json:"maxAge" toml:"maxAge" yaml:"maxAge"`
		MaxBackups int    `json:"maxBackups" toml:"maxBackups" yaml:"maxBackups"`
		Compress   bool   `json:"compress" toml:"compress" yaml:"compress"`
	} `json:"logger" toml:"logger" yaml:"logger"`
	Uploader UploaderConfig `json:"uploader" toml:"uploader" yaml:"uploader"`
	SlackBot SlackBotConfig `json:"slackbot" toml:"slackbot" yaml:"slackbot"`
}

// LoadConfig create new Config by load from config
func LoadConfig(name string) (*Config, error) {
	var c Config
	_, err := toml.DecodeFile(name, &c)
	return &c, err
}
