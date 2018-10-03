package config

import (
	"errors"
	"gopkg.in/ini.v1"
	"strings"
)

type iniConfig struct {
	cache map[string]string
	file  string
}

func DefaultIniConfig() ConfigProvider {
	return iniConfig{cache: make(map[string]string), file: "config.ini"}
}

func IniConfig(fileName string) ConfigProvider {
	return iniConfig{cache: make(map[string]string), file: fileName}
}

// GetConfig returns config for specified key-string
// It is expected that input string can contain # delimiter that separates section from key,
// (e.g. `auth#github` is for section `auth` and key `github`).
func (i iniConfig) GetConfig(sectionKey string) (string, error) {

	if sectionKey == "" {
		return "", errors.New("empty key")
	}

	if i.cache[sectionKey] != "" {
		return i.cache[sectionKey], nil
	}

	cfg, err := ini.Load(i.file)
	if err != nil {
		return "", err
	}
	parts := strings.Split(sectionKey, "#")

	if len(parts) > 2 {
		return "", errors.New("input key should contain at max one # separator")
	}
	section, key := "", ""
	if len(parts) == 2 {
		section = parts[0]
		key = parts[1]
	}
	if len(parts) == 1 {
		key = parts[0]
	}

	value := cfg.Section(section).Key(key).String()
	if value == "" {
		return "", errors.New("github authentication data from `config.ini` is empty")
	}

	i.cache[sectionKey] = value

	return value, nil
}
