package config

import "os"

type Config struct {
	MongoDBConnection string
	AzureToken        string
}

func NewConfig() *Config {
	return &Config{
		MongoDBConnection: os.Getenv("MONGODB"),
		AzureToken:        os.Getenv("AZURE_TOKEN"),
	}
}
