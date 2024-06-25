package config_application

import (
	application_sendepic "github.com/RenzoReccio/project-management.worker/application/epic/event/send-epic"
	application_getepic "github.com/RenzoReccio/project-management.worker/application/epic/query/get-epic"
	application_createevent "github.com/RenzoReccio/project-management.worker/application/event/command/create-event"
	application_sendfeature "github.com/RenzoReccio/project-management.worker/application/feature/event/send-feature"
	application_getfeature "github.com/RenzoReccio/project-management.worker/application/feature/query/get-feature"
	application_sendtask "github.com/RenzoReccio/project-management.worker/application/task/event/send-task"
	application_gettask "github.com/RenzoReccio/project-management.worker/application/task/query/get-task"
	application_senduserstory "github.com/RenzoReccio/project-management.worker/application/userStory/event/send-user-story"
	application_getuserstory "github.com/RenzoReccio/project-management.worker/application/userStory/query/get-user-story"
	application_getworkitemtype "github.com/RenzoReccio/project-management.worker/application/workItemType/query/get-work-item-type"
	config_service "github.com/RenzoReccio/project-management.worker/config/service"
	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
	"github.com/mehdihadeli/go-mediatr"
)

func InitApplication(configService *config_service.ConfigService) {
	createEventCommandHandler := application_createevent.NewCreateEventCommandHandler(configService.EventRepository)
	getWorkItemTypeQueryHandler := application_getworkitemtype.NewGetWorkItemTypeQueryHandler(configService.WorkItemTypeRepository)
	getEpicQueryHandler := application_getepic.NewGetEpicQueryHandler(configService.EpicRepository, configService.CommentRepository)
	sendEpicEventHandler := application_sendepic.NewSendEpicEventHandler(configService.MessageRepository)
	getFeatureQueryHandler := application_getfeature.NewGetFeatureQueryHandler(configService.EpicRepository, configService.CommentRepository, configService.FeatureRepository)
	sendFeatureEventHandler := application_sendfeature.NewSendFeatureEventHandler(configService.MessageRepository)

	getUserStoryQueryHandler := application_getuserstory.NewGetUserStoryQueryHandler(configService.EpicRepository,
		configService.CommentRepository,
		configService.FeatureRepository,
		configService.UserStoryRepository)

	sendUserStoryEventHandler := application_senduserstory.NewSendUserStoryEventHandler(configService.MessageRepository)

	getTaskCommandHandler := application_gettask.NewGetTaskQueryHandler(configService.EpicRepository,
		configService.CommentRepository,
		configService.FeatureRepository,
		configService.UserStoryRepository,
		configService.TaskRepository)
	sendTaskEventHandler := application_sendtask.NewSendTaskEventHandler(configService.MessageRepository)

	mediatr.RegisterRequestHandler[*application_createevent.CreateProductCommand, *model_shared.ResultWithValue[model.Event]](createEventCommandHandler)
	mediatr.RegisterRequestHandler[*application_getworkitemtype.GetWorkItemTypeQuery, *model_shared.ResultWithValue[model.WorkItemType]](getWorkItemTypeQueryHandler)
	mediatr.RegisterRequestHandler[*application_getepic.GetEpicQuery, *model_shared.ResultWithValue[model.Epic]](getEpicQueryHandler)
	mediatr.RegisterRequestHandler[*application_getfeature.GetFeatureQuery, *model_shared.ResultWithValue[model.Feature]](getFeatureQueryHandler)
	mediatr.RegisterRequestHandler[*application_getuserstory.GetUserStoryQuery, *model_shared.ResultWithValue[model.UserStory]](getUserStoryQueryHandler)
	mediatr.RegisterRequestHandler[*application_gettask.GetTaskQuery, *model_shared.ResultWithValue[model.Task]](getTaskCommandHandler)

	mediatr.RegisterNotificationHandlers[*application_sendepic.SendEpicEvent](sendEpicEventHandler)
	mediatr.RegisterNotificationHandlers[*application_sendfeature.SendFeatureEvent](sendFeatureEventHandler)
	mediatr.RegisterNotificationHandlers[*application_senduserstory.SendUserStoryEvent](sendUserStoryEventHandler)
	mediatr.RegisterNotificationHandlers[*application_sendtask.SendTaskEvent](sendTaskEventHandler)
}
