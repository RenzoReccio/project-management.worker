package azureapi_feature

import (
	"net/http"

	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
	"github.com/RenzoReccio/project-management.worker/domain/repository"
	azureapi_utils "github.com/RenzoReccio/project-management.worker/infrastructure/azure-api/utils"
)

type FeatureService struct {
	client              *http.Client
	authorizationHeader string
}

func NewFeatureService(client *http.Client, authorizationHeader string) repository.FeatureRepository {
	return &FeatureService{
		client:              client,
		authorizationHeader: authorizationHeader,
	}
}

func (c *FeatureService) GetFeature(url string) (*model_shared.ResultWithValue[model.Feature], string) {
	req, _ := http.NewRequest("GET", url+"?$expand=relations", nil)
	req.Header.Set("Authorization", "Basic "+c.authorizationHeader)

	resp, _ := c.client.Do(req)
	if resp.StatusCode != 200 {
		return model_shared.NewResultWithValueFailure[model.Feature](model_shared.NewError("AZURE_NOT_WORKING", "Azure not working.")), ""
	}
	featureAzure := new(FeatureAzure)
	azureapi_utils.GetJson(resp, featureAzure)
	assignedPerson := model.NewPerson(featureAzure.Fields.SystemAssignedTo.DisplayName,
		featureAzure.Fields.SystemAssignedTo.ID,
		featureAzure.Fields.SystemAssignedTo.UniqueName)
	parentURL := getParentURL(featureAzure)
	return model_shared.NewResultWithValueSuccess[model.Feature](model.NewFeature(
		featureAzure.ID, featureAzure.Fields.SystemAreaPath, featureAzure.Fields.SystemTeamProject,
		featureAzure.Fields.SystemIterationPath, featureAzure.Fields.SystemWorkItemType,
		featureAzure.Fields.SystemState, featureAzure.Fields.SystemReason, assignedPerson,
		featureAzure.Fields.SystemCreatedDate, featureAzure.Fields.SystemTitle,
		featureAzure.Fields.MicrosoftVSTSCommonPriority, featureAzure.Fields.MicrosoftVSTSCommonValueArea,
		featureAzure.Fields.MicrosoftVSTSCommonRisk, featureAzure.Fields.MicrosoftVSTSSchedulingTargetDate,
		featureAzure.Fields.MicrosoftVSTSCommonBusinessValue, featureAzure.Fields.MicrosoftVSTSCommonTimeCriticality,
		featureAzure.Fields.MicrosoftVSTSSchedulingEffort, featureAzure.Fields.MicrosoftVSTSSchedulingStartDate,
		featureAzure.Fields.SystemDescription, featureAzure.Fields.SystemTags,
		featureAzure.URL, nil, nil,
	)), parentURL
}

func getParentURL(featureAzure *FeatureAzure) string {
	parentURL := ""
	for relation := range featureAzure.Relations {
		if featureAzure.Relations[relation].Attributes.Name == "Parent" {
			parentURL = featureAzure.Relations[relation].URL
			break
		}
	}
	return parentURL
}
