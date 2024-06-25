package factory_workitem

import model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"

func GetWorkItemFactory(workItemType string) IWorkItem {
	switch workItemType {
	case model_shared.EpicType:
		return NewEpicFactory()
	case model_shared.FeatureType:
		return NewFeatureFactory()
	case model_shared.UserStoryType:
		return NewUserStoryFactory()
	default:
		panic("Not implemented")
	}

}
