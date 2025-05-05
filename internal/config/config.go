package config

import (
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v3"
)

type UTMConfig struct {
	Sources []string `yaml:"sources"`
	Mediums []string `yaml:"mediums"`
}

var (
	Config UTMConfig
	Mu     sync.RWMutex
)

func LoadConfig() error {
	// Get the absolute path to the config file
	configPath := filepath.Join("config", "utm.yaml")

	// Read the config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	// Parse the YAML
	var newConfig UTMConfig
	err = yaml.Unmarshal(data, &newConfig)
	if err != nil {
		return err
	}

	// Update the config atomically
	Mu.Lock()
	Config = newConfig
	Mu.Unlock()

	return nil
}
