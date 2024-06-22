package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	MongoDBConnection string
	AzureToken        string
	PubSubConnection  *PubSubConnection
}

type PubSubConnection struct {
	ClientId       string `json:"client_id"`
	ClientSecret   string `json:"client_secret"`
	QuotaProjectId string `json:"quota_project_id"`
	Refresh_token  string `json:"refresh_token"`
	Type           string `json:"type"`
}

func NewConfig() *Config {
	return &Config{
		MongoDBConnection: os.Getenv("MONGODB"),
		AzureToken:        os.Getenv("AZURE_TOKEN"),
		PubSubConnection: &PubSubConnection{
			ClientId:       os.Getenv("CLIENT_ID_PUBSUB"),
			ClientSecret:   os.Getenv("CLIENT_SECRET_PUBSUB"),
			QuotaProjectId: os.Getenv("QUOTA_PROJECT_ID_PUBSUB"),
			Refresh_token:  os.Getenv("REFRESH_TOKEN_PUBSUB"),
			Type:           os.Getenv("TYPE_PUBSUB"),
		},
	}
}

func (c PubSubConnection) ToString() string {
	b, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(b)
}
