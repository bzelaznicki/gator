package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	var config Config
	filePath, err := getConfigFilePath()

	if err != nil {

		return Config{}, err
	}
	fileData, err := os.ReadFile(filePath)

	if err != nil {
		return Config{}, err
	}
	err = json.Unmarshal(fileData, &config)

	if err != nil {

		return Config{}, err
	}

	return config, nil
}

func (c *Config) SetUser(username string) error {

	c.CurrentUserName = username

	err := write(c)

	if err != nil {
		return err
	}

	return nil
}

func write(c *Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, configFileName), nil
}
