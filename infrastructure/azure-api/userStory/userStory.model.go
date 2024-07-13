package azureapi_userstory

import "time"

type UserStoryAzure struct {
	ID     float64 `json:"id"`
	Rev    float64 `json:"rev"`
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
		SystemCommentCount                 float64   `json:"System.CommentCount"`
		SystemTitle                        string    `json:"System.Title"`
		SystemBoardColumn                  string    `json:"System.BoardColumn"`
		SystemBoardColumnDone              bool      `json:"System.BoardColumnDone"`
		MicrosoftVSTSSchedulingStoryPoints float64   `json:"Microsoft.VSTS.Scheduling.StoryPoints"`
		MicrosoftVSTSCommonStateChangeDate time.Time `json:"Microsoft.VSTS.Common.StateChangeDate"`
		MicrosoftVSTSCommonActivatedDate   time.Time `json:"Microsoft.VSTS.Common.ActivatedDate"`
		MicrosoftVSTSCommonActivatedBy     struct {
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
		MicrosoftVSTSCommonPriority                         float64 `json:"Microsoft.VSTS.Common.Priority"`
		MicrosoftVSTSCommonValueArea                        string  `json:"Microsoft.VSTS.Common.ValueArea"`
		MicrosoftVSTSCommonRisk                             string  `json:"Microsoft.VSTS.Common.Risk"`
		WEFE5A77436CD8D4FD8931A44E8FD000363KanbanColumn     string  `json:"WEF_E5A77436CD8D4FD8931A44E8FD000363_Kanban.Column"`
		WEFE5A77436CD8D4FD8931A44E8FD000363KanbanColumnDone bool    `json:"WEF_E5A77436CD8D4FD8931A44E8FD000363_Kanban.Column.Done"`
		SystemDescription                                   string  `json:"System.Description"`
		MicrosoftVSTSCommonAcceptanceCriteria               string  `json:"Microsoft.VSTS.Common.AcceptanceCriteria"`
		SystemTags                                          string  `json:"System.Tags"`
		SystemParent                                        float64 `json:"System.Parent"`
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
