package azureapi_workItemType

import (
	"net/http"

	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
	"github.com/RenzoReccio/project-management.worker/domain/repository"
	azureapi_utils "github.com/RenzoReccio/project-management.worker/infrastructure/azure-api/utils"
)

type WorkItemTypeAzure struct {
	Fields struct {
		SystemWorkItemType string `json:"System.WorkItemType"`
	} `json:"fields"`
}

type WorkItemTypeService struct {
	client              *http.Client
	authorizationHeader string
}

func NewWorkItemTypeService(client *http.Client, authorizationHeader string) repository.WorkItemTypeRepository {

	return &WorkItemTypeService{
		client:              client,
		authorizationHeader: authorizationHeader,
	}
}

func (c *WorkItemTypeService) GetWorkItemType(url *string) *model_shared.ResultWithValue[string] {
	req, _ := http.NewRequest("GET", *url, nil)
	req.Header.Set("Authorization", "Basic "+c.authorizationHeader)

	resp, _ := c.client.Do(req)
	if resp.StatusCode != 200 {
		return model_shared.NewResultWithValueFailure[string](model_shared.NewError("AZURE_NOT_WORKING", "Azure not working."))
	}
	workItemTypeAzure := new(WorkItemTypeAzure)
	azureapi_utils.GetJson(resp, workItemTypeAzure)

	return model_shared.NewResultWithValueSuccess[string](&workItemTypeAzure.Fields.SystemWorkItemType)
}
