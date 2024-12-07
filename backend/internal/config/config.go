package config

import (
	"os"
	"strconv"
)

type Config struct {
	RabbitMQ   `yaml:"rabbitmq"`
	MongoDB    `yaml:"mongodb"`
	PostgreSQL `yaml:"postgresql"`
	LogPath    string `yaml:"log_path"`
}

type RabbitMQ struct {
	URL       string `yaml:"url"`
	QueueName string `yaml:"queue_name"`
}

type MongoDB struct {
	URL              string `yaml:"url"`
	Database         string `yaml:"database"`
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

func LoadConfig() (*Config, error) {
	return &Config{
		RabbitMQ{
			getEnv("RABBIT_URL", ""),
			getEnv("RABBIT_QUEUE_NAME", ""),
		},
		MongoDB{
			getEnv("MONGO_URL", ""),
			getEnv("MONGO_DATABASE", ""),
			getEnv("MONGO_COLLECTION_RESULTS", ""),
		},
		PostgreSQL{
			getEnv("POSTGRES_HOST", ""),
			getEnvInt("POSTGRES_PORT", 5432),
			getEnv("POSTGRES_USER", ""),
			getEnv("POSTGRES_PASSWORD", ""),
			getEnv("POSTGRES_DB", ""),
			getEnv("POSTGRES_TABLE", ""),
		},
		getEnv("LOG_PATH", "video_analizer.log"),
	}, nil
}

func getEnv(envName, defaultValue string) string {
	if env, exists := os.LookupEnv(envName); exists {
		return env
	}
	return defaultValue
}

func getEnvInt(envName string, defaultValue int) int {
	if env, exists := os.LookupEnv(envName); exists {
		res, _ := strconv.Atoi(env)
		return res
	}
	return defaultValue
}
