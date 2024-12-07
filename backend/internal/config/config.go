package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	RabbitMQ   `yaml:"rabbitmq"`
	MongoDB    `yaml:"mongodb"`
	PostgreSQL `yaml:"postgresql"`
	LogPath    string `yaml:"log_path"`
	Debug      bool   `yaml:"debug"`
	Port       int    `yaml:"port"`
}

type RabbitMQ struct {
	URL       string `yaml:"url"`
	QueueName string `yaml:"queue_name"`
}

type MongoDB struct {
	URL              string `yaml:"url"`
	Database         string `yaml:"database"`
	CollectionStatus string `yaml:"collection_status"`
	CollectionResult string `yaml:"collection_results"`
}

type PostgreSQL struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Database    string `yaml:"database"`
	TableStatus string `yaml:"table_status"`
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
