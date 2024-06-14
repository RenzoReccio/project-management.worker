package config_service

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/RenzoReccio/project-management.worker/config"
	"github.com/RenzoReccio/project-management.worker/domain/repository"
	azureapi_comment "github.com/RenzoReccio/project-management.worker/infrastructure/azure-api/comment"
	azureapi_epic "github.com/RenzoReccio/project-management.worker/infrastructure/azure-api/epic"
	azureapi_task "github.com/RenzoReccio/project-management.worker/infrastructure/azure-api/task"
	azureapi_workItemType "github.com/RenzoReccio/project-management.worker/infrastructure/azure-api/workItemType"
	mongoInfraestructure "github.com/RenzoReccio/project-management.worker/infrastructure/mongo/event"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConfigService struct {
	CommentRepository      repository.CommentRepository
	EpicRepository         repository.EpicRepository
	EventRepository        repository.EventRepository
	TaskRepository         repository.TaskRepository
	WorkItemTypeRepository repository.WorkItemTypeRepository
}

func InitMongoDatabase(mongoDBConnetion string) *mongo.Database {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoDBConnetion).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	projectManagementDB := client.Database("ProjectManagementDB")
	if err := projectManagementDB.RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return projectManagementDB
}

func basicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func NewConfigService() *ConfigService {
	config := config.NewConfig()
	mongoDB := InitMongoDatabase(config.MongoDBConnection)

	headerAPI := basicAuth("user", config.AzureToken)
	client := &http.Client{Timeout: 10 * time.Second}
	return &ConfigService{
		EventRepository:        mongoInfraestructure.NewEventService(mongoDB),
		TaskRepository:         azureapi_task.NewTaskService(client, headerAPI),
		WorkItemTypeRepository: azureapi_workItemType.NewWorkItemTypeService(client, headerAPI),
		EpicRepository:         azureapi_epic.NewWorkItemTypeService(client, headerAPI),
		CommentRepository:      azureapi_comment.NewCommentService(client, headerAPI),
	}
}
