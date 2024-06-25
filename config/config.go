package config

import (
	"context"
	"encoding/json"
	"os"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

type Config struct {
	MongoDBConnection string
	AzureToken        string
	PubSubClient      *pubsub.Client
}

type PubSubConnection struct {
	ClientId       string `json:"client_id"`
	ClientSecret   string `json:"client_secret"`
	QuotaProjectId string `json:"quota_project_id"`
	Refresh_token  string `json:"refresh_token"`
	Type           string `json:"type"`
}

func NewConfig() *Config {
	pubSubConnection := &PubSubConnection{
		ClientId:       os.Getenv("CLIENT_ID_PUBSUB"),
		ClientSecret:   os.Getenv("CLIENT_SECRET_PUBSUB"),
		QuotaProjectId: os.Getenv("QUOTA_PROJECT_ID_PUBSUB"),
		Refresh_token:  os.Getenv("REFRESH_TOKEN_PUBSUB"),
		Type:           os.Getenv("TYPE_PUBSUB"),
	}
	pubSubClient, err := pubsub.NewClient(context.TODO(), pubSubConnection.QuotaProjectId, option.WithCredentialsJSON([]byte(pubSubConnection.ToString())))
	if err != nil {
		panic(err)
	}
	return &Config{
		MongoDBConnection: os.Getenv("MONGODB"),
		AzureToken:        os.Getenv("AZURE_TOKEN"),
		PubSubClient:      pubSubClient,
	}
}

func (c PubSubConnection) ToString() string {
	b, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(b)
}
