package azureapi_task

import (
	"net/http"

	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
	"github.com/RenzoReccio/project-management.worker/domain/repository"
	azureapi_utils "github.com/RenzoReccio/project-management.worker/infrastructure/azure-api/utils"
)

type TaskService struct {
	client              *http.Client
	authorizationHeader string
}

func NewTaskService(client *http.Client, authorizationHeader string) repository.TaskRepository {
	return &TaskService{
		client:              client,
		authorizationHeader: authorizationHeader,
	}
}

func (c *TaskService) GetTask(taskURL string) (*model_shared.ResultWithValue[model.Task], string) {
	req, _ := http.NewRequest("GET", taskURL+"?$expand=relations", nil)
	req.Header.Set("Authorization", "Basic "+c.authorizationHeader)

	resp, _ := c.client.Do(req)
	if resp.StatusCode != 200 {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError("AZURE_NOT_WORKING", "Azure not working.")), ""
	}
	taskAzure := new(TaskAzure)
	azureapi_utils.GetJson(resp, taskAzure)
	assignedPerson := model.NewPerson(taskAzure.Fields.SystemAssignedTo.DisplayName,
		taskAzure.Fields.SystemAssignedTo.ID,
		taskAzure.Fields.SystemAssignedTo.UniqueName)
	parentURL := getParentURL(taskAzure)
	return model_shared.NewResultWithValueSuccess[model.Task](model.NewTask(
		azureapi_utils.FloatToInt(taskAzure.ID), taskAzure.Fields.SystemAreaPath, taskAzure.Fields.SystemTeamProject, taskAzure.Fields.SystemIterationPath,
		taskAzure.Fields.SystemWorkItemType, taskAzure.Fields.SystemState,
		taskAzure.Fields.SystemReason, assignedPerson, taskAzure.Fields.SystemTitle,
		azureapi_utils.FloatToInt(taskAzure.Fields.MicrosoftVSTSSchedulingRemainingWork),
		azureapi_utils.FloatToInt(taskAzure.Fields.MicrosoftVSTSSchedulingOriginalEstimate),
		azureapi_utils.FloatToInt(taskAzure.Fields.MicrosoftVSTSSchedulingCompletedWork),
		taskAzure.Fields.MicrosoftVSTSCommonActivity,
		azureapi_utils.FloatToInt(taskAzure.Fields.MicrosoftVSTSCommonPriority),
		taskAzure.Fields.SystemDescription, taskAzure.Fields.SystemTags,
		nil, taskAzure.URL, nil, taskAzure.Links.HTML.Href,
	),
	), parentURL
}

func getParentURL(userStoryAzure *TaskAzure) string {
	parentURL := ""
	for relation := range userStoryAzure.Relations {
		if userStoryAzure.Relations[relation].Attributes.Name == "Parent" {
			parentURL = userStoryAzure.Relations[relation].URL
			break
		}
	}
	return parentURL
}
