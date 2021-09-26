package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

func ReadConfig() *Config {
	config := &Config{}

	// Read the file.
	file, err := os.ReadFile("config/configuration.yaml")
	if err != nil {
		fmt.Println("Could not read the configuration file.")
		os.Exit(1)
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		fmt.Println("The configuration file contains an error.")
		os.Exit(1)
	}
	fmt.Println("Configuration successfully read.")
	return config
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
