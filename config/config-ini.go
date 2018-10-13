package config

import (
	"gopkg.in/ini.v1"
)

type iniConfig struct {
	cache map[string]string
	file  string
}

// DefaultIniConfig creates ini config with default properties source file named `config.ini`
func DefaultIniConfig() ConfigProvider {
	return iniConfig{cache: make(map[string]string), file: "config.ini"}
}

// IniConfig creates ini config with specified properties file as source.
func IniConfig(fileName string) ConfigProvider {
	return iniConfig{cache: make(map[string]string), file: fileName}
}

// GetConfig returns configuration structure read from ini file
func (i iniConfig) GetConfig() (Config, error) {
	cfg, err := ini.Load(i.file)
	if err != nil {
		return Config{}, err
	}
	c := new(Config)
	err = cfg.MapTo(c)
	return *c, err
}