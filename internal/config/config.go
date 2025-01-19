package config

import (
	"encoding/json"
	"os"
)

const configFile = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigPath() (string, error) {
	homePath, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return homePath + "/" + configFile, nil
}

func write(cfg Config) error {

	data, err := json.Marshal(cfg)

	if err != nil {
		return err
	}

	var path string

	path, err = getConfigPath()

	err = os.WriteFile(path, data, 755)

	if err != nil {
		return err
	}

	return nil
}

func (c Config) SetUser(current_user string) error {
	c.CurrentUserName = current_user

	err := write(c)

	if err != nil {
		return err
	}

	return nil
}

func Read() (Config, error) {
	path, err := getConfigPath()

	if err != nil {
		return Config{}, err
	}

	readFile, err := os.ReadFile(path)

	if err != nil {
		return Config{}, err
	}

	var config Config

	if err := json.Unmarshal(readFile, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
