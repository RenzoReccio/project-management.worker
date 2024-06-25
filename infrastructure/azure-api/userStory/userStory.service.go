package azureapi_userstory

import (
	"net/http"

	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
	"github.com/RenzoReccio/project-management.worker/domain/repository"
	azureapi_utils "github.com/RenzoReccio/project-management.worker/infrastructure/azure-api/utils"
)

type UserStoryService struct {
	client              *http.Client
	authorizationHeader string
}

func NewUserStoryService(client *http.Client, authorizationHeader string) repository.UserStoryRepository {
	return &UserStoryService{
		client:              client,
		authorizationHeader: authorizationHeader,
	}
}

func (c *UserStoryService) GetUserStory(url string) (*model_shared.ResultWithValue[model.UserStory], string) {
	req, _ := http.NewRequest("GET", url+"?$expand=relations", nil)
	req.Header.Set("Authorization", "Basic "+c.authorizationHeader)

	resp, _ := c.client.Do(req)
	if resp.StatusCode != 200 {
		return model_shared.NewResultWithValueFailure[model.UserStory](model_shared.NewError("AZURE_NOT_WORKING", "Azure not working.")), ""
	}
	userStoryAzure := new(UserStoryAzure)
	azureapi_utils.GetJson(resp, userStoryAzure)
	assignedPerson := model.NewPerson(userStoryAzure.Fields.SystemAssignedTo.DisplayName,
		userStoryAzure.Fields.SystemAssignedTo.ID,
		userStoryAzure.Fields.SystemAssignedTo.UniqueName)
	parentURL := getParentURL(userStoryAzure)
	return model_shared.NewResultWithValueSuccess[model.UserStory](model.NewUserStory(
		userStoryAzure.ID, userStoryAzure.Fields.SystemAreaPath, userStoryAzure.Fields.SystemTeamProject,
		userStoryAzure.Fields.SystemIterationPath, userStoryAzure.Fields.SystemWorkItemType,
		userStoryAzure.Fields.SystemState, userStoryAzure.Fields.SystemReason, assignedPerson,
		userStoryAzure.Fields.SystemTitle, userStoryAzure.Fields.SystemBoardColumn, userStoryAzure.Fields.SystemBoardColumnDone,
		userStoryAzure.Fields.MicrosoftVSTSCommonPriority, userStoryAzure.Fields.MicrosoftVSTSCommonValueArea,
		userStoryAzure.Fields.MicrosoftVSTSCommonRisk, userStoryAzure.Fields.WEFE5A77436CD8D4FD8931A44E8FD000363KanbanColumn,
		userStoryAzure.Fields.WEFE5A77436CD8D4FD8931A44E8FD000363KanbanColumnDone, userStoryAzure.Fields.SystemDescription,
		userStoryAzure.Fields.MicrosoftVSTSCommonAcceptanceCriteria, userStoryAzure.Fields.SystemTags,
		nil, userStoryAzure.URL, nil, userStoryAzure.Links.HTML.Href),
	), parentURL
}

func getParentURL(userStoryAzure *UserStoryAzure) string {
	parentURL := ""
	for relation := range userStoryAzure.Relations {
		if userStoryAzure.Relations[relation].Attributes.Name == "Parent" {
			parentURL = userStoryAzure.Relations[relation].URL
			break
		}
	}
	return parentURL
}
