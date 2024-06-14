package config_application

import (
	application_getepic "github.com/RenzoReccio/project-management.worker/application/epic/query"
	application_createevent "github.com/RenzoReccio/project-management.worker/application/event/command/create-event"
	application_gettask "github.com/RenzoReccio/project-management.worker/application/task/query/get-task"
	application_getworkitemtype "github.com/RenzoReccio/project-management.worker/application/workItemType/query/get-work-item-type"
	config_service "github.com/RenzoReccio/project-management.worker/config/service"
	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
	"github.com/mehdihadeli/go-mediatr"
)

func InitApplication(configService *config_service.ConfigService) {
	createEventCommandHandler := application_createevent.NewCreateEventCommandHandler(configService.EventRepository)
	getTaskCommandHandler := application_gettask.NewGetTaskQueryHandler(configService.TaskRepository)
	getWorkItemTypeQueryHandler := application_getworkitemtype.NewGetWorkItemTypeQueryHandler(configService.WorkItemTypeRepository)
	getEpicQueryHandler := application_getepic.NewGetEpicQueryHandler(configService.EpicRepository, configService.CommentRepository)

	mediatr.RegisterRequestHandler[*application_createevent.CreateProductCommand, *model_shared.ResultWithValue[model.Event]](createEventCommandHandler)
	mediatr.RegisterRequestHandler[*application_gettask.GetTaskQuery, *model.Task](getTaskCommandHandler)
	mediatr.RegisterRequestHandler[*application_getworkitemtype.GetWorkItemTypeQuery, *model_shared.ResultWithValue[model.WorkItemType]](getWorkItemTypeQueryHandler)
	mediatr.RegisterRequestHandler[*application_getepic.GetEpicQuery, *model_shared.ResultWithValue[model.Epic]](getEpicQueryHandler)
}
