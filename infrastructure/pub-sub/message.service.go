package pubsub_message

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/RenzoReccio/project-management.worker/config"
	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
	"github.com/RenzoReccio/project-management.worker/domain/repository"
	"google.golang.org/api/option"
)

type MessageService struct {
	PubSubConnection *config.PubSubConnection
}

func NewMessageService(pubSubConnection *config.PubSubConnection) repository.MessageRepository {
	return &MessageService{
		PubSubConnection: pubSubConnection,
	}
}

func (c MessageService) SendEpic(in *model.Epic) *model_shared.ResultWithValue[string] {
	topicID := model_shared.EpicType

	// msg := "Hello World"
	ctx := context.Background()
	log.Printf("json %s", c.PubSubConnection.ToString())

	client, err := pubsub.NewClient(ctx, c.PubSubConnection.QuotaProjectId, option.WithCredentialsJSON([]byte(c.PubSubConnection.ToString())))
	// client, err := pubsub.NewClientWithConfig()(ctx, projectID)
	if err != nil {
		log.Printf("error %s", err)
		return model_shared.NewResultWithValueFailure[string](model_shared.NewError("PUB_SUB_FAIL", err.Error()))
	}
	defer client.Close()
	t := client.Topic(topicID)
	jsonEpic, err := json.Marshal(in)
	if err != nil {
		return model_shared.NewResultWithValueFailure[string](model_shared.NewError("PUB_SUB_FAIL", err.Error()))
	}

	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(string(jsonEpic)),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		log.Printf("error %s", err)
		return model_shared.NewResultWithValueFailure[string](model_shared.NewError("PUB_SUB_FAIL", err.Error()))
	}
	log.Printf("id %s", id)

	return model_shared.NewResultWithValueSuccess[string](&id)

}
