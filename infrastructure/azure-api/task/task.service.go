package azureapi_task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/RenzoReccio/project-management.worker/domain/model"
	"github.com/RenzoReccio/project-management.worker/domain/repository"
)

const (
	TaskType      = "Task"
	UserStoryType = "User Story"
	FeatyreType   = "Feature"
	EpicType      = "Epic"
)

type TaskAzure struct {
	ID     int `json:"id"`
	Rev    int `json:"rev"`
	Fields struct {
		SystemAreaPath      string `json:"System.AreaPath"`
		SystemTeamProject   string `json:"System.TeamProject"`
		SystemIterationPath string `json:"System.IterationPath"`
		SystemWorkItemType  string `json:"System.WorkItemType"`
		SystemState         string `json:"System.State"`
		SystemReason        string `json:"System.Reason"`
		SystemAssignedTo    struct {
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
		} `json:"System.AssignedTo"`
		SystemCreatedDate time.Time `json:"System.CreatedDate"`
		SystemCreatedBy   struct {
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
		} `json:"System.CreatedBy"`
		SystemChangedDate time.Time `json:"System.ChangedDate"`
		SystemChangedBy   struct {
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
		} `json:"System.ChangedBy"`
		SystemCommentCount                      int       `json:"System.CommentCount"`
		SystemTitle                             string    `json:"System.Title"`
		MicrosoftVSTSSchedulingRemainingWork    int       `json:"Microsoft.VSTS.Scheduling.RemainingWork"`
		MicrosoftVSTSSchedulingOriginalEstimate int       `json:"Microsoft.VSTS.Scheduling.OriginalEstimate"`
		MicrosoftVSTSSchedulingCompletedWork    int       `json:"Microsoft.VSTS.Scheduling.CompletedWork"`
		MicrosoftVSTSCommonActivity             string    `json:"Microsoft.VSTS.Common.Activity"`
		MicrosoftVSTSCommonStateChangeDate      time.Time `json:"Microsoft.VSTS.Common.StateChangeDate"`
		MicrosoftVSTSCommonActivatedDate        time.Time `json:"Microsoft.VSTS.Common.ActivatedDate"`
		MicrosoftVSTSCommonActivatedBy          struct {
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
		} `json:"Microsoft.VSTS.Common.ActivatedBy"`
		MicrosoftVSTSCommonPriority int    `json:"Microsoft.VSTS.Common.Priority"`
		SystemDescription           string `json:"System.Description"`
		SystemTags                  string `json:"System.Tags"`
		SystemParent                int    `json:"System.Parent"`
	} `json:"fields"`
	Relations []struct {
		Rel        string `json:"rel"`
		URL        string `json:"url"`
		Attributes struct {
			IsLocked bool   `json:"isLocked"`
			Comment  string `json:"comment"`
			Name     string `json:"name"`
		} `json:"attributes"`
	} `json:"relations"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		WorkItemUpdates struct {
			Href string `json:"href"`
		} `json:"workItemUpdates"`
		WorkItemRevisions struct {
			Href string `json:"href"`
		} `json:"workItemRevisions"`
		WorkItemComments struct {
			Href string `json:"href"`
		} `json:"workItemComments"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		WorkItemType struct {
			Href string `json:"href"`
		} `json:"workItemType"`
		Fields struct {
			Href string `json:"href"`
		} `json:"fields"`
	} `json:"_links"`
	URL string `json:"url"`
}

type UserStoryAzure struct {
	ID     int `json:"id"`
	Rev    int `json:"rev"`
	Fields struct {
		SystemAreaPath      string `json:"System.AreaPath"`
		SystemTeamProject   string `json:"System.TeamProject"`
		SystemIterationPath string `json:"System.IterationPath"`
		SystemWorkItemType  string `json:"System.WorkItemType"`
		SystemState         string `json:"System.State"`
		SystemReason        string `json:"System.Reason"`
		SystemAssignedTo    struct {
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
		} `json:"System.AssignedTo"`
		SystemCreatedDate time.Time `json:"System.CreatedDate"`
		SystemCreatedBy   struct {
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
		} `json:"System.CreatedBy"`
		SystemChangedDate time.Time `json:"System.ChangedDate"`
		SystemChangedBy   struct {
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
		} `json:"System.ChangedBy"`
		SystemCommentCount                                  int       `json:"System.CommentCount"`
		SystemTitle                                         string    `json:"System.Title"`
		SystemBoardColumn                                   string    `json:"System.BoardColumn"`
		SystemBoardColumnDone                               bool      `json:"System.BoardColumnDone"`
		MicrosoftVSTSSchedulingStoryPoints                  int       `json:"Microsoft.VSTS.Scheduling.StoryPoints"`
		MicrosoftVSTSCommonStateChangeDate                  time.Time `json:"Microsoft.VSTS.Common.StateChangeDate"`
		MicrosoftVSTSCommonPriority                         int       `json:"Microsoft.VSTS.Common.Priority"`
		MicrosoftVSTSCommonValueArea                        string    `json:"Microsoft.VSTS.Common.ValueArea"`
		MicrosoftVSTSCommonRisk                             string    `json:"Microsoft.VSTS.Common.Risk"`
		WEFE5A77436CD8D4FD8931A44E8FD000363KanbanColumn     string    `json:"WEF_E5A77436CD8D4FD8931A44E8FD000363_Kanban.Column"`
		WEFE5A77436CD8D4FD8931A44E8FD000363KanbanColumnDone bool      `json:"WEF_E5A77436CD8D4FD8931A44E8FD000363_Kanban.Column.Done"`
		SystemDescription                                   string    `json:"System.Description"`
		MicrosoftVSTSCommonAcceptanceCriteria               string    `json:"Microsoft.VSTS.Common.AcceptanceCriteria"`
		SystemParent                                        int       `json:"System.Parent"`
	} `json:"fields"`
	Relations []struct {
		Rel        string `json:"rel"`
		URL        string `json:"url"`
		Attributes struct {
			IsLocked bool   `json:"isLocked"`
			Comment  string `json:"comment"`
			Name     string `json:"name"`
		} `json:"attributes"`
	} `json:"relations"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		WorkItemUpdates struct {
			Href string `json:"href"`
		} `json:"workItemUpdates"`
		WorkItemRevisions struct {
			Href string `json:"href"`
		} `json:"workItemRevisions"`
		WorkItemComments struct {
			Href string `json:"href"`
		} `json:"workItemComments"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		WorkItemType struct {
			Href string `json:"href"`
		} `json:"workItemType"`
		Fields struct {
			Href string `json:"href"`
		} `json:"fields"`
	} `json:"_links"`
	URL string `json:"url"`
}

type FeatureAzure struct {
	ID     int `json:"id"`
	Rev    int `json:"rev"`
	Fields struct {
		SystemAreaPath      string    `json:"System.AreaPath"`
		SystemTeamProject   string    `json:"System.TeamProject"`
		SystemIterationPath string    `json:"System.IterationPath"`
		SystemWorkItemType  string    `json:"System.WorkItemType"`
		SystemState         string    `json:"System.State"`
		SystemReason        string    `json:"System.Reason"`
		SystemCreatedDate   time.Time `json:"System.CreatedDate"`
		SystemCreatedBy     struct {
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
		} `json:"System.CreatedBy"`
		SystemChangedDate time.Time `json:"System.ChangedDate"`
		SystemChangedBy   struct {
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
		} `json:"System.ChangedBy"`
		SystemCommentCount                 int       `json:"System.CommentCount"`
		SystemTitle                        string    `json:"System.Title"`
		MicrosoftVSTSCommonStateChangeDate time.Time `json:"Microsoft.VSTS.Common.StateChangeDate"`
		MicrosoftVSTSCommonPriority        int       `json:"Microsoft.VSTS.Common.Priority"`
		MicrosoftVSTSCommonValueArea       string    `json:"Microsoft.VSTS.Common.ValueArea"`
		SystemDescription                  string    `json:"System.Description"`
		SystemParent                       int       `json:"System.Parent"`
	} `json:"fields"`
	Relations []struct {
		Rel        string `json:"rel"`
		URL        string `json:"url"`
		Attributes struct {
			IsLocked bool   `json:"isLocked"`
			Name     string `json:"name"`
		} `json:"attributes"`
	} `json:"relations"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		WorkItemUpdates struct {
			Href string `json:"href"`
		} `json:"workItemUpdates"`
		WorkItemRevisions struct {
			Href string `json:"href"`
		} `json:"workItemRevisions"`
		WorkItemComments struct {
			Href string `json:"href"`
		} `json:"workItemComments"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		WorkItemType struct {
			Href string `json:"href"`
		} `json:"workItemType"`
		Fields struct {
			Href string `json:"href"`
		} `json:"fields"`
	} `json:"_links"`
	URL string `json:"url"`
}

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

func getJson(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func (c *TaskService) GetTask(taskURL string) (*model.Task, error) {
	req, _ := http.NewRequest("GET", taskURL, nil)
	req.Header.Set("Authorization", "Basic "+c.authorizationHeader)

	resp, _ := c.client.Do(req)

	azureTask := new(TaskAzure)
	getJson(resp, azureTask)

	fmt.Printf("%#v", azureTask)
	return nil, nil

}

func (c *TaskService) GetTaskComments(taskURL string) (*[]model.Comment, error) {
	return nil, nil
}
