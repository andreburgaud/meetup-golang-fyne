package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	configFile = "config.json"
)

const configContent = `
{
    "color": "Default"
}
`

type Config struct {
	Color string `json:"color"`
}

func getConfigPath() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	path := filepath.Dir(exe)
	return fmt.Sprintf("%s/%s", path, configFile), nil
}

func loadConfig() (Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}

	var config Config

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Config file does not exist, create a default one
		err := createConfigFile()
		if err != nil {
			return Config{}, err
		}
	}

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func GetAppConfig() (Config, error) {
	cfg, err := loadConfig()

	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func createConfigFile() error {
	configPath, err := getConfigPath()

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configPath, []byte(configContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

// Save the configuration file when the color changes
func SaveColor(color string) error {
	cfg, err := loadConfig()

	if err != nil {
		return err
	}

	cfg.Color = color

	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	jsonConfig, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configPath, append(jsonConfig, []byte("\n")...), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (cfg *Config) Print() {
	fmt.Printf("Color      : %s\n", cfg.Color)
}
