package main

import (
	"encoding/json"

	config "goframework.io/config"
)

type AppConfig struct {
	PetsURL string `json:"petsURL"`
}

// Implement interfaces from `goframework.io/config/contracts.go`

func (thisRef *AppConfig) CreateOrReturnInstance() config.Config {
	return config.Load(&AppConfig{})
}

func (thisRef *AppConfig) DefaultConfig() config.Config {
	return &AppConfig{}
}

func (thisRef *AppConfig) String() string {
	bytes, err := json.Marshal(thisRef)
	if err != nil {
		// INFO: in normal app you could log this
		return ""
	}
	return string(bytes)
}

func getConfig() *AppConfig {
	appDefaultConfig := &AppConfig{}
	return appDefaultConfig.CreateOrReturnInstance().(*AppConfig)
}
