package config

import (
	"gopkg.in/ini.v1"
)

type Ini struct {
}

// GetDefaultConfig returns configuration structure read from file named `config.ini`
func (Ini) GetDefaultConfig() (Config, error) {
	return loadConfig("config.ini")
}

// GetDefaultConfig returns configuration structure read from specified ini file
func (Ini) GetConfigFromSource(file string) (Config, error) {
	return loadConfig(file)
}

// loadConfig loads config from specified file
func loadConfig(file string) (Config, error) {
	cfg, err := ini.Load(file)
	if err != nil {
		return Config{}, err
	}
	c := new(Config)
	err = cfg.MapTo(c)
	return *c, err
}