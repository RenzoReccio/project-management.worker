package azureapi_comment

import "time"

type CommentAzure struct {
	TotalCount int             `json:"totalCount"`
	Count      int             `json:"count"`
	Comments   []CommentsAzure `json:"comments"`
}

type CommentsAzure struct {
	Mentions []struct {
		ArtifactID   string `json:"artifactId"`
		ArtifactType string `json:"artifactType"`
		CommentID    int    `json:"commentId"`
		TargetID     string `json:"targetId"`
	} `json:"mentions"`
	WorkItemID int    `json:"workItemId"`
	ID         int    `json:"id"`
	Version    int    `json:"version"`
	Text       string `json:"text"`
	CreatedBy  struct {
		DisplayName string `json:"displayName"`
		URL         string `json:"url"`
		Links       struct {
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"_links"`
		ID         string `json:"id"`
		UniqueName string `json:"uniqueName"`
		ImageURL   string `json:"imageUrl"`
		Descriptor string `json:"descriptor"`
	} `json:"createdBy"`
	CreatedDate time.Time `json:"createdDate"`
	ModifiedBy  struct {
		DisplayName string `json:"displayName"`
		URL         string `json:"url"`
		Links       struct {
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"_links"`
		ID         string `json:"id"`
		UniqueName string `json:"uniqueName"`
		ImageURL   string `json:"imageUrl"`
		Descriptor string `json:"descriptor"`
	} `json:"modifiedBy"`
	ModifiedDate time.Time `json:"modifiedDate"`
	Format       string    `json:"format"`
	RenderedText string    `json:"renderedText"`
	URL          string    `json:"url"`
}
