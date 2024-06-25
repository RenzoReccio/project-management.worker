package azureapi_feature

import "time"

type FeatureAzure struct {
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
		SystemCommentCount                 int       `json:"System.CommentCount"`
		SystemTitle                        string    `json:"System.Title"`
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
		MicrosoftVSTSCommonPriority        int       `json:"Microsoft.VSTS.Common.Priority"`
		MicrosoftVSTSCommonValueArea       string    `json:"Microsoft.VSTS.Common.ValueArea"`
		MicrosoftVSTSCommonRisk            string    `json:"Microsoft.VSTS.Common.Risk"`
		MicrosoftVSTSSchedulingTargetDate  time.Time `json:"Microsoft.VSTS.Scheduling.TargetDate"`
		MicrosoftVSTSCommonBusinessValue   int       `json:"Microsoft.VSTS.Common.BusinessValue"`
		MicrosoftVSTSCommonTimeCriticality int       `json:"Microsoft.VSTS.Common.TimeCriticality"`
		MicrosoftVSTSSchedulingEffort      int       `json:"Microsoft.VSTS.Scheduling.Effort"`
		MicrosoftVSTSSchedulingStartDate   time.Time `json:"Microsoft.VSTS.Scheduling.StartDate"`
		SystemDescription                  string    `json:"System.Description"`
		SystemTags                         string    `json:"System.Tags"`
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
