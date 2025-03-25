package server

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Displays []Display `yaml:"displays"`
}

type Display struct {
	ID   int    `yaml:"id" json:"id"`
	Name string `yaml:"name" json:"name"`
	Host string `yaml:"host" json:"host"`
}

var ConfigData *Config

func init() {
	configFile := os.Getenv("PIDISPLAY_CONFIG")
	if configFile == "" {
		log.Println("PIDISPLAY_CONFIG not set, using default pidisplay.yml")
		configFile = "pidisplay.yml"
	}

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error reading config file %s: %v", configFile, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}

	// Automatically assign incrementing IDs for displays if not provided (or set to 0).
	for i := range cfg.Displays {
		if cfg.Displays[i].ID == 0 {
			cfg.Displays[i].ID = i + 1
		}
	}

	ConfigData = &cfg
	log.Printf("Loaded %d displays from configuration", len(ConfigData.Displays))
}
