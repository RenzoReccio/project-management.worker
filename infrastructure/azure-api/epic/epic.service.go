package azureapi_epic

import (
	"net/http"

	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
	"github.com/RenzoReccio/project-management.worker/domain/repository"
	azureapi_utils "github.com/RenzoReccio/project-management.worker/infrastructure/azure-api/utils"
)

type EpicService struct {
	client              *http.Client
	authorizationHeader string
}

func NewWorkItemTypeService(client *http.Client, authorizationHeader string) repository.EpicRepository {
	return &EpicService{
		client:              client,
		authorizationHeader: authorizationHeader,
	}
}

func (c *EpicService) GetEpic(url string) *model_shared.ResultWithValue[model.Epic] {
	req, _ := http.NewRequest("GET", url+"?$expand=relations", nil)
	req.Header.Set("Authorization", "Basic "+c.authorizationHeader)

	resp, _ := c.client.Do(req)
	if resp.StatusCode != 200 {
		return model_shared.NewResultWithValueFailure[model.Epic](model_shared.NewError("AZURE_NOT_WORKING", "Azure not working."))
	}
	epicAzure := new(EpicAzure)
	azureapi_utils.GetJson(resp, epicAzure)

	return model_shared.NewResultWithValueSuccess[model.Epic](model.NewEpic(
		azureapi_utils.FloatToInt(epicAzure.ID), epicAzure.Fields.SystemAreaPath, epicAzure.Fields.SystemTeamProject,
		epicAzure.Fields.SystemIterationPath, epicAzure.Fields.SystemWorkItemType, epicAzure.Fields.SystemState,
		epicAzure.Fields.SystemReason,
		model.NewPerson(epicAzure.Fields.SystemAssignedTo.DisplayName, epicAzure.Fields.SystemAssignedTo.ID, epicAzure.Fields.SystemAssignedTo.UniqueName),
		epicAzure.Fields.SystemCreatedDate, epicAzure.Fields.SystemTitle, epicAzure.Fields.SystemDescription,
		azureapi_utils.FloatToInt(epicAzure.Fields.MicrosoftVSTSCommonPriority),
		epicAzure.Fields.MicrosoftVSTSCommonValueArea, epicAzure.Fields.MicrosoftVSTSCommonRisk,
		azureapi_utils.FloatToInt(epicAzure.Fields.MicrosoftVSTSCommonBusinessValue),
		azureapi_utils.FloatToInt(epicAzure.Fields.MicrosoftVSTSCommonTimeCriticality),
		azureapi_utils.FloatToInt(epicAzure.Fields.MicrosoftVSTSSchedulingEffort),
		epicAzure.Fields.MicrosoftVSTSSchedulingStartDate, epicAzure.Fields.MicrosoftVSTSSchedulingTargetDate,
		epicAzure.URL, epicAzure.Fields.SystemTags, nil, epicAzure.Links.HTML.Href))
}
