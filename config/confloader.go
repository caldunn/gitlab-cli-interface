package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"sync"
)

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		instance = readConfig()
	})
	return instance
}

func readConfig() *Config {
	tempConfig := &Config{}

	// Read the file.
	file, err := os.ReadFile("config/configuration.yaml")
	if err != nil {
		fmt.Println("Could not read the configuration file.")
		os.Exit(1)
	}

	err = yaml.Unmarshal(file, tempConfig)
	if err != nil {
		fmt.Println("The configuration file contains an error.")
		os.Exit(1)
	}
	fmt.Println("Configuration successfully read.")

	// Only save ref to config after configuration is successful.
	return tempConfig
}

type Config struct {
	GitlabConnection *GitlabConnection `yaml:"gitlab-connection"`
}

type GitlabConnection struct {
	Base          string        `yaml:"base"`
	AllowInsecure bool          `yaml:"allow-insecure-cert"`
	AccessTokens  *AccessTokens `yaml:"access-tokens"`
}

type AccessTokens struct {
	ReadOnly string `yaml:"read-only"`
}
