package config

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type Bump struct {
	MajorKeywords []string `yaml:"major_keywords"`
	MinorKeywords []string `yaml:"minor_keywords"`
	PatchKeywords []string `yaml:"patch_keywords"`
}

type Project struct {
	Name      string `yaml:"name"`
	TagPrefix string `yaml:"tag_prefix"`
	Bump      Bump   `yaml:"bump"`
}

type Config struct {
	Projects []Project `yaml:"projects"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler arquivo YAML: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("erro ao parsear YAML: %w", err)
	}

	return &cfg, nil
}
