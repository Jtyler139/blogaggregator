package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DataBaseName 	string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	cfg := Config{}
	fileLocation, err := getConfigFilePath()
	if err != nil {
		return cfg, err
	}
	data, err := os.ReadFile(fileLocation)
	if err != nil {
		return cfg, err
	}
	jsonData := []byte(data)
	
	err = json.Unmarshal(jsonData, &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}

func (cfg *Config) SetUser(userName string) {
	cfg.CurrentUserName = userName
	err := write(cfg)
	if err != nil {
		return
	}
}

func getConfigFilePath() (string, error) {
	fullPath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath = fullPath + "/.gatorconfig.json"
	return fullPath, nil
}

func write(cfg *Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(fullPath, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}
