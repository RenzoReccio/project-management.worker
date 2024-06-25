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
	featureAzure := new(UserStoryAzure)
	azureapi_utils.GetJson(resp, featureAzure)
	assignedPerson := model.NewPerson(featureAzure.Fields.SystemAssignedTo.DisplayName,
		featureAzure.Fields.SystemAssignedTo.ID,
		featureAzure.Fields.SystemAssignedTo.UniqueName)
	parentURL := getParentURL(featureAzure)
	return model_shared.NewResultWithValueSuccess[model.UserStory](model.NewUserStory(
		featureAzure.ID, featureAzure.Fields.SystemAreaPath, featureAzure.Fields.SystemTeamProject,
		featureAzure.Fields.SystemIterationPath, featureAzure.Fields.SystemWorkItemType,
		featureAzure.Fields.SystemState, featureAzure.Fields.SystemReason, assignedPerson,
		featureAzure.Fields.SystemTitle, featureAzure.Fields.SystemBoardColumn, featureAzure.Fields.SystemBoardColumnDone,
		featureAzure.Fields.MicrosoftVSTSCommonPriority, featureAzure.Fields.MicrosoftVSTSCommonValueArea,
		featureAzure.Fields.MicrosoftVSTSCommonRisk, featureAzure.Fields.WEFE5A77436CD8D4FD8931A44E8FD000363KanbanColumn,
		featureAzure.Fields.WEFE5A77436CD8D4FD8931A44E8FD000363KanbanColumnDone, featureAzure.Fields.SystemDescription,
		featureAzure.Fields.MicrosoftVSTSCommonAcceptanceCriteria, featureAzure.Fields.SystemTags,
		nil, featureAzure.URL, nil),
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
