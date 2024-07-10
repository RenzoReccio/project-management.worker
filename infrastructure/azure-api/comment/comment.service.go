package azureapi_comment

import (
	"net/http"

	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
	"github.com/RenzoReccio/project-management.worker/domain/repository"
	azureapi_utils "github.com/RenzoReccio/project-management.worker/infrastructure/azure-api/utils"
)

type CommentService struct {
	client              *http.Client
	authorizationHeader string
}

func NewCommentService(client *http.Client, authorizationHeader string) repository.CommentRepository {
	return &CommentService{
		client:              client,
		authorizationHeader: authorizationHeader,
	}
}

func (c *CommentService) GetComments(url string) *model_shared.ResultWithValue[[]model.Comment] {
	req, _ := http.NewRequest("GET", url+"/comments", nil)
	req.Header.Set("Authorization", "Basic "+c.authorizationHeader)

	resp, _ := c.client.Do(req)
	if resp.StatusCode != 200 {
		return model_shared.NewResultWithValueFailure[[]model.Comment](model_shared.NewError("AZURE_NOT_WORKING", "Azure not working."))
	}
	commentsAzure := new(CommentAzure)
	azureapi_utils.GetJson(resp, commentsAzure)
	comments := azureapi_utils.Map(commentsAzure.Comments, func(item CommentsAzure) model.Comment {
		createdBy := model.NewPerson(item.CreatedBy.DisplayName, item.CreatedBy.ID, item.CreatedBy.UniqueName)
		return *model.NewComment(item.CreatedDate, createdBy, item.Text)
	})

	return model_shared.NewResultWithValueSuccess[[]model.Comment](&comments)
}
