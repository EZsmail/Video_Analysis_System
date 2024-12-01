package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	RabbitMQ `yaml:"rabbitmq"`
	MongoDB  `yaml:"mongodb"`
	LogPath  string `yaml:"log_path"`
	Debug    bool   `yaml:"debug"`
}

type RabbitMQ struct {
	URL       string `yaml:"url"`
	QueueName string `yaml:"queue_name"`
}

type MongoDB struct {
	URL              string `yaml:"url"`
	Database         string `yaml:"database"`
	CollectionStatus string `yaml:"collection_status"`
	CollectionResult string `yaml:"collection_result"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
