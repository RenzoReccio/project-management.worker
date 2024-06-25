package pubsub_message

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"cloud.google.com/go/pubsub"
	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
	"github.com/RenzoReccio/project-management.worker/domain/repository"
)

type MessageService struct {
	pubSubClient *pubsub.Client
}

func NewMessageService(pubSubClient *pubsub.Client) repository.MessageRepository {
	return &MessageService{
		pubSubClient: pubSubClient,
	}
}

func (c MessageService) SendEpic(in *model.Epic) *model_shared.ResultWithValue[string] {
	topicID := model_shared.EpicType
	ctx := context.Background()

	t := c.pubSubClient.Topic(topicID)
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

	return model_shared.NewResultWithValueSuccess[string](&id)
}

func (c MessageService) SendFeature(in *model.Feature) *model_shared.ResultWithValue[string] {
	topicID := model_shared.FeatureType
	ctx := context.Background()

	t := c.pubSubClient.Topic(topicID)
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

	return model_shared.NewResultWithValueSuccess[string](&id)
}

func (c MessageService) SendUserStory(in *model.UserStory) *model_shared.ResultWithValue[string] {
	topicID := strings.Replace(model_shared.UserStoryType, " ", "_", -1)
	ctx := context.Background()

	t := c.pubSubClient.Topic(topicID)
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

	return model_shared.NewResultWithValueSuccess[string](&id)
}
